package utils

import "fr/greytsu/sol_api_products/models"

func MergeWarehouses(baseWarehouse *models.Warehouse, newWarehouse *models.Warehouse) {
	baseWarehouse.Name = newWarehouse.Name
	baseWarehouse.Address = newWarehouse.Address
	baseWarehouse.Complement = newWarehouse.Complement
	baseWarehouse.Zip = newWarehouse.Zip
	baseWarehouse.City = newWarehouse.City
	baseWarehouse.Region = newWarehouse.Region
	baseWarehouse.Country = newWarehouse.Country
	baseWarehouse.Manager = newWarehouse.Manager
	baseWarehouse.Phone = newWarehouse.Phone
	baseWarehouse.Email = newWarehouse.Email
}
