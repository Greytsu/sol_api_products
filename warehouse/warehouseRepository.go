package warehouse

import (
	"context"
	"database/sql"
	"sync"

	"fr/greytsu/sol_api_products/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type WarehouseRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

func (warehouseRepository *WarehouseRepository) getAllWarehouses(companyId string) ([]*models.Warehouse, error) {
	warehouses, err := models.Warehouses(qm.Where("company_id=?", companyId)).All(context.Background(), warehouseRepository.db)
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (warehouseRepository *WarehouseRepository) getWarehouseByName(name string, companyId string) (*models.Warehouse, error) {
	warehouse, err := models.Warehouses(qm.Where("company_id=?", companyId), qm.Where("name=?", name)).One(context.Background(), warehouseRepository.db)
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (warehouseRepository *WarehouseRepository) createWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	warehouseRepository.Lock()
	defer warehouseRepository.Unlock()
	err := warehouse.Insert(context.Background(), warehouseRepository.db, boil.Infer())
	return warehouse, err
}
