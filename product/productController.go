package product

import (
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/variant"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func RegisterProductRoutes(router *gin.Engine, productService *ProductService, variantService *variant.VariantService) {
	router.GET("/products", getAllProducts(productService))
	router.GET("/products/:id", getProduct(productService))
	router.POST("/products", postProduct(productService))
	router.POST("/products/:productId/variants", postVariant(variantService))
}

func getAllProducts(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Query("company_id")

		products, err := productService.GetAllProducts(companyId)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, "error")
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
			log.Printf("Error: %s", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, "error")
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
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}
