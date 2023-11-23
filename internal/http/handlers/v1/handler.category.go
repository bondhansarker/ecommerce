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

type CategoryHandler struct {
	categoryUseCase V1Domains.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase V1Domains.CategoryUseCase) CategoryHandler {
	return CategoryHandler{
		categoryUseCase: categoryUseCase,
	}
}

func (categoryHandler CategoryHandler) Create(ctx *gin.Context) {
	var categoryCreateRequest requests.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&categoryCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inputDomain := categoryCreateRequest.ToV1Domain()
	outputDomain, statusCode, err := categoryHandler.categoryUseCase.Create(ctx.Request.Context(), inputDomain)
	logger.Debug(outputDomain, statusCode, err)

	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "category created successfully", map[string]interface{}{
		"category": responses.FromV1DomainToCategoryResponse(outputDomain),
	})
}

func (categoryHandler CategoryHandler) Get(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	if strings.TrimSpace(categoryID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "category id not found")
		return
	}

	categoryIDInt, parseErr := strconv.ParseInt(categoryID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	outputDomain, statusCode, err := categoryHandler.categoryUseCase.GetByID(ctx, categoryIDInt)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	categoryResponse := responses.FromV1DomainToCategoryResponse(outputDomain)

	NewSuccessResponse(ctx, statusCode, "category fetched successfully", map[string]interface{}{
		"category": categoryResponse,
	})
}

func (categoryHandler CategoryHandler) GetList(ctx *gin.Context) {
	outputDomain, statusCode, err := categoryHandler.categoryUseCase.GetList(ctx)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	categoryListResponse := responses.ToResponseListOfCategories(outputDomain)

	NewSuccessResponse(ctx, statusCode, "categories fetched successfully", map[string]interface{}{
		"categories": categoryListResponse,
	})
}

func (categoryHandler CategoryHandler) GetTree(ctx *gin.Context) {
	outputDomain, statusCode, err := categoryHandler.categoryUseCase.GetList(ctx)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}
	categoryListResponse := responses.ToTreeListOfCategories(outputDomain)

	NewSuccessResponse(ctx, statusCode, "categories fetched successfully", map[string]interface{}{
		"categories": categoryListResponse,
	})
}

func (categoryHandler CategoryHandler) Update(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	if strings.TrimSpace(categoryID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "category id not found")
		return
	}

	categoryIDInt, parseErr := strconv.ParseInt(categoryID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	var categoryUpdateRequest requests.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&categoryUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	categoryUpdateRequest.ID = categoryIDInt
	inputDomain := categoryUpdateRequest.ToV1Domain()

	statusCode, err := categoryHandler.categoryUseCase.Update(ctx.Request.Context(), inputDomain)
	logger.Debug(inputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "category updated successfully", nil)
}

func (categoryHandler CategoryHandler) Delete(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	if strings.TrimSpace(categoryID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "category id not found")
		return
	}

	categoryIDInt, parseErr := strconv.ParseInt(categoryID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	statusCode, err := categoryHandler.categoryUseCase.Delete(ctx.Request.Context(), categoryIDInt)
	logger.Debug(statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "category deleted successfully", nil)
}
