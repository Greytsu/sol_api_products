package product

import (
	"context"
	"database/sql"
	"fr/greytsu/sol_api_products/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (productRepository *ProductRepository) GetAllProducts() ([]*models.PRProduct, error) {

	products, err := models.PRProducts().All(context.Background(), productRepository.db)
	if err != nil {
		return nil, err
	}
	return products, nil
}
