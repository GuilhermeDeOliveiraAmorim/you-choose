package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	repositories_implementation "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type BrandFactory struct {
	CreateBrand *usecases.CreateBrandUseCase
}

func NewBrandFactory(input database.StorageInput) *BrandFactory {
	movieResository := repositories_implementation.NewBrandRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	imageRepository := repositories_implementation.NewImageRepository(input.BucketName)

	createBrand := usecases.NewCreateBrandUseCase(movieResository, userResository, imageRepository)

	return &BrandFactory{
		CreateBrand: createBrand,
	}
}
