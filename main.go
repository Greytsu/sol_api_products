package main

import (
	"fr/greytsu/sol_api_products/config"
	"fr/greytsu/sol_api_products/database"
	"fr/greytsu/sol_api_products/product"
	"fr/greytsu/sol_api_products/stock"
	"fr/greytsu/sol_api_products/variant"
	"fr/greytsu/sol_api_products/warehouse"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func init() {
	config.LoadEnvironmentVariables()
	config.SetupLogger()
}

func main() {

	//Init database connection
	databaseCon := database.DatabaseCon{}
	err := databaseCon.ConnectToDatabaseWithRetry()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database. Exiting")
	}

	//Init product
	productRepository := product.NewProductRepository(databaseCon.GetDatabaseCon())
	productService := product.NewProductService(productRepository)

	//Init stock
	stockRepository := stock.NewStockRepository(databaseCon.GetDatabaseCon())
	stockService := stock.NewStockService(stockRepository)

	//Init variant
	variantRepository := variant.NewVariantRepository(databaseCon.GetDatabaseCon())
	variantService := variant.NewVariantService(variantRepository, stockService)

	//Init warehouse
	warehouseRepository := warehouse.NewWarehouseRepository(databaseCon.GetDatabaseCon())
	warehouseService := warehouse.NewWarehouseService(warehouseRepository)

	//Create the Gin router
	router := gin.Default()
	v1 := router.Group("/api/v1")

	//Routes
	product.RegisterProductsRoutes(v1, productService, variantService)
	warehouse.RegisterWarehousesRoutes(v1, warehouseService)
	variant.RegisterVariantsRoutes(v1, variantService)

	err = router.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server. Exiting")
	}
}
