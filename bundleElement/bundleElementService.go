package bundleElement

import (
	"errors"
	"fr/greytsu/sol_api_products/models"
	"strconv"
)

type BundleElementService struct {
	BundleElementRepository *BundleElementRepository
}

func NewBundleElementService(bundleElementRepository *BundleElementRepository) *BundleElementService {
	return &BundleElementService{
		BundleElementRepository: bundleElementRepository,
	}
}

func (bundleElementService *BundleElementService) GetBundleElementByBundleAndVariant(bundleId string, variantId string, companyId string) (*models.BundleElement, error) {
	foundBundleElement, err := bundleElementService.BundleElementRepository.GetBundleElementByBundleAndVariant(bundleId, variantId, companyId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	return foundBundleElement, nil
}

func (bundleElementService *BundleElementService) CreateBundleElement(bundleElement *models.BundleElement) (*models.BundleElement, error) {
	foundBundleElement, err := bundleElementService.BundleElementRepository.GetBundleElementByBundleAndVariant(strconv.Itoa(bundleElement.FKBundleID), strconv.Itoa(bundleElement.FKVariantID), strconv.Itoa(bundleElement.CompanyID))
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if foundBundleElement != nil {
		return nil, errors.New("Product already exists. ID: " + strconv.Itoa(foundBundleElement.ID))
	}
	err = bundleElementService.BundleElementRepository.CreateBundleElement(bundleElement)
	return bundleElement, nil
}

func (bundleElementService *BundleElementService) DeleteBundleElement(id int, companyId string) error {
	return bundleElementService.BundleElementRepository.DeleteBundleElement(id, companyId)
}
