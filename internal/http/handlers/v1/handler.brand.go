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

type BrandHandler struct {
	brandUseCase V1Domains.BrandUseCase
}

func NewBrandHandler(brandUseCase V1Domains.BrandUseCase) BrandHandler {
	return BrandHandler{
		brandUseCase: brandUseCase,
	}
}

func (brandHandler BrandHandler) Create(ctx *gin.Context) {
	var brandCreateRequest requests.BrandCreateRequest
	if err := ctx.ShouldBindJSON(&brandCreateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	inputDomain := brandCreateRequest.ToV1Domain()
	outputDomain, statusCode, err := brandHandler.brandUseCase.Create(ctx.Request.Context(), inputDomain)
	logger.Debug(outputDomain, statusCode, err)

	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "brand created successfully", map[string]interface{}{
		"brand": responses.FromV1DomainToBrandResponse(outputDomain),
	})
}

func (brandHandler BrandHandler) Get(ctx *gin.Context) {
	brandID := ctx.Param("id")
	if strings.TrimSpace(brandID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "brand id not found")
		return
	}

	brandIDInt, parseErr := strconv.ParseInt(brandID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}
	outputDomain, statusCode, err := brandHandler.brandUseCase.GetByID(ctx, brandIDInt)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	brandResponse := responses.FromV1DomainToBrandResponse(outputDomain)

	NewSuccessResponse(ctx, statusCode, "brand fetched successfully", map[string]interface{}{
		"brand": brandResponse,
	})

}

func (brandHandler BrandHandler) GetList(ctx *gin.Context) {
	//currentPage := ctx.Query("current_page")
	//itemPerPage := ctx.Query("item_per_page")
	//
	//var currentPageInt, itemPerPageInt int
	//var parseErr error
	//
	//if strings.TrimSpace(currentPage) != "" {
	//	currentPageInt, parseErr = strconv.Atoi(currentPage)
	//	if parseErr != nil {
	//		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
	//		return
	//	}
	//}
	//if strings.TrimSpace(itemPerPage) != "" {
	//	itemPerPageInt, parseErr = strconv.Atoi(itemPerPage)
	//	if parseErr != nil {
	//		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
	//		return
	//	}
	//}
	//
	//if currentPageInt <= 0 {
	//	currentPageInt = 1
	//}
	//
	//if itemPerPageInt <= 0 {
	//	itemPerPageInt = 5
	//}

	outputDomain, statusCode, err := brandHandler.brandUseCase.GetList(ctx)
	logger.Debug(outputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	brandListResponse := responses.ToResponseListOfBrands(outputDomain)

	NewSuccessResponse(ctx, statusCode, "brands fetched successfully", map[string]interface{}{
		"brands": brandListResponse,
	})

}

func (brandHandler BrandHandler) Update(ctx *gin.Context) {
	brandID := ctx.Param("id")
	if strings.TrimSpace(brandID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "brand id not found")
		return
	}

	brandIDInt, parseErr := strconv.ParseInt(brandID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	var brandUpdateRequest requests.BrandUpdateRequest

	if err := ctx.ShouldBindJSON(&brandUpdateRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	brandUpdateRequest.ID = brandIDInt
	inputDomain := brandUpdateRequest.ToV1Domain()

	statusCode, err := brandHandler.brandUseCase.Update(ctx.Request.Context(), inputDomain)
	logger.Debug(inputDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "brand updated successfully", nil)

}

func (brandHandler BrandHandler) Delete(ctx *gin.Context) {
	brandID := ctx.Param("id")
	if strings.TrimSpace(brandID) == "" {
		NewErrorResponse(ctx, http.StatusBadRequest, "brand id not found")
		return
	}

	brandIDInt, parseErr := strconv.ParseInt(brandID, 10, 64)
	if parseErr != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, parseErr.Error())
		return
	}

	statusCode, err := brandHandler.brandUseCase.Delete(ctx.Request.Context(), brandIDInt)
	logger.Debug(statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "brand deleted successfully", nil)

}
