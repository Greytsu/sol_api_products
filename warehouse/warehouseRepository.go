package warehouse

import (
	"context"
	"database/sql"
	"fr/greytsu/sol_api_products/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sync"
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
