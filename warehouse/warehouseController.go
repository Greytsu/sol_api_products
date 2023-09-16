package warehouse

import (
	"fr/greytsu/sol_api_products/models"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterWarehousesRoutes(routerGroup *gin.RouterGroup, warehouseService *WarehouseService) {
	routerGroup.GET("/warehouses", getAllWarehouses(warehouseService))
	routerGroup.POST("/warehouses", postWarehouse(warehouseService))
	routerGroup.PUT("/warehouses/:id", putWarehouse(warehouseService))
	routerGroup.DELETE("/warehouses/:id", deleteWarehouse(warehouseService))
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
		warehouse, err := warehouseService.CreateWarehouse(&newWarehouse)
		if err != nil {
			if strings.Contains(err.Error(), "Warehouse already exists") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating warehouse")
			return
		}
		c.IndentedJSON(http.StatusOK, warehouse)
	}
}

func putWarehouse(warehouseService *WarehouseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid warehouse id")
			return
		}

		var updateWarehouse models.Warehouse
		if err := c.BindJSON(&updateWarehouse); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		warehouse, err := warehouseService.UpdateWarehouse(id, companyId, &updateWarehouse)
		if err != nil {
			if strings.Contains(err.Error(), "Warehouse not found") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while updating warehouse")
			return
		}
		c.IndentedJSON(http.StatusOK, warehouse)
	}
}

func deleteWarehouse(warehouseService *WarehouseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid id")
			return
		}

		err = warehouseService.DeleteWarehouse(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Warehouse not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while deleting warehouse")
			return
		}
		c.IndentedJSON(http.StatusNoContent, "")
	}
}
