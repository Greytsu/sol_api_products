package main

import (
	"context"
	"encoding/json"
	"fmt"
	"fr/greytsu/sol_api_products/config"
	"fr/greytsu/sol_api_products/database"
	"fr/greytsu/sol_api_products/models"
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

	product, err := models.PRProducts().One(context.Background(), databaseCon.GetDatabaseCon())
	fmt.Println(product)
	options, err := product.OpFKProductOpOptions().All(context.Background(), databaseCon.GetDatabaseCon())
	for i := 0; i < len(options); i++ {
		jsonOption, _ := json.Marshal(options[i])
		fmt.Println(string(jsonOption))
	}
}
