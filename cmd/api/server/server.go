package server

import (
	"context"
	"fmt"
	V1UseCase "github.com/bondhansarker/ecommerce/internal/business/usecases/v1"
	V1PostgresRepository "github.com/bondhansarker/ecommerce/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/bondhansarker/ecommerce/internal/http/handlers/v1"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bondhansarker/ecommerce/internal/config"
	"github.com/bondhansarker/ecommerce/internal/constants"
	"github.com/bondhansarker/ecommerce/internal/http/routes"
	"github.com/bondhansarker/ecommerce/internal/utils"
	"github.com/bondhansarker/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// setup databases
	utils.SetupPostgresConnection()

	dbClient := utils.GetDbClient()

	// Dependency Injection
	V1BrandRepository := V1PostgresRepository.NewBrandRepository(dbClient)
	V1BrandUseCase := V1UseCase.NewBrandUseCase(V1BrandRepository)
	V1BrandHandler := V1Handler.NewBrandHandler(V1BrandUseCase)

	V1CategoryRepository := V1PostgresRepository.NewCategoryRepository(dbClient)
	V1CategoryUseCase := V1UseCase.NewCategoryUseCase(V1CategoryRepository)
	V1CategoryHandler := V1Handler.NewCategoryHandler(V1CategoryUseCase)

	V1SupplierRepository := V1PostgresRepository.NewSupplierRepository(dbClient)
	V1SupplierUseCase := V1UseCase.NewSupplierUseCase(V1SupplierRepository)
	V1SupplierHandler := V1Handler.NewSupplierHandler(V1SupplierUseCase)

	V1ProductStockRepository := V1PostgresRepository.NewProductStockRepository(dbClient)
	V1ProductStockUseCase := V1UseCase.NewProductStockUseCase(V1ProductStockRepository)
	V1ProductStockHandler := V1Handler.NewProductStockHandler(V1ProductStockUseCase)

	V1ProductRepository := V1PostgresRepository.NewProductRepository(dbClient)
	V1ProductUseCase := V1UseCase.NewProductUseCase(V1ProductRepository, V1ProductStockRepository)
	V1ProductHandler := V1Handler.NewProductHandler(V1ProductUseCase)

	// setup router
	router := setupRouter()

	// API Routes
	api := router.Group("api")
	api.GET("/", routes.RootHandler)

	routes.NewBrandRoute(api, V1BrandHandler).Routes()
	routes.NewCategoryRoute(api, V1CategoryHandler).Routes()
	routes.NewSupplierRoute(api, V1SupplierHandler).Routes()
	routes.NewProductRoute(api, V1ProductHandler).Routes()
	routes.NewProductStockRoute(api, V1ProductStockHandler).Routes()

	// setup http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Gracefully Shutdown
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.AppConfig.Port)
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	return
}

func setupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.DebugMode
	//if config.AppConfig.Debug {
	//	mode = gin.DebugMode
	//}
	gin.SetMode(mode)

	// create a new router instance
	router := gin.New()

	// set up middlewares
	//router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	return router
}
