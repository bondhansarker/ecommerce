package v1

import (
	"context"
	"time"
)

type BrandDomain struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StatusID  *bool     `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
}

type BrandUseCase interface {
	Create(ctx context.Context, inputDomain BrandDomain) (outputDomain BrandDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int64) (outputDomain BrandDomain, statusCode int, err error)
	GetList(ctx context.Context) (outputDomain []BrandDomain, statusCode int, err error)
	//GetList(ctx context.Context, currentPageInt int, itemPerPageInt int) (outputDomain []BrandDomain, statusCode int, err error)
	Update(ctx context.Context, inputDomain BrandDomain) (statusCode int, err error)
	Delete(ctx context.Context, id int64) (statusCode int, err error)
}

type BrandRepository interface {
	CreateRecord(ctx context.Context, inputDomain BrandDomain) (outputDomain BrandDomain, err error)
	GetRecords(ctx context.Context) (outputDomains []BrandDomain, err error)
	GetRecordByID(ctx context.Context, id int64) (outputDomain BrandDomain, err error)
	UpdateRecord(ctx context.Context, inputDomain BrandDomain) (err error)
	DeleteRecordByID(ctx context.Context, id int64) (err error)
}
