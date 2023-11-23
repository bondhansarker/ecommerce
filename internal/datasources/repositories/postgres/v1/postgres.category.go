package v1

import (
	"context"
	"fmt"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/records"
	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	dbClient *sqlx.DB
}

func NewCategoryRepository(dbClient *sqlx.DB) V1Domains.CategoryRepository {
	return &categoryRepository{
		dbClient: dbClient,
	}
}

func (categoryRepo categoryRepository) CreateRecord(ctx context.Context, inputDomain V1Domains.CategoryDomain) (outputDomain V1Domains.CategoryDomain, err error) {
	categoryRecord := records.ToCategoryRecord(inputDomain)

	rows, err := categoryRepo.dbClient.NamedQueryContext(ctx, `INSERT INTO categories(name, parent_id, sequence, status_id) VALUES (:name, :parent_id, :sequence, :status_id) RETURNING id`, categoryRecord)
	if err != nil {
		return V1Domains.CategoryDomain{}, err
	}
	defer rows.Close()

	// Check if there is a row
	if rows.Next() {
		// Scan the id into categoryRecord
		err := rows.Scan(&categoryRecord.ID)
		if err != nil {
			return V1Domains.CategoryDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.CategoryDomain{}, fmt.Errorf("no rows created")
	}

	return categoryRecord.ToDomain(), nil
}

func (categoryRepo categoryRepository) GetRecordByID(ctx context.Context, id int64) (outputDomain V1Domains.CategoryDomain, err error) {
	categoryRecord := records.Category{
		ID: id,
	}
	err = categoryRepo.dbClient.GetContext(ctx, &categoryRecord, `SELECT id, name, parent_id, sequence, status_id, created_at FROM categories WHERE id = $1`, categoryRecord.ID)
	if err != nil {
		return V1Domains.CategoryDomain{}, err
	}

	return categoryRecord.ToDomain(), nil
}

func (categoryRepo categoryRepository) GetRecords(ctx context.Context) (outputDomains []V1Domains.CategoryDomain, err error) {
	var categoryRecords []records.Category
	query := `SELECT id, name, parent_id, sequence, status_id, created_at FROM categories where status_id = true ORDER BY sequence `

	err = categoryRepo.dbClient.SelectContext(ctx, &categoryRecords, query)
	if err != nil {
		return nil, err
	}
	return records.ToArrayOfCategoryDomain(&categoryRecords), nil
}

func (categoryRepo categoryRepository) UpdateRecord(ctx context.Context, inputDomain V1Domains.CategoryDomain) (err error) {
	categoryRecord := records.ToCategoryRecord(inputDomain)
	rows, err := categoryRepo.dbClient.NamedQueryContext(ctx, `UPDATE categories SET name = :name, parent_id = :parent_id, sequence = :sequence, status_id = :status_id WHERE id = :id  RETURNING id`, categoryRecord)
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

func (categoryRepo categoryRepository) DeleteRecordByID(ctx context.Context, id int64) (err error) {
	result, err := categoryRepo.dbClient.ExecContext(ctx, `DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	affectedRowCount, err := result.RowsAffected()
	if affectedRowCount == 0 {
		return constants.ErrDBNoRowsFound
	} else {
		return err
	}
	return nil
}
