package product

import (
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/variant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func RegisterProductRoutes(routerGroup *gin.RouterGroup, productService *ProductService, variantService *variant.VariantService) {
	routerGroup.GET("/products", getAllProducts(productService))
	routerGroup.GET("/products/:id", getProduct(productService))
	routerGroup.POST("/products", postProduct(productService))
	routerGroup.POST("/products/:productId/variants", postVariant(variantService))
	routerGroup.DELETE("/products/:id", deleteProduct(productService))
}

func getAllProducts(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
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
		companyId := c.Request.Header["Company_id"][0]
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
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}

		var newProduct models.Product
		if err := c.BindJSON(&newProduct); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}

		newProduct.CompanyID = companyId
		product, err := productService.CreateProduct(&newProduct)
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

func postVariant(variantService *variant.VariantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}

		productIdStr := c.Param("productId")
		productId, err := strconv.Atoi(productIdStr)

		var newVariant models.Variant
		newVariant.FKProductID = productId
		newVariant.CompanyID = companyId
		if err := c.BindJSON(&newVariant); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		product, err := variantService.CreateVariant(&newVariant)
		if err != nil {
			if strings.Contains(err.Error(), "Variant already exists") {
				c.IndentedJSON(http.StatusBadRequest, "Reference already exists")
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while creating variant")
			return
		}
		c.IndentedJSON(http.StatusOK, product)
	}
}

func deleteProduct(productService *ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyId := c.Request.Header["Company_id"][0]
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
