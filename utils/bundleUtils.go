package utils

import "fr/greytsu/sol_api_products/models"

func MergeBundles(baseBundle *models.Bundle, newBundle *models.Bundle) {
	baseBundle.Reference = newBundle.Reference
	baseBundle.Name = newBundle.Name
	baseBundle.Price = newBundle.Price
}
