package routes

import (
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
)

type brandRoutes struct {
	V1Handler V1Handler.BrandHandler
	router    *gin.RouterGroup
}

func NewBrandRoute(router *gin.RouterGroup, V1Handler V1Handler.BrandHandler) *brandRoutes {
	return &brandRoutes{V1Handler: V1Handler, router: router}
}

func (r *brandRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		V1BrandRoute := V1Route.Group("/brands")
		V1BrandRoute.POST("/", r.V1Handler.Create)
		V1BrandRoute.GET("/", r.V1Handler.GetList)
		V1BrandRoute.GET("/:id", r.V1Handler.Get)
		V1BrandRoute.PATCH("/:id", r.V1Handler.Update)
		V1BrandRoute.DELETE("/:id", r.V1Handler.Delete)
	}

}
