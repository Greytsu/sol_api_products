package bundle

import (
	"fr/greytsu/sol_api_products/bundleElement"
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/utils"
	"github.com/friendsofgo/errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

type BundleService struct {
	BundleRepository     *BundleRepository
	BundleElementService *bundleElement.BundleElementService
}

func NewBundleService(bundleRepo *BundleRepository, bundleElementService *bundleElement.BundleElementService) *BundleService {
	return &BundleService{
		BundleRepository:     bundleRepo,
		BundleElementService: bundleElementService,
	}
}

func (bundleService *BundleService) GetAllBundles(companyId string) ([]*models.Bundle, error) {
	return bundleService.BundleRepository.GetAllBundles(companyId)
}

func (bundleService *BundleService) GetBundle(id string, companyId string) (*dto.BundleDetails, error) {
	bundle, err := bundleService.BundleRepository.GetBundle(id, companyId)
	return bundle, err
}

func (bundleService *BundleService) CreateBundle(newBundle *models.Bundle, companyId string) (*models.Bundle, error) {
	foundBundle, err := bundleService.BundleRepository.GetBundleByReference(newBundle.Reference, companyId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Debug().Msg("Error while checking if name is available for bundle")
		return nil, errors.New("Error while creating bundle")
	}
	if foundBundle != nil {
		return nil, errors.New("Bundle name already taken")
	}
	err = bundleService.BundleRepository.CreateBundle(newBundle)
	if err != nil {
		return nil, errors.New("Error while creating bundle")
	}
	return newBundle, nil
}

func (bundleService *BundleService) UpdateBundle(id int, companyId int, newBundle *models.Bundle) (*models.Bundle, error) {
	baseBundle, _ := bundleService.BundleRepository.FindBundle(strconv.Itoa(id), strconv.Itoa(companyId))
	if baseBundle == nil {
		return nil, errors.New("Bundle not found.")
	}
	founBundleRef, _ := bundleService.BundleRepository.GetBundleByReference(newBundle.Reference, strconv.Itoa(companyId))
	if founBundleRef != nil && founBundleRef.ID != id {
		return nil, errors.New("Reference already taken. ID: " + strconv.Itoa(founBundleRef.ID))
	}
	utils.MergeBundles(baseBundle, newBundle)
	err := bundleService.BundleRepository.UpdateBundle(baseBundle)
	return baseBundle, err
}

func (bundleService *BundleService) DeleteBundle(id int, companyId string) error {
	return bundleService.BundleRepository.DeleteBundle(id, companyId)
}
