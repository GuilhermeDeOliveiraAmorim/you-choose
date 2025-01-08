package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Brand struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type CreateBrandInputDTO struct {
	UserID string `json:"user_id"`
	Brand  Brand  `json:"brand"`
}

type CreateBrandOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
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

func (u *CreateBrandUseCase) Execute(input CreateBrandInputDTO) (CreateBrandOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	} else if !user.IsAdmin {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not an admin",
				Status:   403,
				Detail:   "User is not an admin",
				Instance: util.RFC403,
			},
		}
	}

	brandExists, errThisBrandExist := u.BrandRepository.ThisBrandExist(input.Brand.Name)
	if errThisBrandExist != nil && strings.Compare(errThisBrandExist.Error(), "brand not found") > 0 {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing brand",
				Status:   500,
				Detail:   errThisBrandExist.Error(),
				Instance: util.RFC500,
			},
		}
	}

	if brandExists {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A brand with the same external ID already exists.",
				Instance: util.RFC409,
			},
		}
	}

	brand, problems := entities.NewBrand(
		input.Brand.Name,
		input.Brand.Logo,
	)

	if len(problems) > 0 {
		return CreateBrandOutputDTO{}, problems
	}

	logo, errSaveImage := u.ImageRepository.SaveImage(input.Brand.Logo)
	if errSaveImage != nil {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving logo",
				Status:   500,
				Detail:   errSaveImage.Error(),
				Instance: util.RFC500,
			},
		}
	}

	brand.AddLogo(logo)

	errCreateBrand := u.BrandRepository.CreateBrand(*brand)
	if errCreateBrand != nil {
		return CreateBrandOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating brand",
				Status:   500,
				Detail:   errCreateBrand.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return CreateBrandOutputDTO{
		SuccessMessage: "Brand created successfully!",
		ContentMessage: brand.Name,
	}, nil
}
