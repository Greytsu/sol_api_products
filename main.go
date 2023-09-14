package main

import (
	"fr/greytsu/sol_api_products/config"
	"fr/greytsu/sol_api_products/database"
	"fr/greytsu/sol_api_products/product"
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

	//Init variant
	variantRepository := variant.NewVariantRepository(databaseCon.GetDatabaseCon())
	variantService := variant.NewVariantService(variantRepository)

	//Init warehouse
	warehouseRepository := warehouse.NewWarehouseRepository(databaseCon.GetDatabaseCon())
	warehouseService := warehouse.NewWarehouseService(warehouseRepository)

	//Create the Gin router
	router := gin.Default()

	//Routes
	product.RegisterProductRoutes(router, productService, variantService)
	warehouse.RegisterProductRoutes(router, warehouseService)

	err = router.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server. Exiting")
	}
}
