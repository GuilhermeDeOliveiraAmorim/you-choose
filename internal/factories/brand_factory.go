package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type BrandFactory struct {
	CreateBrand *usecases.CreateBrandUseCase
}

func NewBrandFactory(input util.ImputFactory) *BrandFactory {
	movieResository := repositories_implementation.NewBrandRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	imageRepository := repositories_implementation.NewImageRepository(input.BucketName)

	createBrand := usecases.NewCreateBrandUseCase(movieResository, userResository, imageRepository)

	return &BrandFactory{
		CreateBrand: createBrand,
	}
}
