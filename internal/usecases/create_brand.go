package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type Brand struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type CreateBrandInputDTO struct {
	UserID string `json:"user_id"`
	Brand  Brand  `json:"brand"`
}

type CreateBrandUseCase struct {
	BrandRepository repositories.BrandRepository
	UserRepository  repositories.UserRepository
	ImageRepository repositories.ImageRepository
}

func NewCreateBrandUseCase(
	BrandRepository repositories.BrandRepository,
	UserRepository repositories.UserRepository,
	ImageRepository repositories.ImageRepository,
) *CreateBrandUseCase {
	return &CreateBrandUseCase{
		BrandRepository: BrandRepository,
		UserRepository:  UserRepository,
		ImageRepository: ImageRepository,
	}
}

func (u *CreateBrandUseCase) Execute(input CreateBrandInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	brandExists, errThisBrandExist := u.BrandRepository.ThisBrandExist(input.Brand.Name)
	if errThisBrandExist != nil && strings.Compare(errThisBrandExist.Error(), "brand not found") > 0 {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing brand",
				Status:   500,
				Detail:   "An unexpected error occurred while verifying existing brand data.",
				Instance: exceptions.RFC500,
			},
		}
	}

	if brandExists {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A brand with the same name already exists in the system.",
				Instance: exceptions.RFC409,
			},
		}
	}

	brand, problems := entities.NewBrand(
		input.Brand.Name,
		input.Brand.Logo,
	)

	if len(problems) > 0 {
		return presenters.SuccessOutputDTO{}, problems
	}

	logo, errSaveImage := u.ImageRepository.SaveImage(input.Brand.Logo)
	if errSaveImage != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving logo",
				Status:   500,
				Detail:   "We encountered an issue while saving the brand's logo. Please try again later.",
				Instance: exceptions.RFC500,
			},
		}
	}

	brand.AddLogo(logo)

	errCreateBrand := u.BrandRepository.CreateBrand(*brand)
	if errCreateBrand != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating brand",
				Status:   500,
				Detail:   "Something went wrong while creating the brand. Please contact support if the issue persists.",
				Instance: exceptions.RFC500,
			},
		}
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Brand created successfully!",
		ContentMessage: brand.Name,
	}, nil
}
