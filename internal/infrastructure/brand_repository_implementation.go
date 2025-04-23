package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"gorm.io/gorm"
)

type BrandRepository struct {
	gorm *gorm.DB
}

func NewBrandRepository(gorm *gorm.DB) *BrandRepository {
	return &BrandRepository{
		gorm: gorm,
	}
}

func (c *BrandRepository) CreateBrand(brand entities.Brand) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&models.Brands{
		ID:            brand.ID,
		Active:        brand.Active,
		CreatedAt:     brand.CreatedAt,
		UpdatedAt:     brand.UpdatedAt,
		DeactivatedAt: brand.DeactivatedAt,
		Name:          brand.Name,
		Logo:          brand.Logo,
		VotesCount:    brand.VotesCount,
	}).Error; err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateBrand",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *BrandRepository) GetBrandByID(brandID string) (entities.Brand, error) {
	var brandModel models.Brands

	result := c.gorm.Model(&models.Brands{}).Where("id =? AND active =?", brandID, true).First(&brandModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Brand{}, errors.New("brand not found")
		}
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetBrandByID",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return entities.Brand{}, result.Error
	}

	return *brandModel.ToEntity(), nil
}

func (c *BrandRepository) ThisBrandExist(brandName string) (bool, error) {
	var brandModel models.Brands

	result := c.gorm.Model(&models.Brands{}).Where("name =?", brandName).First(&brandModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisBrandExist",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return false, result.Error
	}

	return true, nil
}

func (c *BrandRepository) GetBrandsByIDs(brandsIDs []string) ([]entities.Brand, error) {
	var brandsModel []models.Brands

	result := c.gorm.Model(&models.Brands{}).Where("id IN?", brandsIDs).Find(&brandsModel)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetBrandsByIDs",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var brands []entities.Brand
	for _, brandModel := range brandsModel {
		brands = append(brands, *brandModel.ToEntity())
	}

	return brands, nil
}

func (c *BrandRepository) UpdadeBrand(brand entities.Brand) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Model(&models.Brands{}).Where("id =?", brand.ID).Updates(models.Brands{
		Active:        brand.Active,
		Name:          brand.Name,
		VotesCount:    brand.VotesCount,
		DeactivatedAt: brand.DeactivatedAt,
		UpdatedAt:     brand.UpdatedAt,
		Logo:          brand.Logo,
	}).Error; err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "UpdadeBrand",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *BrandRepository) GetBrands() ([]entities.Brand, error) {
	var brandsModel []models.Brands

	result := c.gorm.Model(&models.Brands{}).Where("active =?", true).Find(&brandsModel)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetBrands",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var brands []entities.Brand
	for _, brandModel := range brandsModel {
		brands = append(brands, *brandModel.ToEntity())
	}

	return brands, nil
}
