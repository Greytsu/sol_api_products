package product

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/utils"
	"github.com/friendsofgo/errors"

	"strconv"
)

type ProductService struct {
	productRepository productRepository
}

type productRepository interface {
	GetAllProducts(companyId string) ([]*models.Product, error)
	GetProductsLike(name string, companyId string) ([]*models.Product, error)
	FindProduct(id string, companyId string) (*models.Product, error)
	GetProduct(id string, companyId string) (*dto.ProductDetails, error)
	GetProductByReference(reference string, companyId string) (*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) error
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

func (productService ProductService) GetProductByReference(reference string, companyId string) (*models.Product, error) {
	product, err := productService.productRepository.GetProductByReference(reference, companyId)
	return product, err
}

func (productService ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	productFound, _ := productService.GetProductByReference(product.Reference, strconv.Itoa(product.CompanyID))
	if productFound != nil {
		return nil, errors.New("Product already exists. ID: " + strconv.Itoa(productFound.ID))
	}
	return productService.productRepository.CreateProduct(product)
}

func (productService ProductService) UpdateProduct(id int, companyId int, newProduct *models.Product) (*models.Product, error) {
	baseProduct, _ := productService.productRepository.FindProduct(strconv.Itoa(id), strconv.Itoa(companyId))
	if baseProduct == nil {
		return nil, errors.New("Product not found.")
	}
	foundProduct, _ := productService.GetProductByReference(newProduct.Reference, strconv.Itoa(companyId))
	if foundProduct != nil && foundProduct.ID != id {
		return nil, errors.New("Reference already taken. ID: " + strconv.Itoa(foundProduct.ID))
	}
	baseProduct = utils.MergeProducts(baseProduct, newProduct)
	err := productService.productRepository.UpdateProduct(baseProduct)
	return baseProduct, err
}

func (productService ProductService) DeleteProduct(id int, companyId string) error {
	return productService.productRepository.DeleteProduct(id, companyId)
}
