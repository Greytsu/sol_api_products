package warehouse

import (
	"errors"
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/utils"
	"github.com/rs/zerolog/log"
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

func (warehouseService *WarehouseService) GetAllWarehouses(companyId string) ([]*models.Warehouse, error) {
	return warehouseService.warehouseRepository.getAllWarehouses(companyId)
}

func (warehouseService *WarehouseService) GetWarehouse(id string, companyId string) (*models.Warehouse, error) {
	return warehouseService.warehouseRepository.getWarehouse(id, companyId)
}
func (warehouseService *WarehouseService) GetWarehouseByName(name string, companyId string) (*models.Warehouse, error) {
	return warehouseService.warehouseRepository.getWarehouseByName(name, companyId)
}

func (warehouseService *WarehouseService) CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	warehouseFound, _ := warehouseService.GetWarehouseByName(warehouse.Name, strconv.Itoa(warehouse.CompanyID))
	if warehouseFound != nil {
		return nil, errors.New("Warehouse already exists. ID: " + strconv.Itoa(warehouseFound.ID))
	}
	return warehouseService.warehouseRepository.createWarehouse(warehouse)
}

func (warehouseService *WarehouseService) UpdateWarehouse(id int, companyId int, newWarehouse *models.Warehouse) (*models.Warehouse, error) {
	warehouseFoundName, _ := warehouseService.GetWarehouseByName(newWarehouse.Name, strconv.Itoa(newWarehouse.CompanyID))
	if warehouseFoundName != nil && warehouseFoundName.ID == id {
		log.Debug().Msg("Warehouse name already exists. ID: " + strconv.Itoa(warehouseFoundName.ID))
		return nil, errors.New("Warehouse already exists")
	}
	warehouseFoundId, _ := warehouseService.GetWarehouse(strconv.Itoa(id), strconv.Itoa(companyId))
	if warehouseFoundId == nil {
		return nil, errors.New("Warehouse not found")
	}

	utils.MergeWarehouses(warehouseFoundId, newWarehouse)
	err := warehouseService.warehouseRepository.updateWarehouse(warehouseFoundId)

	return warehouseFoundId, err
}
