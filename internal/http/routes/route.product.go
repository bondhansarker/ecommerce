package routes

import (
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
)

type productRoutes struct {
	V1Handler V1Handler.ProductHandler
	router    *gin.RouterGroup
}

func NewProductRoute(router *gin.RouterGroup, V1Handler V1Handler.ProductHandler) *productRoutes {
	return &productRoutes{V1Handler: V1Handler, router: router}
}

func (r *productRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// product
		V1ProductRoute := V1Route.Group("/products")
		V1ProductRoute.POST("/", r.V1Handler.Create)
		V1ProductRoute.GET("/", r.V1Handler.GetList)
		V1ProductRoute.GET("/:id", r.V1Handler.Get)
		V1ProductRoute.PATCH("/:id", r.V1Handler.Update)
		V1ProductRoute.DELETE("/:id", r.V1Handler.Delete)
	}

}
