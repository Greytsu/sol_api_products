package main

import (
	"fr/greytsu/sol_api_products/config"
	"fr/greytsu/sol_api_products/database"
	"fr/greytsu/sol_api_products/product"
	"fr/greytsu/sol_api_products/variant"
	"fr/greytsu/sol_api_products/warehouse"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	config.LoadEnvironmentVariables()
}

func main() {

	//Init database connection
	databaseCon := database.DatabaseCon{}
	err := databaseCon.ConnectToDatabaseWithRetry()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Exiting")
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
		log.Fatalf("Failed to start the server: %s", err.Error())
	}
}
