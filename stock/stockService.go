package stock

type StockService struct {
	stockRepository *StockRepository
}

func NewStockService(stockRepository *StockRepository) *StockService {
	return &StockService{
		stockRepository: stockRepository,
	}
}
