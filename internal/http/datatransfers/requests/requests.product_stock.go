package requests

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

type ProductStockCreateRequest struct {
	ID            int64     `json:"id"`
	ProductID     int64     `json:"product_id" validate:"required"`
	StockQuantity int64     `json:"stock_quantity" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (productStockCreateRequest ProductStockCreateRequest) ToV1Domain() V1Domains.ProductStockDomain {
	return V1Domains.ProductStockDomain{
		ID:            productStockCreateRequest.ID,
		ProductID:     productStockCreateRequest.ProductID,
		StockQuantity: productStockCreateRequest.StockQuantity,
		UpdatedAt:     productStockCreateRequest.UpdatedAt,
	}
}

type ProductStockUpdateRequest struct {
	ID            int64     `json:"id"`
	ProductID     int64     `json:"product_id" validate:"required"`
	StockQuantity int64     `json:"stock_quantity" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (productStockUpdateRequest ProductStockUpdateRequest) ToV1Domain() V1Domains.ProductStockDomain {
	return V1Domains.ProductStockDomain{
		ID:            productStockUpdateRequest.ID,
		ProductID:     productStockUpdateRequest.ProductID,
		StockQuantity: productStockUpdateRequest.StockQuantity,
		UpdatedAt:     productStockUpdateRequest.UpdatedAt,
	}
}
