package product

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *gin.Engine, productService *ProductService) {
	router.GET("/products", getAllProducts(productService))
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
