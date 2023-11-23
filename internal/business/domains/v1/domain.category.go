package v1

import (
	"context"
	"time"
)

type CategoryDomain struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ParentID  *int64    `json:"parent_id"`
	Sequence  uint      `json:"sequence"`
	StatusID  *bool     `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryUseCase interface {
	Create(ctx context.Context, inputDomain CategoryDomain) (outputDomain CategoryDomain, statusCode int, err error)
	GetByID(ctx context.Context, id int64) (outputDomain CategoryDomain, statusCode int, err error)
	GetList(ctx context.Context) (outputDomain []CategoryDomain, statusCode int, err error)
	Update(ctx context.Context, inputDomain CategoryDomain) (statusCode int, err error)
	Delete(ctx context.Context, id int64) (statusCode int, err error)
}

type CategoryRepository interface {
	CreateRecord(ctx context.Context, inputDomain CategoryDomain) (outputDomain CategoryDomain, err error)
	GetRecords(ctx context.Context) (outputDomains []CategoryDomain, err error)
	GetRecordByID(ctx context.Context, id int64) (outputDomain CategoryDomain, err error)
	UpdateRecord(ctx context.Context, inputDomain CategoryDomain) (err error)
	DeleteRecordByID(ctx context.Context, id int64) (err error)
}
