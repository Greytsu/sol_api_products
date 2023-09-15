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
	routerGroup.PUT("/products/:id", putProduct(productService))
	routerGroup.POST("/products/:id/variants", postVariant(variantService))
	routerGroup.PUT("/variants/:variantId", putVariant(variantService))
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

func putProduct(productService *ProductService) gin.HandlerFunc {
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
		var updateProduct models.Product
		if err := c.BindJSON(&updateProduct); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Error while parsing JSON")
			return
		}
		product, err := productService.UpdateProduct(id, companyId, &updateProduct)
		if err != nil {
			if strings.Contains(err.Error(), "Product not found") || strings.Contains(err.Error(), "Reference already taken") {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			c.IndentedJSON(http.StatusInternalServerError, "Error while updating variant")
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

		productIdStr := c.Param("id")
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

func putVariant(variantService *variant.VariantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		companyIdStr := c.Request.Header["Company_id"][0]
		companyId, err := strconv.Atoi(companyIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid company id")
			return
		}
		variantIdStr := c.Param("variantId")
		variantId, err := strconv.Atoi(variantIdStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid product id")
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
