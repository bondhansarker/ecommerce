package v1

import (
	"context"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"net/http"
)

type productStockUseCase struct {
	repo V1Domains.ProductStockRepository
}

func NewProductStockUseCase(repo V1Domains.ProductStockRepository) V1Domains.ProductStockUseCase {
	return productStockUseCase{
		repo: repo,
	}
}

func (p productStockUseCase) GetByID(ctx context.Context, id int64) (outputDomain V1Domains.ProductStockDomain, statusCode int, err error) {
	outputDomain, err = p.repo.GetRecordByID(ctx, id)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.ProductStockDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (p productStockUseCase) GetByProductID(ctx context.Context, productID int64) (outputDomain V1Domains.ProductStockDomain, statusCode int, err error) {
	outputDomain, err = p.repo.GetRecordByProductID(ctx, productID)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.ProductStockDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (p productStockUseCase) GetList(ctx context.Context) (outputDomains []V1Domains.ProductStockDomain, statusCode int, err error) {
	outputDomains, err = p.repo.GetRecords(ctx)
	if err != nil {
		return []V1Domains.ProductStockDomain{}, http.StatusInternalServerError, err
	}
	return outputDomains, http.StatusOK, nil
}

func (p productStockUseCase) Update(ctx context.Context, inputDomain V1Domains.ProductStockDomain) (statusCode int, err error) {
	err = p.repo.UpdateRecord(ctx, inputDomain)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}

func (p productStockUseCase) Delete(ctx context.Context, id int64) (statusCode int, err error) {
	err = p.repo.DeleteRecordByID(ctx, id)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}
