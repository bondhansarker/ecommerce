package v1

import (
	"context"
	"time"

	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type ProductDomain struct {
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
	StatusID       *bool     `json:"status_id"`
	StockQuantity  int64     `json:"stock_quantity"`
	CreatedAt      time.Time `json:"created_at"`
}

type ProductFilterParams struct {
	Name               string   `json:"name,omitempty"`                 // Search term for the product name
	MinPrice           *float64 `json:"min_price,omitempty"`            // Minimum unit price
	MaxPrice           *float64 `json:"max_price,omitempty"`            // Maximum unit price
	BrandIDs           []int64  `json:"brand_ids,omitempty"`            // Brand IDs to filter by
	CategoryID         *int64   `json:"category_id,omitempty"`          // Category ID to filter by
	SupplierID         *int64   `json:"supplier_id,omitempty"`          // Supplier ID to filter by
	IsVerifiedSupplier *bool    `json:"is_verified_supplier,omitempty"` // Flag to filter by verified supplier
}

type ProductUseCase interface {
	Create(ctx context.Context, inputDomain ProductDomain) (outputDomain ProductDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int64) (outputDomain ProductDomain, statusCode int, err error)
	GetList(ctx context.Context, productFilterParams ProductFilterParams, currentPageInt int, itemPerPageInt int) (outputDomain []ProductDomain, paginationResult *helpers.PaginationResult, statusCode int, err error)
	Update(ctx context.Context, inputDomain ProductDomain) (statusCode int, err error)
	Delete(ctx context.Context, id int64) (statusCode int, err error)
}

type ProductRepository interface {
	CreateRecord(ctx context.Context, inputDomain ProductDomain) (outputDomain ProductDomain, err error)
	CreateRecordWithTransaction(tx *sqlx.Tx, inputDomain ProductDomain) (outputDomain ProductDomain, err error)
	GetRecords(ctx context.Context, productFilterParams ProductFilterParams, currentPageInt int, itemPerPageInt int) (outputDomains []ProductDomain, paginationResult *helpers.PaginationResult, err error)
	GetRecordByID(ctx context.Context, id int64) (outputDomain ProductDomain, err error)
	UpdateRecord(ctx context.Context, inputDomain ProductDomain) (err error)
	DeleteRecordByID(ctx context.Context, id int64) (err error)
}
