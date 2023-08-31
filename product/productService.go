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
	return productService.productRepository.GetAllProducts(companyId)
}

func (productService ProductService) GetProduct(id string, companyId string) (*dto.ProductDetails, error) {
	product, err := productService.productRepository.GetProduct(id, companyId)
	return product, err
}

func (productService ProductService) createProduct(product *models.Product) (*models.Product, error) {
	return productService.productRepository.CreateProduct(product)
}
