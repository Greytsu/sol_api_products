package warehouse

import (
	"errors"
	"fr/greytsu/sol_api_products/models"
	"strconv"
)

type WarehouseService struct {
	warehouseRepository *WarehouseRepository
}

func NewWarehouseService(warehouseRepo *WarehouseRepository) *WarehouseService {
	return &WarehouseService{
		warehouseRepository: warehouseRepo,
	}
}

func (warehouseService WarehouseService) GetAllWarehouses(companyId string) ([]*models.Warehouse, error) {
	return warehouseService.warehouseRepository.getAllWarehouses(companyId)
}

func (warehouseService WarehouseService) getWarehouseByName(name string, companyId string) (*models.Warehouse, error) {
	return warehouseService.warehouseRepository.getWarehouseByName(name, companyId)
}

func (warehouseService WarehouseService) CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	warehouseFound, _ := warehouseService.getWarehouseByName(warehouse.Name, strconv.Itoa(warehouse.CompanyID))
	if warehouseFound != nil {
		return nil, errors.New("Warehouse already exists. ID: " + strconv.Itoa(warehouseFound.ID))
	}
	return warehouseService.warehouseRepository.createWarehouse(warehouse)
}
