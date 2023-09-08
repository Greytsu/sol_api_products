package warehouse

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *gin.Engine, warehouseService *WarehouseService) {
	router.GET("/warehouses", getAllWarehouses(warehouseService))
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
