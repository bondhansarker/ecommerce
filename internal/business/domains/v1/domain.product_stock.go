package v1

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type ProductStockDomain struct {
	ID            int64     `json:"id"`
	ProductID     int64     `json:"product_id"`
	StockQuantity int64     `json:"stock_quantity"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductStockUseCase interface {
	GetByID(ctx context.Context, id int64) (outputDomain ProductStockDomain, statusCode int, err error)
	GetByProductID(ctx context.Context, id int64) (outputDomain ProductStockDomain, statusCode int, err error)
	GetList(ctx context.Context) (outputDomain []ProductStockDomain, statusCode int, err error)
	Update(ctx context.Context, inputDomain ProductStockDomain) (statusCode int, err error)
	Delete(ctx context.Context, id int64) (statusCode int, err error)
}

type ProductStockRepository interface {
	CreateRecordWithTransaction(tx *sqlx.Tx, inputDomain ProductStockDomain) (outputDomain ProductStockDomain, err error)
	GetRecords(ctx context.Context) (outputDomains []ProductStockDomain, err error)
	GetRecordByID(ctx context.Context, id int64) (outputDomain ProductStockDomain, err error)
	GetRecordByProductID(ctx context.Context, productID int64) (outputDomain ProductStockDomain, err error)
	UpdateRecord(ctx context.Context, inputDomain ProductStockDomain) (err error)
	DeleteRecordByID(ctx context.Context, id int64) (err error)
}
