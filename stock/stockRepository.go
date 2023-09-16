package stock

import (
	"fr/greytsu/sol_api_products/models"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"context"
	"database/sql"
	"sync"
)

type StockRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (stockRepository *StockRepository) GetStockByVariantIdAndWarehouseId(variantId int, warehouseId int, companyId int) (*models.Stock, error) {
	stock, err := models.Stocks(qm.Where("company_id=?", companyId), qm.Where("fk_variant_id=?", variantId), qm.Where("fk_warehouse_id=?", warehouseId)).One(context.Background(), stockRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Int("Quantity", stock.Quantity).Msg("Stock found")
	return stock, nil
}

func (stockRepository *StockRepository) CreateStock(variantId int, warehouseId int, quantity int, companyId int) (*models.Stock, error) {
	stockRepository.Lock()
	defer stockRepository.Unlock()
	log.Debug().Msg("Creating stock")
	stock := models.Stock{CompanyID: companyId, FKWarehouseID: warehouseId, FKVariantID: variantId, Quantity: quantity}
	err := stock.Insert(context.Background(), stockRepository.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (stockRepository *StockRepository) UpdateStock(stock *models.Stock) error {
	stockRepository.Lock()
	defer stockRepository.Unlock()
	log.Debug().Int("Stock ID", stock.ID).Msg("Updating stock")
	_, err := stock.Update(context.Background(), stockRepository.db, boil.Infer())
	return err
}
