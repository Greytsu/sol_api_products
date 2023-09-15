package variant

import (
	"fr/greytsu/sol_api_products/models"
	"github.com/rs/zerolog/log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"context"
	"database/sql"
	"sync"
)

type VariantRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewVariantRepository(db *sql.DB) *VariantRepository {
	return &VariantRepository{db: db}
}

func (variantRepository *VariantRepository) GetVariant(id string, companyId string) (*models.Variant, error) {
	variant, err := models.Variants(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), variantRepository.db)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("Name", variant.Name).Msg("Variant found")
	return variant, nil
}

func (variantRepository *VariantRepository) GetVariantByReference(reference string, companyId string) (*models.Variant, error) {
	variant, err := models.Variants(qm.Where("reference=?", reference), qm.Where("company_id=?", companyId)).One(context.Background(), variantRepository.db)
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func (variantRepository *VariantRepository) CreateVariant(variant *models.Variant) (*models.Variant, error) {
	variantRepository.Lock()
	defer variantRepository.Unlock()
	err := variant.Insert(context.Background(), variantRepository.db, boil.Infer())
	log.Debug().Str("Name", variant.Name).Msg("Variant found")
	return variant, err
}

func (variantRepository *VariantRepository) UpdateVariant(variant *models.Variant) error {
	variantRepository.Lock()
	defer variantRepository.Unlock()
	log.Debug().Int("Variant ID", variant.ID).Msg("Updating variant")
	rows, err := variant.Update(context.Background(), variantRepository.db, boil.Infer())
	log.Debug().Int64("rows", rows)
	return err
}

func (variantRepository *VariantRepository) DeleteVariant(id int, companyId string) error {
	variantRepository.Lock()
	defer variantRepository.Unlock()
	log.Debug().Int("Variant ID", id).Msg("Deleting variant")
	variant, err := models.Variants(qm.Where("id=?", id), qm.Where("company_id=?", companyId)).One(context.Background(), variantRepository.db)
	if err != nil {
		return err
	}
	_, err = variant.Delete(context.Background(), variantRepository.db)
	return err
}
