package v1

import (
	"context"
	"fmt"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/datasources/records"
	"github.com/jmoiron/sqlx"
)

type supplierRepository struct {
	dbClient *sqlx.DB
}

func NewSupplierRepository(dbClient *sqlx.DB) V1Domains.SupplierRepository {
	return &supplierRepository{
		dbClient: dbClient,
	}
}

func (supplierRepo supplierRepository) CreateRecord(ctx context.Context, inputDomain V1Domains.SupplierDomain) (outputDomain V1Domains.SupplierDomain, err error) {
	supplierRecord := records.ToSupplierRecord(inputDomain)

	rows, err := supplierRepo.dbClient.NamedQueryContext(ctx, `INSERT INTO suppliers(name, email, phone, status_id, is_verified_supplier) VALUES (:name, :email, :phone, :status_id, :is_verified_supplier) RETURNING id`, supplierRecord)
	if err != nil {
		return V1Domains.SupplierDomain{}, err
	}
	defer rows.Close()
	// Check if there is a row
	if rows.Next() {
		// Scan the id into supplierRecord
		err := rows.Scan(&supplierRecord.ID)
		if err != nil {
			return V1Domains.SupplierDomain{}, fmt.Errorf("failed to scan id: %v", err)
		}
	} else {
		return V1Domains.SupplierDomain{}, fmt.Errorf("no rows created")
	}

	return supplierRecord.ToDomain(), nil
}

func (supplierRepo supplierRepository) GetRecordByID(ctx context.Context, id int64) (outputDomain V1Domains.SupplierDomain, err error) {
	supplierRecord := records.Supplier{
		ID: id,
	}
	err = supplierRepo.dbClient.GetContext(ctx, &supplierRecord, `SELECT id, name, email, phone, status_id, is_verified_supplier, created_at FROM suppliers WHERE id = $1`, supplierRecord.ID)
	if err != nil {
		return V1Domains.SupplierDomain{}, err
	}

	return supplierRecord.ToDomain(), nil
}

func (supplierRepo supplierRepository) GetRecords(ctx context.Context) (outputDomains []V1Domains.SupplierDomain, err error) {
	var supplierRecords []records.Supplier
	query := `SELECT id, name, email, phone, status_id, is_verified_supplier, created_at FROM suppliers WHERE status_id = true`
	err = supplierRepo.dbClient.SelectContext(ctx, &supplierRecords, query)
	if err != nil {
		return nil, err
	}
	return records.ToArrayOfSupplierDomain(&supplierRecords), nil
}

func (supplierRepo supplierRepository) UpdateRecord(ctx context.Context, inputDomain V1Domains.SupplierDomain) (err error) {
	supplierRecord := records.ToSupplierRecord(inputDomain)
	rows, err := supplierRepo.dbClient.NamedQueryContext(ctx, `UPDATE suppliers SET name = :name, email = :email, phone = :phone, status_id = :status_id, is_verified_supplier = :is_verified_supplier WHERE id = :id RETURNING id`, supplierRecord)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Check if there is a row
	if !rows.Next() {
		return fmt.Errorf("no rows found")
	}
	return nil
}

func (supplierRepo supplierRepository) DeleteRecordByID(ctx context.Context, id int64) (err error) {
	result, err := supplierRepo.dbClient.ExecContext(ctx, `DELETE FROM suppliers WHERE id = $1`, id)
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
