package requests

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type CategoryCreateRequest struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	ParentID  *int64    `json:"parent_id"`
	Sequence  uint      `json:"sequence"`
	StatusID  int8      `json:"status_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (categoryCreateRequest CategoryCreateRequest) ToV1Domain() V1Domains.CategoryDomain {
	return V1Domains.CategoryDomain{
		ID:        categoryCreateRequest.ID,
		Name:      categoryCreateRequest.Name,
		ParentID:  categoryCreateRequest.ParentID,
		Sequence:  categoryCreateRequest.Sequence,
		StatusID:  helpers.IntegerToBoolean[categoryCreateRequest.StatusID],
		CreatedAt: categoryCreateRequest.CreatedAt,
	}
}

type CategoryUpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	ParentID *int64 `json:"parent_id"`
	Sequence uint   `json:"sequence"`
	StatusID int8   `json:"status_id" validate:"required"`
}

func (categoryUpdateRequest CategoryUpdateRequest) ToV1Domain() V1Domains.CategoryDomain {
	return V1Domains.CategoryDomain{
		ID:       categoryUpdateRequest.ID,
		Name:     categoryUpdateRequest.Name,
		ParentID: categoryUpdateRequest.ParentID,
		Sequence: categoryUpdateRequest.Sequence,
		StatusID: helpers.IntegerToBoolean[categoryUpdateRequest.StatusID],
	}
}
