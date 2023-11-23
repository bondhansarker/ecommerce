package v1

import (
	V1Domains "github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/http/datatransfers/requests"
	"github.com/bondhansarker/ecommerce/internal/http/datatransfers/responses"
	"github.com/bondhansarker/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ProductStockHandler struct {
	productStockUseCase V1Domains.ProductStockUseCase
}

func NewProductStockHandler(productStockUseCase V1Domains.ProductStockUseCase) ProductStockHandler {
	return ProductStockHandler{
		productStockUseCase: productStockUseCase,
	}
}

func (p ProductStockHandler) Get(ctx *gin.Context) {
	productStockID := ctx.Param("id")
	if strings.TrimSpace(productStockID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product stock id not found")
		return
	}

	productStockIDInt, parseErr := strconv.ParseInt(productStockID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}
	outputDomain, statusCode, err := p.productStockUseCase.GetByID(ctx, productStockIDInt)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	productStockResponse := responses.FromV1DomainToProductStockResponse(outputDomain)

	NewSuccessResponse(ctx, statusCode, "product stock fetched successfully", map[string]interface{}{
		"product_stock": productStockResponse,
	})

}

func (p ProductStockHandler) GetList(ctx *gin.Context) {
	outputDomains, statusCode, err := p.productStockUseCase.GetList(ctx)
	logger.Debug(outputDomains, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	productStockListResponse := responses.ToResponseListOfProductStocks(outputDomains)

	NewSuccessResponse(ctx, statusCode, "product stocks fetched successfully", map[string]interface{}{
		"product_stocks": productStockListResponse,
	})

}

func (p ProductStockHandler) Update(ctx *gin.Context) {
	productStockID := ctx.Param("id")
	if strings.TrimSpace(productStockID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product stock id not found")
		return
	}

	productStockIDInt, parseErr := strconv.ParseInt(productStockID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	var productStockUpdateRequest requests.ProductStockUpdateRequest

	if err := ctx.ShouldBindJSON(&productStockUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	productStockUpdateRequest.ID = productStockIDInt
	inputDomain := productStockUpdateRequest.ToV1Domain()

	statusCode, err := p.productStockUseCase.Update(ctx.Request.Context(), inputDomain)
	logger.Debug(inputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "product stock updated successfully", nil)

}

func (p ProductStockHandler) Delete(ctx *gin.Context) {
	productStockID := ctx.Param("id")
	if strings.TrimSpace(productStockID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product stock id not found")
		return
	}

	productStockIDInt, parseErr := strconv.ParseInt(productStockID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	statusCode, err := p.productStockUseCase.Delete(ctx.Request.Context(), productStockIDInt)
	logger.Debug(statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "product stock deleted successfully", nil)

}
