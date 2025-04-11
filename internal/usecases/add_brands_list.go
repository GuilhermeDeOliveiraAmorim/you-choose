package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Brands struct {
	ListID string   `json:"list_id"`
	Brands []string `json:"brands"`
}

type AddBrandsListInputDTO struct {
	UserID string `json:"user_id"`
	Brands Brands `json:"add_brands_list"`
}

type AddBrandsListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
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

func (u *AddBrandsListUseCase) Execute(input AddBrandsListInputDTO) (AddBrandsListOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.NotFound,
				util.GetErrorMessage("AddBrandsListUseCase", "UserNotFound"),
			),
		}
	} else if !user.Active {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.Forbidden,
				util.GetErrorMessage("AddBrandsListUseCase", "UserNotActive"),
			),
		}
	} else if !user.IsAdmin {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.Forbidden,
				util.GetErrorMessage("AddBrandsListUseCase", "UserNotAdmin"),
			),
		}
	}

	var problems []util.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.Brands.ListID)
	if errGetList != nil {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.InternalServerError,
				util.GetErrorMessage("AddBrandsListUseCase", "ListNotFound"),
			),
		}
	} else if list.ListType != entities.BRAND_TYPE {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.InternalServerError,
				util.GetErrorMessage("AddBrandsListUseCase", "InvalidListType"),
			),
		}
	}

	for _, brandID := range input.Brands.Brands {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Brand:
				if item.ID == brandID {
					problems = append(problems,
						util.NewProblemDetails(
							util.BadRequest,
							util.GetErrorMessage("AddBrandsListUseCase", "BrandAlreadyInList"),
						),
					)
				}
			}
		}
	}

	if len(problems) > 0 {
		return AddBrandsListOutputDTO{}, problems
	}

	brands, errGetBrandsByID := u.BrandRepository.GetBrandsByIDs(input.Brands.Brands)
	if errGetBrandsByID != nil {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.InternalServerError,
				util.GetErrorMessage("AddBrandsListUseCase", "ErrorFetchingBrands"),
			),
		}
	}

	brandIDs := []string{}

	getOldBrandIDs, errGetBrandIDs := list.GetItemIDs()
	if len(errGetBrandIDs) > 0 {
		return AddBrandsListOutputDTO{}, errGetBrandIDs
	}

	brandIDs = append(brandIDs, getOldBrandIDs...)

	var items []interface{}
	for _, brand := range brands {
		items = append(items, brand)
	}

	list.AddItems(items)

	getNewBrandIDs, errGetBrandIDs := list.GetItemIDs()
	if len(errGetBrandIDs) > 0 {
		return AddBrandsListOutputDTO{}, errGetBrandIDs
	}

	brandIDs = append(brandIDs, getNewBrandIDs...)

	combinations, errGetCombinations := list.GetCombinations(brandIDs)
	if len(errGetCombinations) > 0 {
		return AddBrandsListOutputDTO{}, errGetCombinations
	}

	list.AddCombinations(combinations)

	errAddBrands := u.ListRepository.AddBrands(list)
	if errAddBrands != nil {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			util.NewProblemDetails(
				util.InternalServerError,
				util.GetErrorMessage("AddBrandsListUseCase", "ErrorAddingBrands"),
			),
		}
	}

	return AddBrandsListOutputDTO{
		SuccessMessage: "Brands added successfully.",
		ContentMessage: "The brands were successfully added to the list.",
	}, nil
}
