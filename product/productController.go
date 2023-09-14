package product

import (
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/variant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterProductRoutes(router *gin.Engine, productService *ProductService, variantService *variant.VariantService) {
	router.GET("/products", getAllProducts(productService))
	router.GET("/products/:id", getProduct(productService))
	router.POST("/products", postProduct(productService))
	router.POST("/products/:productId/variants", postVariant(variantService))
	router.DELETE("/products/:id", deleteProduct(productService))
}

func getAllProducts(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Query("company_id")
		name := c.Query("name")

		products, err := productService.GetAllProducts(name, companyId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal error, please try later")
			return
		}
		c.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Query("company_id")
		id := c.Param("id")

		product, err := productService.GetProduct(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Product not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Internal error, please try later")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}

func postProduct(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct models.Product
		if err := c.BindJSON(&newProduct); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
		}
		product, err := productService.CreateProduct(&newProduct)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}

func postVariant(variantService *variant.VariantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Param("productId")
		productId, err := strconv.Atoi(productIdStr)

		var newVariant models.Variant
		newVariant.FKProductID = productId
		if err := c.BindJSON(&newVariant); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
		}
		product, err := variantService.CreateVariant(&newVariant)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating variant")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}

func deleteProduct(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Query("company_id")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid id")
		}

		err = productService.DeleteProduct(id, companyId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.IndentedJSON(http.StatusNotFound, "Product not found")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while deleting product")
			return
		}
		c.IndentedJSON(http.StatusNoContent, "")
	}
}
