package v1

import (
	"context"
	"fmt"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/records"
	"github.com/bondhansarker/ecommerce/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type productRepository struct {
	dbClient *sqlx.DB
}

func NewProductRepository(dbClient *sqlx.DB) V1Domains.ProductRepository {
	return &productRepository{
		dbClient: dbClient,
	}
}

func (productRepo productRepository) CreateRecord(ctx context.Context, inputDomain V1Domains.ProductDomain) (outputDomain V1Domains.ProductDomain, err error) {
	productRecord := records.ToProductRecord(inputDomain)

	rows, err := productRepo.dbClient.NamedQueryContext(ctx, `
		INSERT INTO products(name, description, specifications, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id)
		VALUES (:name, :description, :specifications, :brand_id, :category_id, :supplier_id, :unit_price, :discount_price, :tags, :status_id)
		RETURNING id, created_at`, productRecord)
	if err != nil {
		return V1Domains.ProductDomain{}, err
	}
	defer rows.Close()

	// Check if there is a row
	if rows.Next() {
		// Scan the id into productRecord
		err := rows.Scan(&productRecord.ID, &productRecord.CreatedAt)
		if err != nil {
			return V1Domains.ProductDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.ProductDomain{}, fmt.Errorf("no rows created")
	}

	return productRecord.ToDomain(), nil
}

func (productRepo productRepository) CreateRecordWithTransaction(tx *sqlx.Tx, inputDomain V1Domains.ProductDomain) (outputDomain V1Domains.ProductDomain, err error) {
	productRecord := records.ToProductRecord(inputDomain)

	rows, err := tx.NamedQuery(`
		INSERT INTO products(name, description, specifications, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id)
		VALUES (:name, :description, :specifications, :brand_id, :category_id, :supplier_id, :unit_price, :discount_price, :tags, :status_id)
		RETURNING id, created_at`, productRecord)
	if err != nil {
		return V1Domains.ProductDomain{}, err
	}
	defer rows.Close()

	// Check if there is a row
	if rows.Next() {
		// Scan the id into productRecord
		err := rows.Scan(&productRecord.ID, &productRecord.CreatedAt)
		if err != nil {
			return V1Domains.ProductDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.ProductDomain{}, fmt.Errorf("no rows created")
	}

	return productRecord.ToDomain(), nil
}

func (productRepo productRepository) GetRecordByID(ctx context.Context, id int64) (outputDomain V1Domains.ProductDomain, err error) {
	productRecord := records.Product{
		ID: id,
	}
	err = productRepo.dbClient.GetContext(ctx, &productRecord, `
		SELECT id, name, description, specifications, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, created_at
		FROM products WHERE id = $1`, productRecord.ID)
	if err != nil {
		return V1Domains.ProductDomain{}, err
	}

	return productRecord.ToDomain(), nil
}

func (productRepo productRepository) GetRecords(ctx context.Context, productFilterParams V1Domains.ProductFilterParams) (outputDomains []V1Domains.ProductDomain, err error) {
	var productRecords []records.Product
	logger.Debug(productFilterParams.CategoryID)
	// Build the SQL query
	query := `
		SELECT p.id, p.name, p.description, p.brand_id, p.category_id, p.supplier_id,
			p.unit_price, p.discount_price, p.tags, p.status_id, p.created_at
		FROM products p
		JOIN suppliers s ON p.supplier_id = s.id
		WHERE
			p.status_id = true
			AND (LOWER(p.name) LIKE LOWER($1) OR $1 IS NULL)
			AND (p.unit_price BETWEEN $2 AND $3 OR $2 IS NULL OR $3 IS NULL)
			AND (p.brand_id = ANY($4) OR $4 IS NULL)
			AND (p.category_id = $5 OR $5 IS NULL)
			AND (p.supplier_id = $6 OR $6 IS NULL)
			AND (s.is_verified_supplier = $7 OR $7 IS NULL)
		ORDER BY p.unit_price`

	brandIDs := pq.Array(productFilterParams.BrandIDs)
	err = productRepo.dbClient.Select(&productRecords, query, "%"+productFilterParams.Name+"%", &productFilterParams.MinPrice, &productFilterParams.MaxPrice, &brandIDs, &productFilterParams.CategoryID, &productFilterParams.SupplierID, &productFilterParams.IsVerifiedSupplier)
	if err != nil {
		return nil, err
		return
	}

	if err != nil {
		return nil, err
	}
	return records.ToArrayOfProductDomain(&productRecords), nil
}

func (productRepo productRepository) UpdateRecord(ctx context.Context, inputDomain V1Domains.ProductDomain) (err error) {
	productRecord := records.ToProductRecord(inputDomain)
	rows, err := productRepo.dbClient.NamedQueryContext(ctx, `
		UPDATE products SET name = :name, description = :description, specifications = :specifications, 
		brand_id = :brand_id, category_id = :category_id, supplier_id = :supplier_id, unit_price = :unit_price, 
		discount_price = :discount_price, tags = :tags, status_id = :status_id WHERE id = :id 
		RETURNING id`, productRecord)
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

func (productRepo productRepository) DeleteRecordByID(ctx context.Context, id int64) (err error) {
	result, err := productRepo.dbClient.ExecContext(ctx, `DELETE FROM products WHERE id = $1`, id)
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
