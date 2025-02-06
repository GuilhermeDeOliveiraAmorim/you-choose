package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type BrandRepository interface {
	CreateBrand(brand entities.Brand) error
	GetBrandByID(brandID string) (entities.Brand, error)
	ThisBrandExist(brandName string) (bool, error)
	GetBrandsByIDs(brandsIDs []string) ([]entities.Brand, error)
	UpdadeBrand(brand entities.Brand) error
	GetBrands() ([]entities.Brand, error)
}
