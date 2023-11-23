package requests

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
)

type ProductQueryParams struct {
	Name               string   `form:"name"`
	MinPrice           *float64 `form:"min_price"`
	MaxPrice           *float64 `form:"max_price"`
	BrandIDs           []int64  `form:"brand_ids"`
	CategoryID         *int64   `form:"category_id"`
	SupplierID         *int64   `form:"supplier_id"`
	IsVerifiedSupplier *bool    `form:"is_verified_supplier"`
}

func (productQueryParams ProductQueryParams) ToV1ProductFilterParams() V1Domains.ProductFilterParams {
	return V1Domains.ProductFilterParams{
		Name:               productQueryParams.Name,
		MinPrice:           productQueryParams.MinPrice,
		MaxPrice:           productQueryParams.MaxPrice,
		BrandIDs:           productQueryParams.BrandIDs,
		CategoryID:         productQueryParams.CategoryID,
		SupplierID:         productQueryParams.SupplierID,
		IsVerifiedSupplier: productQueryParams.IsVerifiedSupplier,
	}
}

type ProductCreateRequest struct {
	Name           string  `json:"name" validate:"required"`
	Description    string  `json:"description"`
	Specifications string  `json:"specifications"`
	BrandID        int64   `json:"brand_id" validate:"required"`
	CategoryID     int64   `json:"category_id" validate:"required"`
	SupplierID     int64   `json:"supplier_id" validate:"required"`
	UnitPrice      float64 `json:"unit_price" validate:"required"`
	DiscountPrice  float64 `json:"discount_price"`
	Tags           string  `json:"tags"`
	StockQuantity  int64   `json:"stock_quantity" validate:"required"`
	StatusID       int8    `json:"status_id" validate:"required"`
}

func (createRequest ProductCreateRequest) ToV1Domain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		Name:           createRequest.Name,
		Description:    createRequest.Description,
		Specifications: createRequest.Specifications,
		BrandID:        createRequest.BrandID,
		CategoryID:     createRequest.CategoryID,
		SupplierID:     createRequest.SupplierID,
		UnitPrice:      createRequest.UnitPrice,
		DiscountPrice:  createRequest.DiscountPrice,
		Tags:           createRequest.Tags,
		StatusID:       helpers.IntegerToBoolean[createRequest.StatusID],
		StockQuantity:  createRequest.StockQuantity,
	}
}

type ProductUpdateRequest struct {
	ID             int64   `json:"id" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	Description    string  `json:"description"`
	Specifications string  `json:"specifications"`
	BrandID        int64   `json:"brand_id" validate:"required"`
	CategoryID     int64   `json:"category_id" validate:"required"`
	SupplierID     int64   `json:"supplier_id" validate:"required"`
	UnitPrice      float64 `json:"unit_price" validate:"required"`
	DiscountPrice  float64 `json:"discount_price"`
	Tags           string  `json:"tags"`
	StatusID       int8    `json:"status_id" validate:"required"`
}

func (updateRequest ProductUpdateRequest) ToV1Domain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:             updateRequest.ID,
		Name:           updateRequest.Name,
		Description:    updateRequest.Description,
		Specifications: updateRequest.Specifications,
		BrandID:        updateRequest.BrandID,
		CategoryID:     updateRequest.CategoryID,
		SupplierID:     updateRequest.SupplierID,
		UnitPrice:      updateRequest.UnitPrice,
		DiscountPrice:  updateRequest.DiscountPrice,
		Tags:           updateRequest.Tags,
		StatusID:       helpers.IntegerToBoolean[updateRequest.StatusID],
	}
}
