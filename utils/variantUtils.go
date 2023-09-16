package utils

import "fr/greytsu/sol_api_products/models"

func MergeVariants(baseVariant *models.Variant, newVariant *models.Variant) {
	baseVariant.Reference = newVariant.Reference
	baseVariant.Name = newVariant.Name
	baseVariant.StockTracking = newVariant.StockTracking
	baseVariant.PurchasePrice = newVariant.PurchasePrice
	baseVariant.SellingPrice = newVariant.SellingPrice
}
