package v1

import (
	"context"
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
	"net/http"
)

type categoryUseCase struct {
	repo V1Domains.CategoryRepository
}

func NewCategoryUseCase(repo V1Domains.CategoryRepository) V1Domains.CategoryUseCase {
	return categoryUseCase{
		repo: repo,
	}
}

func (c categoryUseCase) Create(ctx context.Context, inputDomain V1Domains.CategoryDomain) (outputDomain V1Domains.CategoryDomain, statusCode int, err error) {
	outputDomain, err = c.repo.CreateRecord(ctx, inputDomain)
	if err != nil {
		return V1Domains.CategoryDomain{}, http.StatusInternalServerError, err
	}

	return outputDomain, http.StatusCreated, nil
}

func (c categoryUseCase) GetByID(ctx context.Context, id int64) (outputDomain V1Domains.CategoryDomain, statusCode int, err error) {
	outputDomain, err = c.repo.GetRecordByID(ctx, id)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.CategoryDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (c categoryUseCase) GetList(ctx context.Context) (outputDomains []V1Domains.CategoryDomain, statusCode int, err error) {
	outputDomains, err = c.repo.GetRecords(ctx)
	if err != nil {
		return []V1Domains.CategoryDomain{}, http.StatusInternalServerError, err
	}
	return outputDomains, http.StatusOK, nil
}

func (c categoryUseCase) Update(ctx context.Context, inputDomain V1Domains.CategoryDomain) (statusCode int, err error) {
	err = c.repo.UpdateRecord(ctx, inputDomain)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}

func (c categoryUseCase) Delete(ctx context.Context, id int64) (statusCode int, err error) {
	err = c.repo.DeleteRecordByID(ctx, id)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}
