package variant

import (
	"fr/greytsu/sol_api_products/models"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
	"strings"
)

func RegisterVariantsRoutes(routerGroup *gin.RouterGroup, variantService *VariantService) {
	routerGroup.PUT("/variants/:id", putVariant(variantService))
	routerGroup.DELETE("/variants/:id", deleteVariant(variantService))
}

func putVariant(variantService *VariantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}
		variantIdStr := c.Param("id")
		variantId, err := strconv.Atoi(variantIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid variant id")
			return
		}
		var updateVariant models.Variant
		if err := c.BindJSON(&updateVariant); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		updatedVariant, err := variantService.UpdateVariant(variantId, companyId, &updateVariant)
		if err != nil {
			if strings.Contains(err.Error(), "Variant not found") || strings.Contains(err.Error(), "Reference already taken") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while updating variant")
			return
		}
		c.IndentedJSON(http.StatusOK, updatedVariant)
	}
}

func deleteVariant(variantService *VariantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid id")
			return
		}

		err = variantService.DeleteVariant(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Variant not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while deleting variant")
			return
		}
		c.IndentedJSON(http.StatusNoContent, "")
	}
}
