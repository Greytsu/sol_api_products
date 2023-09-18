package dto

import (
	"fr/greytsu/sol_api_products/models"
	"time"
)

type BundleDetails struct {
	ID         int            `json:"id" toml:"id" yaml:"id"`
	CompanyID  int            `json:"company_id" toml:"company_id" yaml:"company_id"`
	Reference  string         `json:"reference" toml:"reference" yaml:"reference"`
	Name       string         `json:"name" toml:"name" yaml:"name"`
	Price      float64        `boil:"price" json:"price" toml:"price" yaml:"price"`
	Variants   []VariantStock `json:"variants" toml:"variants" yaml:"variants"`
	CreateTime time.Time      `json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime time.Time      `json:"update_time" toml:"update_time" yaml:"update_time"`
}

func NewBundleDetails(bundle *models.Bundle) *BundleDetails {
	var bundleDetails BundleDetails
	bundleDetails.ID = bundle.ID
	bundleDetails.Reference = bundle.Reference
	bundleDetails.CompanyID = bundle.CompanyID
	bundleDetails.Name = bundle.Name
	bundleDetails.Price = bundle.Price
	bundleDetails.CreateTime = bundle.CreateTime
	bundleDetails.UpdateTime = bundle.UpdateTime

	var variants []VariantStock
	for _, bundleElement := range bundle.R.GetFKBundleBundleElements() {

		var variantStock VariantStock

		variant := bundleElement.R.GetFKVariant()
		variantStock.ID = variant.ID
		variantStock.CompanyID = variant.CompanyID
		variantStock.Reference = variant.Reference
		variantStock.Name = variant.Name
		variantStock.StockTracking = variant.StockTracking
		variantStock.PurchasePrice = variant.PurchasePrice
		variantStock.SellingPrice = variant.SellingPrice
		variantStock.CreateTime = variant.CreateTime
		variantStock.UpdateTime = variant.UpdateTime
		variantStock.FKProductID = variant.FKProductID

		stocks := variant.R.GetFKVariantStocks()
		if stocks != nil {
			var stockAvailable int
			for _, stock := range stocks {
				stockAvailable = stockAvailable + stock.Quantity
			}
			variantStock.Stock = stockAvailable
		}

		variants = append(variants, variantStock)
	}
	bundleDetails.Variants = variants
	return &bundleDetails
}
