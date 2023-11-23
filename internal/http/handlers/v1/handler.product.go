package v1

import (
	"github.com/bondhansarker/ecommerce/internal/business/domains/v1"
	"github.com/bondhansarker/ecommerce/internal/http/datatransfers/requests"
	"github.com/bondhansarker/ecommerce/internal/http/datatransfers/responses"
	"github.com/bondhansarker/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productUseCase v1.ProductUseCase
}

func NewProductHandler(productUseCase v1.ProductUseCase) ProductHandler {
	return ProductHandler{
		productUseCase: productUseCase,
	}
}

func (productHandler ProductHandler) Create(ctx *gin.Context) {
	var productCreateRequest requests.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&productCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inputDomain := productCreateRequest.ToV1Domain()
	outputDomain, statusCode, err := productHandler.productUseCase.Create(ctx.Request.Context(), inputDomain)
	logger.Debug(outputDomain, statusCode, err)

	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "product created successfully", map[string]interface{}{
		"product": responses.FromV1DomainToProductResponse(outputDomain),
	})
}

func (productHandler ProductHandler) Get(ctx *gin.Context) {
	productID := ctx.Param("id")
	if strings.TrimSpace(productID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product id not found")
		return
	}

	productIDInt, parseErr := strconv.ParseInt(productID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}
	outputDomain, statusCode, err := productHandler.productUseCase.GetByID(ctx, productIDInt)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	productResponse := responses.FromV1DomainToProductResponse(outputDomain)

	NewSuccessResponse(ctx, statusCode, "product fetched successfully", map[string]interface{}{
		"product": productResponse,
	})
}

func (productHandler ProductHandler) GetList(ctx *gin.Context) {
	var productQueryParams requests.ProductQueryParams
	// Bind query parameters to the struct
	if err := ctx.ShouldBindQuery(&productQueryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outputDomain, statusCode, err := productHandler.productUseCase.GetList(ctx, productQueryParams.ToV1ProductFilterParams())
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	productListResponse := responses.ToResponseListOfProducts(outputDomain)

	NewSuccessResponse(ctx, statusCode, "products fetched successfully", map[string]interface{}{
		"products": productListResponse,
	})
}

func (productHandler ProductHandler) Update(ctx *gin.Context) {
	productID := ctx.Param("id")
	if strings.TrimSpace(productID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product id not found")
		return
	}

	productIDInt, parseErr := strconv.ParseInt(productID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	var productUpdateRequest requests.ProductUpdateRequest

	if err := ctx.ShouldBindJSON(&productUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	productUpdateRequest.ID = productIDInt
	inputDomain := productUpdateRequest.ToV1Domain()

	statusCode, err := productHandler.productUseCase.Update(ctx.Request.Context(), inputDomain)
	logger.Debug(inputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "product updated successfully", nil)
}

func (productHandler ProductHandler) Delete(ctx *gin.Context) {
	productID := ctx.Param("id")
	if strings.TrimSpace(productID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "product id not found")
		return
	}

	productIDInt, parseErr := strconv.ParseInt(productID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	statusCode, err := productHandler.productUseCase.Delete(ctx.Request.Context(), productIDInt)
	logger.Debug(statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "product deleted successfully", nil)
}
