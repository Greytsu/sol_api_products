package variant

import (
	"errors"
	"fr/greytsu/sol_api_products/models"
	"strconv"
)

type VariantService struct {
	VariantRepository *VariantRepository
}

func NewVariantService(VariantRepo *VariantRepository) *VariantService {
	return &VariantService{
		VariantRepository: VariantRepo,
	}
}

func (variantService VariantService) GetVariantByReference(reference string, companyId string) (*models.Variant, error) {
	return variantService.VariantRepository.GetVariantByReference(reference, companyId)
}

func (variantService VariantService) CreateVariant(variant *models.Variant) (*models.Variant, error) {
	variantFound, _ := variantService.GetVariantByReference(variant.Reference, strconv.Itoa(variant.CompanyID))
	if variantFound != nil {
		return nil, errors.New("Variant already exists. ID: " + strconv.Itoa(variantFound.ID))
	}
	return variantService.VariantRepository.CreateVariant(variant)
}
