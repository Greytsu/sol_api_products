package dto

type StockOperation struct {
	WarehouseId int    `json:"warehouse_id" toml:"warehouse_id" yaml:"warehouse_id"`
	Quantity    int    `json:"quantity" toml:"quantity" yaml:"quantity"`
	Type        string `json:"type" toml:"type" yaml:"type"`
	VariantId   int
}
