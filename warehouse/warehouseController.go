package warehouse

import (
	"fr/greytsu/sol_api_products/models"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterProductRoutes(routerGroup *gin.RouterGroup, warehouseService *WarehouseService) {
	routerGroup.GET("/warehouses", getAllWarehouses(warehouseService))
	routerGroup.POST("/warehouses", postWarehouse(warehouseService))
}

func getAllWarehouses(warehouseService *WarehouseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
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
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}

		var newWarehouse models.Warehouse
		if err := c.BindJSON(&newWarehouse); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		newWarehouse.CompanyID = companyId
		product, err := warehouseService.CreateWarehouse(&newWarehouse)
		if err != nil {
			if strings.Contains(err.Error(), "Warehouse already exists") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating warehouse")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}
