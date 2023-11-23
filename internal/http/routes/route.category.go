package routes

import (
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
)

type categoryRoutes struct {
	V1Handler V1Handler.CategoryHandler
	router    *gin.RouterGroup
}

func NewCategoryRoute(router *gin.RouterGroup, V1Handler V1Handler.CategoryHandler) *categoryRoutes {
	return &categoryRoutes{V1Handler: V1Handler, router: router}
}

func (r *categoryRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// category
		V1CategoryRoute := V1Route.Group("/categories")
		V1CategoryRoute.POST("/", r.V1Handler.Create)
		V1CategoryRoute.GET("/", r.V1Handler.GetTree)
		V1CategoryRoute.GET("/list", r.V1Handler.GetList)
		V1CategoryRoute.GET("/:id", r.V1Handler.Get)
		V1CategoryRoute.PATCH("/:id", r.V1Handler.Update)
		V1CategoryRoute.DELETE("/:id", r.V1Handler.Delete)
	}
}
