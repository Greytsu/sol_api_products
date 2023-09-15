package dto

import (
	"fr/greytsu/sol_api_products/models"
	"time"
)

type ProductDetails struct {
	ID         int            `json:"id" toml:"id" yaml:"id"`
	CompanyID  int            `json:"company_id" toml:"company_id" yaml:"company_id"`
	Reference  string         `json:"reference" toml:"reference" yaml:"reference"`
	Name       string         `json:"name" toml:"name" yaml:"name"`
	Variants   []VariantStock `json:"variants" toml:"variants" yaml:"variants"`
	CreateTime time.Time      `json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime time.Time      `json:"update_time" toml:"update_time" yaml:"update_time"`
}

func NewProductDetails(product *models.Product) *ProductDetails {
	var productVariants ProductDetails
	productVariants.ID = product.ID
	productVariants.Reference = product.Reference
	productVariants.CompanyID = product.CompanyID
	productVariants.Name = product.Name
	productVariants.CreateTime = product.CreateTime
	productVariants.UpdateTime = product.UpdateTime

	var variantsStock []VariantStock
	for _, value := range product.R.GetFKProductVariants() {
		var variant VariantStock
		variant.ID = value.ID
		variant.CompanyID = value.CompanyID
		variant.Reference = value.Reference
		variant.Name = value.Name
		variant.StockTracking = value.StockTracking
		variant.PurchasePrice = value.PurchasePrice
		variant.SellingPrice = value.SellingPrice
		variant.CreateTime = value.CreateTime
		variant.UpdateTime = value.UpdateTime
		variant.FKProductID = value.FKProductID
		stocks := value.R.GetFKVariantStocks()
		if stocks != nil {
			variant.Stock = *stocks[0]
		}
		variantsStock = append(variantsStock, variant)
	}
	productVariants.Variants = variantsStock
	return &productVariants
}
