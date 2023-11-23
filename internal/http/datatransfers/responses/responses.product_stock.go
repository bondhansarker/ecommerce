package responses

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

type ProductStockResponse struct {
	ID            int64     `json:"id"`
	ProductID     int64     `json:"product_id"`
	StockQuantity int64     `json:"stock_quantity"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (productStockResponse *ProductStockResponse) ToV1Domain() V1Domains.ProductStockDomain {
	return V1Domains.ProductStockDomain{
		ID:            productStockResponse.ID,
		ProductID:     productStockResponse.ProductID,
		StockQuantity: productStockResponse.StockQuantity,
		UpdatedAt:     productStockResponse.UpdatedAt,
	}
}

func FromV1DomainToProductStockResponse(productStockDomain V1Domains.ProductStockDomain) ProductStockResponse {
	return ProductStockResponse{
		ID:            productStockDomain.ID,
		ProductID:     productStockDomain.ProductID,
		StockQuantity: productStockDomain.StockQuantity,
		UpdatedAt:     productStockDomain.UpdatedAt,
	}
}

func ToResponseListOfProductStocks(domains []V1Domains.ProductStockDomain) []ProductStockResponse {
	var result = make([]ProductStockResponse, 0)

	for _, val := range domains {
		result = append(result, FromV1DomainToProductStockResponse(val))
	}

	return result
}
