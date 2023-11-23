package requests

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type BrandCreateRequest struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	StatusID  int8      `json:"status_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (brandCreateRequest BrandCreateRequest) ToV1Domain() V1Domains.BrandDomain {
	return V1Domains.BrandDomain{
		ID:        brandCreateRequest.ID,
		Name:      brandCreateRequest.Name,
		StatusID:  helpers.IntegerToBoolean[brandCreateRequest.StatusID],
		CreatedAt: brandCreateRequest.CreatedAt,
	}
}

type BrandUpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	StatusID int8   `json:"status_id" validate:"required"`
}

func (brandUpdateRequest BrandUpdateRequest) ToV1Domain() V1Domains.BrandDomain {
	return V1Domains.BrandDomain{
		ID:       brandUpdateRequest.ID,
		Name:     brandUpdateRequest.Name,
		StatusID: helpers.IntegerToBoolean[brandUpdateRequest.StatusID],
	}
}
