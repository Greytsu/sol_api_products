package stock

import (
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
