package product

import (
	"fr/greytsu/sol_api_products/mock"
	"fr/greytsu/sol_api_products/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var products = []*models.Product{
	{ID: 1, CompanyID: 2508, Reference: "TSHIRT", Name: "T-shirt", CreateTime: time.Time{}, UpdateTime: time.Time{}},
	{ID: 2, CompanyID: 2508, Reference: "BAGUE", Name: "Bague", CreateTime: time.Time{}, UpdateTime: time.Time{}},
	{ID: 3, CompanyID: 8, Reference: "TABLE", Name: "Table", CreateTime: time.Time{}, UpdateTime: time.Time{}},
}

func TestGetAllProducts(t *testing.T) {
	mockRepo := mock.NewProductRepositoryMock(products)
	productService := NewProductService(mockRepo)

	products, err := productService.GetAllProducts("", "2508")

	assert.Nil(t, err)
	assert.True(t, len(products) == 2)
}
