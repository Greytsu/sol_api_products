package product

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *gin.Engine, productService *ProductService) {
	router.GET("/products", getAllProducts(productService))
	router.GET("/products/:id", getProduct(productService))
	router.POST("/products", postProduct(productService))
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
		product, err := productService.createProduct(&newProduct)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating product")
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}
