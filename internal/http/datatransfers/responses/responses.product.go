package responses

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"time"
)

type ProductResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Specifications string    `json:"specifications"`
	BrandID        int64     `json:"brand_id"`
	CategoryID     int64     `json:"category_id"`
	SupplierID     int64     `json:"supplier_id"`
	UnitPrice      float64   `json:"unit_price"`
	DiscountPrice  float64   `json:"discount_price"`
	Tags           string    `json:"tags"`
	StatusID       int8      `json:"status_id"`
	StockQuantity  int64     `json:"stock_quantity,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

func (productResponse *ProductResponse) ToV1Domain() V1Domains.ProductDomain {
	return V1Domains.ProductDomain{
		ID:             productResponse.ID,
		Name:           productResponse.Name,
		Description:    productResponse.Description,
		Specifications: productResponse.Specifications,
		BrandID:        productResponse.BrandID,
		CategoryID:     productResponse.CategoryID,
		SupplierID:     productResponse.SupplierID,
		UnitPrice:      productResponse.UnitPrice,
		DiscountPrice:  productResponse.DiscountPrice,
		Tags:           productResponse.Tags,
		StatusID:       helpers.IntegerToBoolean[productResponse.StatusID],
		StockQuantity:  productResponse.StockQuantity,
		CreatedAt:      productResponse.CreatedAt,
	}
}

func FromV1DomainToProductResponse(productDomain V1Domains.ProductDomain) ProductResponse {
	return ProductResponse{
		ID:             productDomain.ID,
		Name:           productDomain.Name,
		Description:    productDomain.Description,
		Specifications: productDomain.Specifications,
		BrandID:        productDomain.BrandID,
		CategoryID:     productDomain.CategoryID,
		SupplierID:     productDomain.SupplierID,
		UnitPrice:      productDomain.UnitPrice,
		DiscountPrice:  productDomain.DiscountPrice,
		Tags:           productDomain.Tags,
		StatusID:       helpers.BooleanToInteger[*productDomain.StatusID],
		StockQuantity:  productDomain.StockQuantity,
		CreatedAt:      productDomain.CreatedAt,
	}
}

func ToResponseListOfProducts(domains []V1Domains.ProductDomain) []ProductResponse {
	var result = make([]ProductResponse, 0)

	for _, val := range domains {
		result = append(result, FromV1DomainToProductResponse(val))
	}

	return result
}
