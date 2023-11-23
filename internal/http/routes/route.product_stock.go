package routes

import (
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
)

type productStockRoutes struct {
	V1Handler V1Handler.ProductStockHandler
	router    *gin.RouterGroup
}

func NewProductStockRoute(router *gin.RouterGroup, V1Handler V1Handler.ProductStockHandler) *productStockRoutes {
	return &productStockRoutes{V1Handler: V1Handler, router: router}
}

func (r *productStockRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// product_stock
		V1ProductStockRoute := V1Route.Group("/product_stocks")
		V1ProductStockRoute.GET("/", r.V1Handler.GetList)
		V1ProductStockRoute.GET("/:id", r.V1Handler.Get)
		V1ProductStockRoute.PATCH("/:id", r.V1Handler.Update)
		V1ProductStockRoute.DELETE("/:id", r.V1Handler.Delete)
	}
}
