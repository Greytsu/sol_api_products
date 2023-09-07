package dto

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/volatiletech/null/v8"
	"time"
)

type ProductDetails struct {
	ID         int            `json:"id" toml:"id" yaml:"id"`
	CompanyID  int            `json:"company_id" toml:"company_id" yaml:"company_id"`
	Name       string         `json:"name" toml:"name" yaml:"name"`
	Type       string         `json:"type" toml:"type" yaml:"type"`
	Variants   []VariantStock `json:"variants" toml:"variants" yaml:"variants"`
	CreateTime time.Time      `json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime time.Time      `json:"update_time" toml:"update_time" yaml:"update_time"`
	Deleted    null.Bool      `json:"deleted,omitempty" toml:"deleted" yaml:"deleted,omitempty"`
}

func NewProductDetails(product *models.Product) *ProductDetails {
	var productVariants ProductDetails
	productVariants.ID = product.ID
	productVariants.CompanyID = product.CompanyID
	productVariants.Name = product.Name
	productVariants.Type = product.Type
	productVariants.CreateTime = product.CreateTime
	productVariants.UpdateTime = product.UpdateTime
	productVariants.Deleted = product.Deleted

	var variantsStock []VariantStock
	for _, value := range product.R.GetFKProductVariants() {
		var variant VariantStock
		variant.ID = value.ID
		variant.CompanyID = value.CompanyID
		variant.Name = value.Name
		variant.StockTracking = value.StockTracking
		variant.PurchasePrice = value.PurchasePrice
		variant.SellingPrice = value.SellingPrice
		variant.CreateTime = value.CreateTime
		variant.UpdateTime = value.UpdateTime
		variant.Deleted = value.Deleted
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
