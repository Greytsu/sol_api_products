package product

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
)

type ProductService struct {
	productRepository *ProductRepository
}

func NewProductService(productRepo *ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (productService ProductService) GetAllProducts(companyId string) ([]*models.Product, error) {
	return productService.productRepository.getAllProducts(companyId)
}

func (productService ProductService) GetProduct(id string, companyId string) (*dto.ProductDetails, error) {
	product, err := productService.productRepository.getProduct(id, companyId)
	return product, err
}

func (productService ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return productService.productRepository.createProduct(product)
}
