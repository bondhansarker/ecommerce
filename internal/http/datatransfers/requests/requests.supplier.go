package requests

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type SupplierCreateRequest struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name" validate:"required"`
	Email              string    `json:"email" validate:"required,email"`
	Phone              string    `json:"phone" validate:"required"`
	StatusID           int8      `json:"status_id" validate:"required"`
	IsVerifiedSupplier *bool     `json:"is_verified_supplier"`
	CreatedAt          time.Time `json:"created_at"`
}

func (supplierCreateRequest SupplierCreateRequest) ToV1Domain() V1Domains.SupplierDomain {
	return V1Domains.SupplierDomain{
		ID:                 supplierCreateRequest.ID,
		Name:               supplierCreateRequest.Name,
		Email:              supplierCreateRequest.Email,
		Phone:              supplierCreateRequest.Phone,
		StatusID:           helpers.IntegerToBoolean[supplierCreateRequest.StatusID],
		IsVerifiedSupplier: supplierCreateRequest.IsVerifiedSupplier,
		CreatedAt:          supplierCreateRequest.CreatedAt,
	}
}

type SupplierUpdateRequest struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name" validate:"required"`
	Email              string `json:"email" validate:"required,email"`
	Phone              string `json:"phone" validate:"required"`
	StatusID           int8   `json:"status_id" validate:"required"`
	IsVerifiedSupplier *bool  `json:"is_verified_supplier"`
}

func (supplierUpdateRequest SupplierUpdateRequest) ToV1Domain() V1Domains.SupplierDomain {
	return V1Domains.SupplierDomain{
		ID:                 supplierUpdateRequest.ID,
		Name:               supplierUpdateRequest.Name,
		Email:              supplierUpdateRequest.Email,
		Phone:              supplierUpdateRequest.Phone,
		StatusID:           helpers.IntegerToBoolean[supplierUpdateRequest.StatusID],
		IsVerifiedSupplier: supplierUpdateRequest.IsVerifiedSupplier,
	}
}
