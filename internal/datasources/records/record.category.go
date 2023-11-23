package records

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

type Category struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	ParentID  *int64    `db:"parent_id"`
	Sequence  uint      `db:"sequence"`
	StatusID  *bool     `db:"status_id"`
	CreatedAt time.Time `db:"created_at"`
}

func ToCategoryRecord(domain V1Domains.CategoryDomain) Category {
	return Category{
		ID:        domain.ID,
		Name:      domain.Name,
		ParentID:  domain.ParentID,
		Sequence:  domain.Sequence,
		StatusID:  domain.StatusID,
		CreatedAt: domain.CreatedAt,
	}
}

func (model Category) ToDomain() V1Domains.CategoryDomain {
	return V1Domains.CategoryDomain{
		ID:        model.ID,
		Name:      model.Name,
		ParentID:  model.ParentID,
		Sequence:  model.Sequence,
		StatusID:  model.StatusID,
		CreatedAt: model.CreatedAt,
	}
}

func ToArrayOfCategoryDomain(records *[]Category) []V1Domains.CategoryDomain {
	var result []V1Domains.CategoryDomain

	for _, record := range *records {
		result = append(result, record.ToDomain())
	}

	return result
}
