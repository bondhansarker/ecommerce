package v1

import (
	"context"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"net/http"
)

type brandUseCase struct {
	repo V1Domains.BrandRepository
}

func NewBrandUseCase(repo V1Domains.BrandRepository) V1Domains.BrandUseCase {
	return brandUseCase{
		repo: repo,
	}
}

func (b brandUseCase) Create(ctx context.Context, inputDomain V1Domains.BrandDomain) (outputDomain V1Domains.BrandDomain, statusCode int, err error) {
	outputDomain, err = b.repo.CreateRecord(ctx, inputDomain)
	if err != nil {
		return V1Domains.BrandDomain{}, http.StatusInternalServerError, err
	}

	return outputDomain, http.StatusCreated, nil

}

func (b brandUseCase) GetByID(ctx context.Context, id int64) (outputDomain V1Domains.BrandDomain, statusCode int, err error) {
	outputDomain, err = b.repo.GetRecordByID(ctx, id)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.BrandDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (b brandUseCase) GetList(ctx context.Context) (outputDomains []V1Domains.BrandDomain, statusCode int, err error) {
	outputDomains, err = b.repo.GetRecords(ctx)
	if err != nil {
		return []V1Domains.BrandDomain{}, http.StatusInternalServerError, err
	}
	return outputDomains, http.StatusOK, nil
}

func (b brandUseCase) Update(ctx context.Context, inputDomain V1Domains.BrandDomain) (statusCode int, err error) {
	err = b.repo.UpdateRecord(ctx, inputDomain)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}

func (b brandUseCase) Delete(ctx context.Context, id int64) (statusCode int, err error) {
	err = b.repo.DeleteRecordByID(ctx, id)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}
