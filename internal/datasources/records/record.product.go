package records

import (
	"time"

	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
)

type Product struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Specifications string    `db:"specifications"`
	BrandID        int64     `db:"brand_id"`
	CategoryID     int64     `db:"category_id"`
	SupplierID     int64     `db:"supplier_id"`
	UnitPrice      float64   `db:"unit_price"`
	DiscountPrice  float64   `db:"discount_price"`
	Tags           string    `db:"tags"`
	StatusID       *bool     `db:"status_id"`
	CreatedAt      time.Time `db:"created_at"`
	TotalCount     int64     `db:"total_count" json:"-"`
}

// ToProductRecord converts a ProductDomain to a Product record
func ToProductRecord(domain V1Domains.ProductDomain) Product {
	return Product{
		ID:             domain.ID,
		Name:           domain.Name,
		Description:    domain.Description,
		Specifications: domain.Specifications,
		BrandID:        domain.BrandID,
		CategoryID:     domain.CategoryID,
		SupplierID:     domain.SupplierID,
		UnitPrice:      domain.UnitPrice,
		DiscountPrice:  domain.DiscountPrice,
		Tags:           domain.Tags,
		StatusID:       domain.StatusID,
		CreatedAt:      domain.CreatedAt,
	}
}

// ToDomain converts a Product record to a ProductDomain
func (model Product) ToDomain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:             model.ID,
		Name:           model.Name,
		Description:    model.Description,
		Specifications: model.Specifications,
		BrandID:        model.BrandID,
		CategoryID:     model.CategoryID,
		SupplierID:     model.SupplierID,
		UnitPrice:      model.UnitPrice,
		DiscountPrice:  model.DiscountPrice,
		Tags:           model.Tags,
		StatusID:       model.StatusID,
		CreatedAt:      model.CreatedAt,
	}
}

// ToArrayOfProductDomain converts a slice of Product records to a slice of ProductDomain
func ToArrayOfProductDomain(records *[]Product) []V1Domains.ProductDomain {
	var result []V1Domains.ProductDomain

	for _, record := range *records {
		result = append(result, record.ToDomain())
	}

	return result
}
