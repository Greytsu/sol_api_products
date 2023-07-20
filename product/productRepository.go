package product

import (
	"context"
	"database/sql"
	"fr/greytsu/sol_api_products/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"sync"
)

type ProductRepository struct {
	db *sql.DB
	sync.Mutex
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

func (productRepository *ProductRepository) CreateProduct(product *models.PRProduct) (*models.PRProduct, error) {
	productRepository.Lock()
	defer productRepository.Unlock()
	err := product.Insert(context.Background(), productRepository.db, boil.Infer())
	return product, err
}
