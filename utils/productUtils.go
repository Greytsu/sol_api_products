package utils

import "fr/greytsu/sol_api_products/models"

func MergeProducts(baseProduct *models.Product, newProduct *models.Product) *models.Product {
	if newProduct.Name != "" {
		baseProduct.Name = newProduct.Name
	}

	if newProduct.Reference != "" {
		baseProduct.Reference = newProduct.Reference
	}

	return baseProduct
}
