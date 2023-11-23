package records

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

// Supplier represents a supplier database record
type Supplier struct {
	ID                 int64     `db:"id"`
	Name               string    `db:"name"`
	Email              string    `db:"email"`
	Phone              string    `db:"phone"`
	StatusID           *bool     `db:"status_id"`
	IsVerifiedSupplier *bool     `db:"is_verified_supplier"`
	CreatedAt          time.Time `db:"created_at"`
}

// ToSupplierRecord creates a new Supplier record from SupplierDomain
func ToSupplierRecord(domain V1Domains.SupplierDomain) Supplier {
	return Supplier{
		ID:                 domain.ID,
		Name:               domain.Name,
		Email:              domain.Email,
		Phone:              domain.Phone,
		StatusID:           domain.StatusID,
		IsVerifiedSupplier: domain.IsVerifiedSupplier,
	}
}

// ToDomain converts Supplier record to SupplierDomain
func (r *Supplier) ToDomain() V1Domains.SupplierDomain {
	return V1Domains.SupplierDomain{
		ID:                 r.ID,
		Name:               r.Name,
		Email:              r.Email,
		Phone:              r.Phone,
		StatusID:           r.StatusID,
		IsVerifiedSupplier: r.IsVerifiedSupplier,
		CreatedAt:          r.CreatedAt,
	}
}

// ToArrayOfSupplierDomain converts a slice of Supplier records to a slice of SupplierDomains
func ToArrayOfSupplierDomain(records *[]Supplier) []V1Domains.SupplierDomain {
	var result []V1Domains.SupplierDomain

	for _, record := range *records {
		result = append(result, record.ToDomain())
	}

	return result
}
