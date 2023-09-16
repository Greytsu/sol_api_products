package bundle

import (
	"context"
	"database/sql"
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sync"
)

type BundleRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewBundleRepository(db *sql.DB) *BundleRepository {
	return &BundleRepository{db: db}
}

func (bundleRepository *BundleRepository) GetAllBundles(companyId string) ([]*models.Bundle, error) {

	bundles, err := models.Bundles(qm.Where("company_id=?", companyId)).All(context.Background(), bundleRepository.db)
	if err != nil {
		return nil, err
	}
	if bundles == nil {
		bundles = []*models.Bundle{}
	}
	log.Debug().Int("amount", len(bundles)).Msg("bundles found")
	return bundles, nil
}

func (bundleRepository *BundleRepository) FindBundle(id string, companyId string) (*models.Bundle, error) {
	bundle, err := models.Bundles(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), bundleRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("Bundle name", bundle.Name).Msg("bundles found")
	return bundle, nil
}

func (bundleRepository *BundleRepository) GetBundle(id string, companyId string) (*dto.BundleDetails, error) {
	bundle, err := models.Bundles(qm.Load(qm.Rels(models.BundleRels.FKBundleBundleElements, models.BundleElementRels.FKVariant, models.VariantRels.FKVariantStocks)), qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), bundleRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("Bundle name", bundle.Name).Msg("bundles found")
	bundleDetails := dto.NewBundleDetails(bundle)
	return bundleDetails, nil
}

func (bundleRepository *BundleRepository) GetBundleByReference(reference string, companyId string) (*models.Bundle, error) {

	bundle, err := models.Bundles(qm.Where("company_id=?", companyId), qm.Where("reference=?", reference)).One(context.Background(), bundleRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Int("Bundle id", bundle.ID).Msg("bundles found")
	return bundle, nil
}

func (bundleRepository *BundleRepository) CreateBundle(bundle *models.Bundle) error {
	bundleRepository.Lock()
	defer bundleRepository.Unlock()
	log.Debug().Msg("Creating bundle")
	err := bundle.Insert(context.Background(), bundleRepository.db, boil.Infer())
	return err
}

func (bundleRepository *BundleRepository) UpdateBundle(bundle *models.Bundle) error {
	bundleRepository.Lock()
	defer bundleRepository.Unlock()
	log.Debug().Msg("Updating bundle")
	_, err := bundle.Update(context.Background(), bundleRepository.db, boil.Infer())
	return err
}

func (bundleRepository *BundleRepository) DeleteBundle(id int, companyId string) error {
	bundleRepository.Lock()
	defer bundleRepository.Unlock()
	log.Debug().Int("Bundle ID", id).Msg("Deleting bundle")
	bundle, err := models.Bundles(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), bundleRepository.db)
	if err != nil {
		return err
	}
	_, err = bundle.Delete(context.Background(), bundleRepository.db)
	return err
}
