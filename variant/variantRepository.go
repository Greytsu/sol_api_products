package variant

import (
	"fr/greytsu/sol_api_products/models"

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
	return variant, err
}
