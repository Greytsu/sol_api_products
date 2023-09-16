package bundle

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"

	"net/http"
)

func RegisterBundlesRoutes(routerGroup *gin.RouterGroup, bundleService *BundleService) {
	routerGroup.GET("/bundles", getAllBundles(bundleService))
	routerGroup.GET("/bundles/:id", getBundle(bundleService))
	routerGroup.POST("/bundles", postBundle(bundleService))
	routerGroup.PUT("/bundles/:id", putBundle(bundleService))
	routerGroup.DELETE("/bundles/:id", deleteBundle(bundleService))

	routerGroup.POST("/bundles/:id/elements", postBundleElement(bundleService))
}

func getAllBundles(bundleService *BundleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		bundles, err := bundleService.GetAllBundles(companyId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal error, please try later")
			return
		}
		c.IndentedJSON(http.StatusOK, bundles)
	}
}

func getBundle(bundleService *BundleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		_, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}
		idStr := c.Param("id")
		_, err = strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid product id")
			return
		}

		bundles, err := bundleService.GetBundle(idStr, companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal error, please try later")
			return
		}
		c.IndentedJSON(http.StatusOK, bundles)
	}
}

func postBundle(bundleService *BundleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}

		var newBundle models.Bundle
		if err := c.BindJSON(&newBundle); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}

		newBundle.CompanyID = companyId
		product, err := bundleService.CreateBundle(&newBundle, companyIdStr)
		if err != nil {
			if strings.Contains(err.Error(), "Product already exists") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
			return
		}

		c.IndentedJSON(http.StatusOK, product)
	}
}

func putBundle(bundleService *BundleService) gin.HandlerFunc {
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
			c.IndentedJSON(http.StatusBadRequest, "Invalid product id")
			return
		}
		var updateBundle models.Bundle
		if err := c.BindJSON(&updateBundle); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		updatedBundle, err := bundleService.UpdateBundle(id, companyId, &updateBundle)
		if err != nil {
			if strings.Contains(err.Error(), "Bundle not found") || strings.Contains(err.Error(), "Reference already taken") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while updating variant")
			return
		}
		c.IndentedJSON(http.StatusOK, updatedBundle)

	}
}

func deleteBundle(bundleService *BundleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid id")
			return
		}

		err = bundleService.DeleteBundle(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Bundle not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while deleting bundle")
			return
		}
		c.IndentedJSON(http.StatusNoContent, "")
	}
}

func postBundleElement(bundleService *BundleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}

		var newBundleElement models.BundleElement
		if err := c.BindJSON(&newBundleElement); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}

		newBundleElement.CompanyID = companyId
		bundleElement, err := bundleService.BundleElementService.CreateBundleElement(&newBundleElement)
		if err != nil {
			if strings.Contains(err.Error(), "Product already exists") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
			return
		}

		c.IndentedJSON(http.StatusOK, bundleElement)
	}
}
