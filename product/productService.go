package product

import "fr/greytsu/sol_api_products/models"

type ProductService struct {
	productRepository *ProductRepository
}

func NewProductService(productRepo *ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepo,
	}
}

func (productService ProductService) GetAllProducts() ([]*models.PRProduct, error) {
	return productService.productRepository.GetAllProducts()
}

func (productService ProductService) createProduct(product *models.PRProduct) (*models.PRProduct, error) {
	return productService.productRepository.CreateProduct(product)
}
