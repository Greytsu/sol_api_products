package mock

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
	"strconv"
)

type ProductRepositoryMock struct {
	products []*models.Product
}

func NewProductRepositoryMock(productsMock []*models.Product) *ProductRepositoryMock {

	return &ProductRepositoryMock{
		products: productsMock,
	}
}

func (productRepositoryMock *ProductRepositoryMock) GetAllProducts(companyId string) ([]*models.Product, error) {

	compId, err := strconv.Atoi(companyId)
	if err != nil {
		return nil, err
	}

	condition := func(product *models.Product) bool {
		return product.CompanyID == compId
	}

	filteredProducts := filter(productRepositoryMock.products, condition)

	return filteredProducts, nil
}

func (productRepositoryMock *ProductRepositoryMock) GetProductsLike(name string, companyId string) ([]*models.Product, error) {
	return nil, nil
}
func (productRepositoryMock *ProductRepositoryMock) GetProduct(id string, companyId string) (*dto.ProductDetails, error) {
	return nil, nil
}
func (productRepositoryMock *ProductRepositoryMock) CreateProduct(product *models.Product) (*models.Product, error) {
	return nil, nil
}
func (productRepositoryMock *ProductRepositoryMock) DeleteProduct(id int, companyId string) error {
	return nil
}

func filter(arr []*models.Product, condition func(*models.Product) bool) []*models.Product {
	var filtered []*models.Product
	for _, item := range arr {
		if condition(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
