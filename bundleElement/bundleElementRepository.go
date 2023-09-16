package bundleElement

import (
	"context"
	"database/sql"
	"fr/greytsu/sol_api_products/models"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sync"
)

type BundleElementRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewBundleElementRepository(db *sql.DB) *BundleElementRepository {
	return &BundleElementRepository{db: db}
}

func (bundleElementRepository *BundleElementRepository) GetBundleElementByWarehouseAndVariant(bundleId string, variantId string, companyId string) (*models.BundleElement, error) {
	bundleElement, err := models.BundleElements(qm.Where("fk_bundle_id=?", bundleId), qm.Where("fk_variant_id=?", variantId), qm.Where("company_id=?", companyId)).One(context.Background(), bundleElementRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Msg("bundles found")
	return bundleElement, nil
}

func (bundleElementRepository *BundleElementRepository) CreateBundleElement(bundleElement *models.BundleElement) error {
	bundleElementRepository.Lock()
	defer bundleElementRepository.Unlock()
	log.Debug().Msg("Creating bundle")
	err := bundleElement.Insert(context.Background(), bundleElementRepository.db, boil.Infer())
	return err
}

func (bundleElementRepository *BundleElementRepository) DeleteBundleElement(id int, companyId string) error {
	bundleElementRepository.Lock()
	defer bundleElementRepository.Unlock()
	log.Debug().Int("Bundle element ID", id).Msg("Deleting bundle element")
	bundleElement, err := models.BundleElements(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), bundleElementRepository.db)
	if err != nil {
		return err
	}
	_, err = bundleElement.Delete(context.Background(), bundleElementRepository.db)
	return err
}
