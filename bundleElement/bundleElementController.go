package bundleElement

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterBundleElementsRoutes(routerGroup *gin.RouterGroup, bundleElementService *BundleElementService) {
	routerGroup.GET("/bundle-elements", getBundleElementByBundleAndVariant(bundleElementService))
	routerGroup.DELETE("/bundle-elements/:id", deleteBundleElement(bundleElementService))
}

func getBundleElementByBundleAndVariant(bundleElementService *BundleElementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		bundleId := c.Query("bundle_id")
		variantId := c.Query("variant_id")

		products, err := bundleElementService.GetBundleElementByBundleAndVariant(bundleId, variantId, companyId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal error, please try later")
			return
		}
		c.IndentedJSON(http.StatusOK, products)
	}
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
