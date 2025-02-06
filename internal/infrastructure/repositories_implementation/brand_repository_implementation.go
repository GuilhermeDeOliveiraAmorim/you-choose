package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
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

	if err := tx.Create(&Brands{
		ID:            brand.ID,
		Active:        brand.Active,
		CreatedAt:     brand.CreatedAt,
		UpdatedAt:     brand.UpdatedAt,
		DeactivatedAt: brand.DeactivatedAt,
		Name:          brand.Name,
		Logo:          brand.Logo,
		VotesCount:    brand.VotesCount,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *BrandRepository) GetBrandByID(brandID string) (entities.Brand, error) {
	var brandModel Brands

	result := c.gorm.Model(&Brands{}).Where("id =? AND active =?", brandID, true).First(&brandModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Brand{}, errors.New("brand not found")
		}
		return entities.Brand{}, errors.New(result.Error.Error())
	}

	return *brandModel.ToEntity(), nil
}

func (c *BrandRepository) ThisBrandExist(brandName string) (bool, error) {
	var brandModel Brands

	result := c.gorm.Model(&Brands{}).Where("name =?", brandName).First(&brandModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}

func (c *BrandRepository) GetBrandsByIDs(brandsIDs []string) ([]entities.Brand, error) {
	var brandsModel []Brands

	result := c.gorm.Model(&Brands{}).Where("id IN?", brandsIDs).Find(&brandsModel)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
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

	if err := tx.Model(&Brands{}).Where("id =?", brand.ID).Updates(Brands{
		Active:        brand.Active,
		Name:          brand.Name,
		VotesCount:    brand.VotesCount,
		DeactivatedAt: brand.DeactivatedAt,
		UpdatedAt:     brand.UpdatedAt,
		Logo:          brand.Logo,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *BrandRepository) GetBrands() ([]entities.Brand, error) {
	var brandsModel []Brands

	result := c.gorm.Model(&Brands{}).Where("active =?", true).Find(&brandsModel)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}

	var brands []entities.Brand
	for _, brandModel := range brandsModel {
		brands = append(brands, *brandModel.ToEntity())
	}

	return brands, nil
}
