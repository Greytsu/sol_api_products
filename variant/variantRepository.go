package variant

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"sync"

	"fr/greytsu/sol_api_products/models"
)

type VariantRepository struct {
	db *sql.DB
	sync.Mutex
}

func NewVariantRepository(db *sql.DB) *VariantRepository {
	return &VariantRepository{db: db}
}

func (variantRepository *VariantRepository) createVariant(variant *models.Variant) (*models.Variant, error) {
	variantRepository.Lock()
	defer variantRepository.Unlock()
	err := variant.Insert(context.Background(), variantRepository.db, boil.Infer())
	return variant, err
}
