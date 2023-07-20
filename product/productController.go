package product

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *gin.Engine, productService *ProductService) {
	router.GET("/products", getAllProducts(productService))
	router.POST("/products", postProduct(productService))
}

func getAllProducts(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := productService.GetAllProducts()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, "error")
		}
		c.IndentedJSON(http.StatusOK, products)
	}
}

func postProduct(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct models.PRProduct
		if err := c.BindJSON(&newProduct); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
		}
		product, err := productService.createProduct(&newProduct)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}
