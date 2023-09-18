package variant

import (
	"fr/greytsu/sol_api_products/dto"
	"fr/greytsu/sol_api_products/models"
	"fr/greytsu/sol_api_products/stock"
	"fr/greytsu/sol_api_products/utils"

	"errors"
	"strconv"
)

type VariantService struct {
	VariantRepository *VariantRepository
	StockService      *stock.StockService
}

func NewVariantService(variantRepo *VariantRepository, stockService *stock.StockService) *VariantService {
	return &VariantService{
		VariantRepository: variantRepo,
		StockService:      stockService,
	}
}

func (variantService *VariantService) GetVariant(id string, companyId string) (*models.Variant, error) {
	return variantService.VariantRepository.GetVariant(id, companyId)
}

func (variantService *VariantService) GetVariantBundles(id string, companyId string) ([]*models.Bundle, error) {
	return variantService.VariantRepository.GetVariantBundles(id, companyId)
}

func (variantService *VariantService) GetVariantByReference(reference string, companyId string) (*models.Variant, error) {
	return variantService.VariantRepository.GetVariantByReference(reference, companyId)
}

func (variantService *VariantService) CreateVariant(variant *models.Variant) (*models.Variant, error) {
	variantFound, _ := variantService.GetVariantByReference(variant.Reference, strconv.Itoa(variant.CompanyID))
	if variantFound != nil {
		return nil, errors.New("Variant already exists. ID: " + strconv.Itoa(variantFound.ID))
	}
	return variantService.VariantRepository.CreateVariant(variant)
}

func (variantService *VariantService) UpdateVariant(id int, companyId int, newVariant *models.Variant) (*models.Variant, error) {
	baseVariant, _ := variantService.GetVariant(strconv.Itoa(id), strconv.Itoa(companyId))
	if baseVariant == nil {
		return nil, errors.New("Variant not found.")
	}
	foundVariant, _ := variantService.GetVariantByReference(newVariant.Reference, strconv.Itoa(companyId))
	if foundVariant != nil && foundVariant.ID != id {
		return nil, errors.New("Reference already taken. ID: " + strconv.Itoa(foundVariant.ID))
	}
	utils.MergeVariants(baseVariant, newVariant)
	err := variantService.VariantRepository.UpdateVariant(baseVariant)
	return baseVariant, err
}

func (variantService *VariantService) DeleteVariant(id int, companyId string) error {
	return variantService.VariantRepository.DeleteVariant(id, companyId)
}

func (variantService *VariantService) StockOperation(operation dto.StockOperation, companyId int) (*models.Stock, error) {
	variant, err := variantService.GetVariant(strconv.Itoa(operation.VariantId), strconv.Itoa(companyId))
	if err != nil || variant == nil {
		return nil, errors.New("Variant not found")
	}
	newStock, err := variantService.StockService.StockOperation(operation, companyId)
	if err != nil {
		return nil, err
	}
	return newStock, nil
}
