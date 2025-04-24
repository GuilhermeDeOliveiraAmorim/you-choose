package usecases

import (
	"context"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type Brands struct {
	ListID string   `json:"list_id"`
	Brands []string `json:"brands"`
}

type AddBrandsListInputDTO struct {
	UserID string `json:"user_id"`
	Brands Brands `json:"add_brands_list"`
}

type AddBrandsListUseCase struct {
	ListRepository  repositories.ListRepository
	BrandRepository repositories.BrandRepository
	UserRepository  repositories.UserRepository
}

func NewAddBrandsListUseCase(
	ListRepository repositories.ListRepository,
	BrandRepository repositories.BrandRepository,
	UserRepository repositories.UserRepository,
) *AddBrandsListUseCase {
	return &AddBrandsListUseCase{
		ListRepository:  ListRepository,
		BrandRepository: BrandRepository,
		UserRepository:  UserRepository,
	}
}

func (u *AddBrandsListUseCase) Execute(ctx context.Context, input AddBrandsListInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	var problems []exceptions.ProblemDetails

	logging.NewLogger(logging.Logger{
		Context: ctx,
		TypeLog: logging.LoggerTypes.INFO,
		Layer:   logging.LoggerLayers.USECASES,
		Code:    exceptions.RFC200_CODE,
		From:    "AddBrandsListUseCase",
		Message: "starting add brands to list process",
	})

	list, errGetList := u.ListRepository.GetListByID(input.Brands.ListID)
	if errGetList != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddBrandsListUseCase", "ListNotFound")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddBrandsListUseCase",
			Message:  "error getting list by ID: " + input.Brands.ListID,
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	} else if list.ListType != entities.BRAND_TYPE {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddBrandsListUseCase", "InvalidListType")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddBrandsListUseCase",
			Message:  "error getting list by ID: " + input.Brands.ListID,
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	for _, brandID := range input.Brands.Brands {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Brand:
				if item.ID == brandID {
					problems = append(problems,
						exceptions.NewProblemDetails(
							exceptions.BadRequest,
							language.GetErrorMessage("AddBrandsListUseCase", "BrandAlreadyInList"),
						),
					)
				}
			}
		}
	}

	if len(problems) > 0 {
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "AddBrandsListUseCase",
			Message:  "error adding brands to list: " + input.Brands.ListID,
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	brands, errGetBrandsByID := u.BrandRepository.GetBrandsByIDs(input.Brands.Brands)
	if errGetBrandsByID != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddBrandsListUseCase", "BrandNotFound")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddBrandsListUseCase",
			Message:  "error getting brands by ID: " + input.Brands.ListID,
			Error:    errGetBrandsByID,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	brandIDs := []string{}

	getOldBrandIDs := list.GetItemIDs()

	brandIDs = append(brandIDs, getOldBrandIDs...)

	var items []interface{}
	for _, brand := range brands {
		items = append(items, brand)
	}

	list.AddItems(items)

	getNewBrandIDs := list.GetItemIDs()

	brandIDs = append(brandIDs, getNewBrandIDs...)

	combinations := list.GetCombinations(brandIDs)

	list.AddCombinations(combinations)

	errAddBrands := u.ListRepository.AddBrands(list)
	if errAddBrands != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddBrandsListUseCase", "ErrorAddingBrands")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddBrandsListUseCase",
			Message:  "error adding brands to list: " + input.Brands.ListID,
			Error:    errAddBrands,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Brands added successfully.",
		ContentMessage: "The brands were successfully added to the list.",
	}, nil
}
