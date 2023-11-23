package v1

import (
	"context"
	"time"
)

type SupplierDomain struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	StatusID           *bool     `json:"status_id"`
	IsVerifiedSupplier *bool     `json:"is_verified_supplier"`
	CreatedAt          time.Time `json:"created_at"`
}

type SupplierUseCase interface {
	Create(ctx context.Context, inputDomain SupplierDomain) (outputDomain SupplierDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int64) (outputDomain SupplierDomain, statusCode int, err error)
	GetList(ctx context.Context) (outputDomain []SupplierDomain, statusCode int, err error)
	Update(ctx context.Context, inputDomain SupplierDomain) (statusCode int, err error)
	Delete(ctx context.Context, id int64) (statusCode int, err error)
}

type SupplierRepository interface {
	CreateRecord(ctx context.Context, inputDomain SupplierDomain) (outputDomain SupplierDomain, err error)
	GetRecords(ctx context.Context) (outputDomains []SupplierDomain, err error)
	GetRecordByID(ctx context.Context, id int64) (outputDomain SupplierDomain, err error)
	UpdateRecord(ctx context.Context, inputDomain SupplierDomain) (err error)
	DeleteRecordByID(ctx context.Context, id int64) (err error)
}
