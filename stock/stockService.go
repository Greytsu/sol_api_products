package stock

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"

	"errors"
)

type StockService struct {
	stockRepository *StockRepository
}

func NewStockService(stockRepository *StockRepository) *StockService {
	return &StockService{
		stockRepository: stockRepository,
	}
}

func (stockService *StockService) StockOperation(operation dto.StockOperation, companyId int) (*models.Stock, error) {
	foundStock, err := stockService.stockRepository.GetStockByVariantIdAndWarehouseId(operation.VariantId, operation.WarehouseId, companyId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, errors.New("Error, try later")
	}

	if foundStock == nil && operation.Type == "ADD" {
		//No stock yet for this variant
		stock, err := stockService.stockRepository.CreateStock(operation.VariantId, operation.WarehouseId, operation.Quantity, companyId)
		if err != nil {
			return nil, errors.New("Error while creating stock")
		}
		return stock, nil
	}

	//Operation type, increments or decrements
	if operation.Type == "ADD" {
		foundStock.Quantity = foundStock.Quantity + operation.Quantity
	}
	if operation.Type == "REMOVE" {
		if operation.Quantity > foundStock.Quantity {
			return nil, errors.New("Not enough stock")
		}
		foundStock.Quantity = foundStock.Quantity - operation.Quantity
	}

	//Update stock
	err = stockService.stockRepository.UpdateStock(foundStock)
	if err != nil {
		return nil, errors.New("Error while updating stock")
	}

	return foundStock, nil
}
