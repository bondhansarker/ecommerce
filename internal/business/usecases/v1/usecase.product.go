package v1

import (
	"context"
	"net/http"

	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/utils"
	"github.com/bondhansarker/ecommerce/pkg/helpers"
)

type productUseCase struct {
	productRepo      V1Domains.ProductRepository
	productStockRepo V1Domains.ProductStockRepository
}

func NewProductUseCase(productRepo V1Domains.ProductRepository, productStockRepo V1Domains.ProductStockRepository) V1Domains.ProductUseCase {
	return productUseCase{
		productRepo:      productRepo,
		productStockRepo: productStockRepo,
	}
}

func (p productUseCase) Create(ctx context.Context, inputProductDomain V1Domains.ProductDomain) (outputProductDomain V1Domains.ProductDomain, statusCode int, err error) {
	// Start a new transaction
	tx, err := utils.BeginTransaction(ctx)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	// Defer rollback if an error occurs
	defer func() {
		if err != nil {
			utils.RollbackTransaction(tx)
		}
	}()

	// Create the product
	outputProductDomain, err = p.productRepo.CreateRecordWithTransaction(tx, inputProductDomain)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	// assign product stock
	productStockDomain := V1Domains.ProductStockDomain{
		ProductID:     outputProductDomain.ID,
		StockQuantity: inputProductDomain.StockQuantity,
	}

	// Create the stock for the product
	_, err = p.productStockRepo.CreateRecordWithTransaction(tx, productStockDomain)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}

	// Commit the transaction if everything is successful
	err = utils.CommitTransaction(tx)
	if err != nil {
		return V1Domains.ProductDomain{}, http.StatusInternalServerError, err
	}
	outputProductDomain.StockQuantity = productStockDomain.StockQuantity

	return outputProductDomain, http.StatusCreated, nil
}

func (p productUseCase) GetByID(ctx context.Context, id int64) (outputDomain V1Domains.ProductDomain, statusCode int, err error) {
	outputDomain, err = p.productRepo.GetRecordByID(ctx, id)
	if err != nil {
		status, err := helpers.HandleCommonRepositoryError(err)
		return V1Domains.ProductDomain{}, status, err
	}
	return outputDomain, http.StatusOK, nil
}

func (p productUseCase) GetList(ctx context.Context, productFilterParams V1Domains.ProductFilterParams, currentPageInt int, itemPerPageInt int) (outputDomains []V1Domains.ProductDomain, paginationResult *helpers.PaginationResult, statusCode int, err error) {
	outputDomains, paginationResult, err = p.productRepo.GetRecords(ctx, productFilterParams, currentPageInt, itemPerPageInt)
	if err != nil {
		return []V1Domains.ProductDomain{}, paginationResult, http.StatusInternalServerError, err
	}
	return outputDomains, paginationResult, http.StatusOK, nil
}

func (p productUseCase) Update(ctx context.Context, inputDomain V1Domains.ProductDomain) (statusCode int, err error) {
	err = p.productRepo.UpdateRecord(ctx, inputDomain)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}

func (p productUseCase) Delete(ctx context.Context, id int64) (statusCode int, err error) {
	err = p.productRepo.DeleteRecordByID(ctx, id)
	if err != nil {
		return helpers.HandleCommonRepositoryError(err)
	}
	return http.StatusOK, nil
}
