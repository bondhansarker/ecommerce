package records

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"time"
)

type ProductStock struct {
	ID            int64     `db:"id"`
	ProductID     int64     `db:"product_id"`
	StockQuantity int64     `db:"stock_quantity"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func ToProductStockRecord(domain V1Domains.ProductStockDomain) ProductStock {
	return ProductStock{
		ID:            domain.ID,
		ProductID:     domain.ProductID,
		StockQuantity: domain.StockQuantity,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (model ProductStock) ToDomain() V1Domains.ProductStockDomain {
	return V1Domains.ProductStockDomain{
		ID:            model.ID,
		ProductID:     model.ProductID,
		StockQuantity: model.StockQuantity,
		UpdatedAt:     model.UpdatedAt,
	}
}

func ToArrayOfProductStockDomain(records *[]ProductStock) []V1Domains.ProductStockDomain {
	var result []V1Domains.ProductStockDomain

	for _, record := range *records {
		result = append(result, record.ToDomain())
	}

	return result
}
