package warehouse

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *gin.Engine, warehouseService *WarehouseService) {
	router.GET("/warehouses", getAllWarehouses(warehouseService))
	router.POST("/warehouses", postWarehouse(warehouseService))
}

func getAllWarehouses(warehouseService *WarehouseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Query("company_id")

		warehouses, err := warehouseService.GetAllWarehouses(companyId)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, "error")
			return
		}
		c.IndentedJSON(http.StatusOK, warehouses)
	}
}

func postWarehouse(warehouseService *WarehouseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newWarehouse models.Warehouse
		if err := c.BindJSON(&newWarehouse); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		product, err := warehouseService.CreateWarehouse(&newWarehouse)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating warehouse")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}
