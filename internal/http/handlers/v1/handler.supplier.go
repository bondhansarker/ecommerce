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

type SupplierHandler struct {
	supplierUseCase V1Domains.SupplierUseCase
}

// NewSupplierHandler creates a new instance of SupplierHandler
func NewSupplierHandler(supplierUseCase V1Domains.SupplierUseCase) SupplierHandler {
	return SupplierHandler{
		supplierUseCase: supplierUseCase,
	}
}

func (supplierHandler SupplierHandler) Create(ctx *gin.Context) {
	var supplierCreateRequest requests.SupplierCreateRequest
	if err := ctx.ShouldBindJSON(&supplierCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inputDomain := supplierCreateRequest.ToV1Domain()
	outputDomain, statusCode, err := supplierHandler.supplierUseCase.Create(ctx.Request.Context(), inputDomain)
	logger.Debug(outputDomain, statusCode, err)

	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "supplier created successfully", map[string]interface{}{
		"supplier": responses.FromV1DomainToSupplierResponse(outputDomain),
	})
}

func (supplierHandler SupplierHandler) Get(ctx *gin.Context) {
	supplierID := ctx.Param("id")
	if strings.TrimSpace(supplierID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "supplier id not found")
		return
	}

	supplierIDInt, parseErr := strconv.ParseInt(supplierID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}
	outputDomain, statusCode, err := supplierHandler.supplierUseCase.GetByID(ctx, supplierIDInt)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	supplierResponse := responses.FromV1DomainToSupplierResponse(outputDomain)

	NewSuccessResponse(ctx, statusCode, "supplier fetched successfully", map[string]interface{}{
		"supplier": supplierResponse,
	})
}

func (supplierHandler SupplierHandler) GetList(ctx *gin.Context) {
	outputDomain, statusCode, err := supplierHandler.supplierUseCase.GetList(ctx)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	supplierListResponse := responses.ToResponseListOfSuppliers(outputDomain)

	NewSuccessResponse(ctx, statusCode, "suppliers fetched successfully", map[string]interface{}{
		"suppliers": supplierListResponse,
	})
}

func (supplierHandler SupplierHandler) Update(ctx *gin.Context) {
	supplierID := ctx.Param("id")
	if strings.TrimSpace(supplierID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "supplier id not found")
		return
	}

	supplierIDInt, parseErr := strconv.ParseInt(supplierID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	var supplierUpdateRequest requests.SupplierUpdateRequest

	if err := ctx.ShouldBindJSON(&supplierUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	supplierUpdateRequest.ID = supplierIDInt
	inputDomain := supplierUpdateRequest.ToV1Domain()

	statusCode, err := supplierHandler.supplierUseCase.Update(ctx.Request.Context(), inputDomain)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "supplier updated successfully", nil)
}

func (supplierHandler SupplierHandler) Delete(ctx *gin.Context) {
	supplierID := ctx.Param("id")
	if strings.TrimSpace(supplierID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "supplier id not found")
		return
	}

	supplierIDInt, parseErr := strconv.ParseInt(supplierID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	statusCode, err := supplierHandler.supplierUseCase.Delete(ctx.Request.Context(), supplierIDInt)
	logger.Debug(statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "supplier deleted successfully", nil)
}
