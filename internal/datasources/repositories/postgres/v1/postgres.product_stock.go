package v1

import (
	"context"
	"fmt"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/records"
	"github.com/jmoiron/sqlx"
	"time"
)

type productStockRepository struct {
	dbClient *sqlx.DB
}

func NewProductStockRepository(dbClient *sqlx.DB) V1Domains.ProductStockRepository {
	return &productStockRepository{
		dbClient: dbClient,
	}
}

func (productStockRepo productStockRepository) CreateRecord(ctx context.Context, inputDomain V1Domains.ProductStockDomain) (outputDomain V1Domains.ProductStockDomain, err error) {
	productStockRecord := records.ToProductStockRecord(inputDomain)
	rows, err := productStockRepo.dbClient.NamedQueryContext(ctx, `INSERT INTO product_stocks (product_id, stock_quantity) VALUES (:product_id, :stock_quantity) RETURNING id`, productStockRecord)
	if err != nil {
		return V1Domains.ProductStockDomain{}, err
	}
	defer rows.Close()
	// Check if there is a row
	if rows.Next() {
		// Scan the id into productStockRecord
		err := rows.Scan(&productStockRecord.ID)
		if err != nil {
			return V1Domains.ProductStockDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.ProductStockDomain{}, fmt.Errorf("no rows created")
	}

	return productStockRecord.ToDomain(), nil
}

func (productStockRepo productStockRepository) CreateRecordWithTransaction(tx *sqlx.Tx, inputDomain V1Domains.ProductStockDomain) (outputDomain V1Domains.ProductStockDomain, err error) {
	productStockRecord := records.ToProductStockRecord(inputDomain)
	rows, err := tx.NamedQuery(`INSERT INTO product_stocks (product_id, stock_quantity) VALUES (:product_id, :stock_quantity) RETURNING id`, productStockRecord)
	if err != nil {
		return V1Domains.ProductStockDomain{}, err
	}
	defer rows.Close()
	// Check if there is a row
	if rows.Next() {
		// Scan the id into productStockRecord
		err := rows.Scan(&productStockRecord.ID)
		if err != nil {
			return V1Domains.ProductStockDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.ProductStockDomain{}, fmt.Errorf("no rows created")
	}

	return productStockRecord.ToDomain(), nil
}

func (productStockRepo productStockRepository) GetRecordByID(ctx context.Context, id int64) (outputDomain V1Domains.ProductStockDomain, err error) {
	productStockRecord := records.ProductStock{
		ID: id,
	}
	err = productStockRepo.dbClient.GetContext(ctx, &productStockRecord, `SELECT id, product_id, stock_quantity, updated_at FROM product_stocks WHERE id = $1`, productStockRecord.ID)
	if err != nil {
		return V1Domains.ProductStockDomain{}, err
	}

	return productStockRecord.ToDomain(), nil
}

func (productStockRepo productStockRepository) GetRecordByProductID(ctx context.Context, productID int64) (outputDomain V1Domains.ProductStockDomain, err error) {
	productStockRecord := records.ProductStock{
		ProductID: productID,
	}
	err = productStockRepo.dbClient.GetContext(ctx, &productStockRecord, `SELECT id, product_id, stock_quantity, updated_at FROM product_stocks WHERE product_id = $1`, productStockRecord.ProductID)
	if err != nil {
		return V1Domains.ProductStockDomain{}, err
	}

	return productStockRecord.ToDomain(), nil
}

func (productStockRepo productStockRepository) GetRecords(ctx context.Context) (outputDomains []V1Domains.ProductStockDomain, err error) {
	var productStockRecords []records.ProductStock
	query := `SELECT id, product_id, stock_quantity, updated_at FROM product_stocks`

	err = productStockRepo.dbClient.SelectContext(ctx, &productStockRecords, query)
	if err != nil {
		return nil, err
	}
	return records.ToArrayOfProductStockDomain(&productStockRecords), nil
}

func (productStockRepo productStockRepository) UpdateRecord(ctx context.Context, inputDomain V1Domains.ProductStockDomain) (err error) {
	productStockRecord := records.ToProductStockRecord(inputDomain)
	productStockRecord.UpdatedAt = time.Now().UTC()

	rows, err := productStockRepo.dbClient.NamedQueryContext(ctx, `UPDATE product_stocks SET stock_quantity = :stock_quantity, updated_at = :updated_at WHERE product_id = :product_id RETURNING id`, productStockRecord)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Check if there is a row
	if !rows.Next() {
		return constants.ErrDBNoRowsFound
	}
	return nil
}

func (productStockRepo productStockRepository) DeleteRecordByID(ctx context.Context, id int64) (err error) {
	result, err := productStockRepo.dbClient.ExecContext(ctx, `DELETE FROM product_stocks WHERE id = $1`, id)
	if err != nil {
		return err
	}
	affectedRowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRowCount == 0 {
		return constants.ErrDBNoRowsFound
	}
	return nil
}
