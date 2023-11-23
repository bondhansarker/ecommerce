package routes

import (
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
)

type supplierRoutes struct {
	V1Handler V1Handler.SupplierHandler
	router    *gin.RouterGroup
}

func NewSupplierRoute(router *gin.RouterGroup, V1Handler V1Handler.SupplierHandler) *supplierRoutes {
	return &supplierRoutes{V1Handler: V1Handler, router: router}
}

func (r *supplierRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// suppliers
		V1SupplierRoute := V1Route.Group("/suppliers")
		V1SupplierRoute.POST("/", r.V1Handler.Create)
		V1SupplierRoute.GET("/", r.V1Handler.GetList)
		V1SupplierRoute.GET("/:id", r.V1Handler.Get)
		V1SupplierRoute.PATCH("/:id", r.V1Handler.Update)
		V1SupplierRoute.DELETE("/:id", r.V1Handler.Delete)
	}
}
