package v1

import (
	"context"
	"fmt"

	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/records"
	"github.com/jmoiron/sqlx"
)

type brandRepository struct {
	dbClient *sqlx.DB
}

func NewBrandRepository(dbClient *sqlx.DB) V1Domains.BrandRepository {
	return &brandRepository{
		dbClient: dbClient,
	}
}

func (brandRepo brandRepository) CreateRecord(ctx context.Context, inputDomain V1Domains.BrandDomain) (outputDomain V1Domains.BrandDomain, err error) {
	brandRecord := records.ToBrandRecord(inputDomain)

	rows, err := brandRepo.dbClient.NamedQueryContext(ctx, `INSERT INTO brands(name, status_id) VALUES (:name, :status_id) RETURNING id, created_at`, brandRecord)
	if err != nil {
		return V1Domains.BrandDomain{}, err
	}
	defer rows.Close()
	// Check if there is a row
	if rows.Next() {
		// Scan the id into brandRecord
		err := rows.Scan(&brandRecord.ID, &brandRecord.CreatedAt)
		if err != nil {
			return V1Domains.BrandDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.BrandDomain{}, fmt.Errorf("no rows created")
	}

	return brandRecord.ToDomain(), nil
}

func (brandRepo brandRepository) GetRecordByID(ctx context.Context, id int64) (outputDomain V1Domains.BrandDomain, err error) {
	brandRecord := records.Brand{
		ID: id,
	}
	err = brandRepo.dbClient.GetContext(ctx, &brandRecord, `SELECT id, name, status_id, created_at FROM brands WHERE id = $1`, brandRecord.ID)
	if err != nil {
		return V1Domains.BrandDomain{}, err
	}

	return brandRecord.ToDomain(), nil
}

func (brandRepo brandRepository) GetRecords(ctx context.Context) (outputDomains []V1Domains.BrandDomain, err error) {
	var brandRecords []records.Brand
	query := `SELECT id, name, status_id, created_at FROM brands where status_id = true`

	err = brandRepo.dbClient.SelectContext(ctx, &brandRecords, query)
	if err != nil {
		return nil, err
	}
	return records.ToArrayOfBrandDomain(&brandRecords), nil
}

func (brandRepo brandRepository) UpdateRecord(ctx context.Context, inputDomain V1Domains.BrandDomain) (err error) {
	brandRecord := records.ToBrandRecord(inputDomain)
	rows, err := brandRepo.dbClient.NamedQueryContext(ctx, `UPDATE brands SET name = :name , status_id = :status_id WHERE id = :id RETURNING id`, brandRecord)
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

func (brandRepo brandRepository) DeleteRecordByID(ctx context.Context, id int64) (err error) {
	result, err := brandRepo.dbClient.ExecContext(ctx, `DELETE FROM brands WHERE id = $1`, id)
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
