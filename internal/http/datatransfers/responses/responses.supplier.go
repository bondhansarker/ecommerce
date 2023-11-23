package responses

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type SupplierResponse struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	StatusID           int8      `json:"status_id"`
	IsVerifiedSupplier *bool     `json:"is_verified_supplier"`
	CreatedAt          time.Time `json:"created_at"`
}

func (supplierResponse *SupplierResponse) ToV1Domain() V1Domains.SupplierDomain {
	return V1Domains.SupplierDomain{
		ID:                 supplierResponse.ID,
		Name:               supplierResponse.Name,
		Email:              supplierResponse.Email,
		Phone:              supplierResponse.Phone,
		StatusID:           helpers.IntegerToBoolean[supplierResponse.StatusID],
		IsVerifiedSupplier: supplierResponse.IsVerifiedSupplier,
		CreatedAt:          supplierResponse.CreatedAt,
	}
}

func FromV1DomainToSupplierResponse(supplierDomain V1Domains.SupplierDomain) SupplierResponse {
	return SupplierResponse{
		ID:                 supplierDomain.ID,
		Name:               supplierDomain.Name,
		Email:              supplierDomain.Email,
		Phone:              supplierDomain.Phone,
		StatusID:           helpers.BooleanToInteger[*supplierDomain.StatusID],
		IsVerifiedSupplier: supplierDomain.IsVerifiedSupplier,
		CreatedAt:          supplierDomain.CreatedAt,
	}
}

func ToResponseListOfSuppliers(domains []V1Domains.SupplierDomain) []SupplierResponse {
	var result = make([]SupplierResponse, 0)

	for _, val := range domains {
		result = append(result, FromV1DomainToSupplierResponse(val))
	}

	return result
}
