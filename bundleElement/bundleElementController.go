package bundleElement

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterBundleElementsRoutes(routerGroup *gin.RouterGroup, bundleElementService *BundleElementService) {
	routerGroup.DELETE("/bundle-elements/:id", deleteBundleElement(bundleElementService))
}

func deleteBundleElement(bundleElementService *BundleElementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid id")
			return
		}

		err = bundleElementService.DeleteBundleElement(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Bundle element not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while deleting bundle")
			return
		}
		c.IndentedJSON(http.StatusNoContent, "")
	}
}
