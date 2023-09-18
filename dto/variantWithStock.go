package dto

import (
	"github.com/volatiletech/null/v8"

	"time"
)

type VariantStock struct {
	ID            int          `json:"id" toml:"id" yaml:"id"`
	CompanyID     int          `json:"company_id" toml:"company_id" yaml:"company_id"`
	Reference     string       `json:"reference" toml:"reference" yaml:"reference"`
	Name          string       `json:"name" toml:"name" yaml:"name"`
	StockTracking bool         `json:"stock_tracking" toml:"stock_tracking" yaml:"stock_tracking"`
	PurchasePrice null.Float64 `json:"purchase_price,omitempty" toml:"purchase_price" yaml:"purchase_price,omitempty"`
	SellingPrice  null.Float64 `json:"selling_price,omitempty" toml:"selling_price" yaml:"selling_price,omitempty"`
	CreateTime    time.Time    `json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime    time.Time    `json:"update_time" toml:"update_time" yaml:"update_time"`
	FKProductID   int          `json:"product_id" toml:"fk_product_id" yaml:"fk_product_id"`
	Stock         int          `json:"stock" toml:"stocks" yaml:"stocks"`
}
