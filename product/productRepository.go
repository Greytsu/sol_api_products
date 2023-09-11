package product

import (
	"context"
	"database/sql"
	"sync"

	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProductRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (productRepository *ProductRepository) getAllProducts(companyId string) ([]*models.Product, error) {
	products, err := models.Products(qm.Where("company_id=?", companyId)).All(context.Background(), productRepository.db)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *ProductRepository) getProductsLike(name string, companyId string) ([]*models.Product, error) {
	products, err := models.Products(qm.Where("company_id=?", companyId), qm.Where("reference like ? or name like ?", "%"+name+"%", "%"+name+"%")).All(context.Background(), productRepository.db)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *ProductRepository) getProduct(id string, companyId string) (*dto.ProductDetails, error) {
	product, err := models.Products(qm.Load(qm.Rels(models.ProductRels.FKProductVariants, models.VariantRels.FKVariantStocks)), qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), productRepository.db)
	if err != nil {
		return nil, err
	}

	return dto.NewProductDetails(product), nil
}

func (productRepository *ProductRepository) createProduct(product *models.Product) (*models.Product, error) {
	productRepository.Lock()
	defer productRepository.Unlock()
	err := product.Insert(context.Background(), productRepository.db, boil.Infer())
	return product, err
}
