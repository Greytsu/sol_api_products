package warehouse

import "fr/greytsu/sol_api_products/models"

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
