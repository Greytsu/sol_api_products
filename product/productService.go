package product

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
)

type ProductService struct {
	productRepository productRepository
}

type productRepository interface {
	GetAllProducts(companyId string) ([]*models.Product, error)
	GetProductsLike(name string, companyId string) ([]*models.Product, error)
	GetProduct(id string, companyId string) (*dto.ProductDetails, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id int, companyId string) error
}

func NewProductService(productRepo productRepository) *ProductService {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (productService ProductService) GetAllProducts(name string, companyId string) ([]*models.Product, error) {
	if name != "" {
		return productService.productRepository.GetProductsLike(name, companyId)
	}
	return productService.productRepository.GetAllProducts(companyId)
}

func (productService ProductService) GetProduct(id string, companyId string) (*dto.ProductDetails, error) {
	product, err := productService.productRepository.GetProduct(id, companyId)
	return product, err
}

func (productService ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return productService.productRepository.CreateProduct(product)
}

func (productService ProductService) DeleteProduct(id int, companyId string) error {
	return productService.productRepository.DeleteProduct(id, companyId)
}
