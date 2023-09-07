package variant

import "fr/greytsu/sol_api_products/models"

type VariantService struct {
	VariantRepository *VariantRepository
}

func NewVariantService(VariantRepo *VariantRepository) *VariantService {
	return &VariantService{
		VariantRepository: VariantRepo,
	}
}

func (variantService VariantService) CreateVariant(variant *models.Variant) (*models.Variant, error) {
	return variantService.VariantRepository.createVariant(variant)
}
