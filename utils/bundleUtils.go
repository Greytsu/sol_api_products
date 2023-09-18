package utils

import "fr/greytsu/sol_api_products/models"

func MergeBundles(baseBundle *models.Bundle, newBundle *models.Bundle) {
	baseBundle.Reference = newBundle.Reference
	baseBundle.Name = newBundle.Name
	baseBundle.Price = newBundle.Price
}

func GetBundlesFromProduct(variant *models.Variant) []*models.Bundle {
	var bundles []*models.Bundle

	for _, bundleElement := range variant.R.GetFKVariantBundleElements() {
		bundles = append(bundles, bundleElement.R.GetFKBundle())
	}
	return removeDuplicates(bundles)
}

func removeDuplicates(bundles []*models.Bundle) []*models.Bundle {
	// Create a map to store unique IDs
	uniqueIDs := make(map[int]struct{})
	result := []*models.Bundle{}

	for _, bundle := range bundles {
		// Check if the item's ID is already in the map
		if _, found := uniqueIDs[bundle.ID]; !found {
			// If not found, add it to the map and the result slice
			uniqueIDs[bundle.ID] = struct{}{}
			result = append(result, bundle)
		}
	}

	return result
}
