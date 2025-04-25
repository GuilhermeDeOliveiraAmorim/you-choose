package usecases

import (
	"context"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
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

func (u *CreateBrandUseCase) Execute(ctx context.Context, input CreateBrandInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	problems := []exceptions.ProblemDetails{}

	brandExists, errThisBrandExist := u.BrandRepository.ThisBrandExist(input.Brand.Name)
	if errThisBrandExist != nil && strings.Compare(errThisBrandExist.Error(), "brand not found") > 0 {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateBrandUseCase", "ErrorFetchingBrand")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateBrandUseCase",
			Message:  "error checking if brand exists",
			Error:    errThisBrandExist,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	if brandExists {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Conflict, language.GetErrorMessage("CreateBrandUseCase", "BrandAlreadyExists")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC409_CODE,
			From:     "CreateBrandUseCase",
			Message:  "brand already exists",
			Error:    errThisBrandExist,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	brand, problems := entities.NewBrand(
		input.Brand.Name,
		input.Brand.Logo,
	)

	if len(problems) > 0 {
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "CreateBrandUseCase",
			Message:  "error creating brand",
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	logo, errSaveImage := u.ImageRepository.SaveImage(input.Brand.Logo)
	if errSaveImage != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateBrandUseCase", "ErrorSavingLogo")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateBrandUseCase",
			Message:  "error saving logo",
			Error:    errSaveImage,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	brand.AddLogo(logo)

	errCreateBrand := u.BrandRepository.CreateBrand(*brand)
	if errCreateBrand != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateBrandUseCase", "ErrorCreatingBrand")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateBrandUseCase",
			Message:  "error creating brand",
			Error:    errCreateBrand,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Brand created successfully!",
		ContentMessage: brand.Name,
	}, nil
}
