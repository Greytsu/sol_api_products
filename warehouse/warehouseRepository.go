package warehouse

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
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
	if warehouses == nil {
		warehouses = []*models.Warehouse{}
	}
	log.Debug().Int("amount", len(warehouses)).Msg("Warehouses found")
	return warehouses, nil
}

func (warehouseRepository *WarehouseRepository) getWarehouse(id string, companyId string) (*models.Warehouse, error) {
	warehouse, err := models.Warehouses(qm.Where("company_id=?", companyId), qm.Where("id=?", id)).One(context.Background(), warehouseRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("Name", warehouse.Name).Msg("Warehouse found")
	return warehouse, nil
}

func (warehouseRepository *WarehouseRepository) getWarehouseByName(name string, companyId string) (*models.Warehouse, error) {
	warehouse, err := models.Warehouses(qm.Where("company_id=?", companyId), qm.Where("name=?", name)).One(context.Background(), warehouseRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Int("ID", warehouse.ID).Msg("Warehouse found")
	return warehouse, nil
}

func (warehouseRepository *WarehouseRepository) createWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	warehouseRepository.Lock()
	defer warehouseRepository.Unlock()
	log.Debug().Msg("Creating warehouse")
	err := warehouse.Insert(context.Background(), warehouseRepository.db, boil.Infer())
	return warehouse, err
}

func (warehouseRepository *WarehouseRepository) updateWarehouse(warehouse *models.Warehouse) error {
	warehouseRepository.Lock()
	defer warehouseRepository.Unlock()
	log.Debug().Msg("Updating warehouse")
	_, err := warehouse.Update(context.Background(), warehouseRepository.db, boil.Infer())
	return err
}

func (warehouseRepository *WarehouseRepository) DeleteWarehouse(id int, companyId string) error {
	warehouseRepository.Lock()
	defer warehouseRepository.Unlock()
	log.Debug().Int("Warehouse ID", id).Msg("Deleting warehouse")
	warehouse, err := models.Warehouses(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), warehouseRepository.db)
	if err != nil {
		return err
	}
	_, err = warehouse.Delete(context.Background(), warehouseRepository.db)
	return err
}
