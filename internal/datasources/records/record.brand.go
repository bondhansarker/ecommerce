package records

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

type Brand struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	StatusID  *bool     `db:"status_id"`
	CreatedAt time.Time `db:"created_at"`
}

func ToBrandRecord(domain V1Domains.BrandDomain) Brand {
	return Brand{
		ID:        domain.ID,
		Name:      domain.Name,
		StatusID:  domain.StatusID,
		CreatedAt: domain.CreatedAt,
	}
}

func (model Brand) ToDomain() V1Domains.BrandDomain {
	return V1Domains.BrandDomain{
		ID:        model.ID,
		Name:      model.Name,
		StatusID:  model.StatusID,
		CreatedAt: model.CreatedAt,
	}
}

func ToArrayOfBrandDomain(records *[]Brand) []V1Domains.BrandDomain {
	var result []V1Domains.BrandDomain

	for _, record := range *records {
		result = append(result, record.ToDomain())
	}

	return result
}
