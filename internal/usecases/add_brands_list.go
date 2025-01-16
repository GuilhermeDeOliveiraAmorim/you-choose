package usecases

import (
	"fmt"

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
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	} else if !user.IsAdmin {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not an admin",
				Status:   403,
				Detail:   "User is not an admin",
				Instance: util.RFC403,
			},
		}
	}

	var problems []util.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.Brands.ListID)
	if errGetList != nil {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	} else if list.ListType != entities.BRAND_TYPE {
		return AddBrandsListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   "List type must be 'brand'",
				Instance: util.RFC400,
			},
		}
	}

	for _, brandID := range input.Brands.Brands {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Brand:
				if item.ID == brandID {
					problems = append(problems,
						util.ProblemDetails{
							Type:     "Validation Error",
							Title:    "Brand already in list",
							Status:   400,
							Detail:   fmt.Sprintf("Brand with ID %s already exists in the list.", brandID),
							Instance: util.RFC400,
						},
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
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching brands",
				Status:   500,
				Detail:   errGetBrandsByID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	brandIDs := []string{}

	getOldBrandIDs, errGetBrandIDs := list.GetItemIDs()
	if len(errGetBrandIDs) > 0 {
		return AddBrandsListOutputDTO{}, problems
	}

	brandIDs = append(brandIDs, getOldBrandIDs...)

	var items []interface{}
	for _, brand := range brands {
		items = append(items, brand)
	}

	list.AddItems(items)

	getNewBrandIDs, errGetBrandIDs := list.GetItemIDs()
	if len(errGetBrandIDs) > 0 {
		return AddBrandsListOutputDTO{}, problems
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
			{
				Type:     "Internal Server Error",
				Title:    "Error adding brands to list",
				Status:   500,
				Detail:   errAddBrands.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return AddBrandsListOutputDTO{
		SuccessMessage: "Brands added successfully.",
		ContentMessage: "The brands were successfully added to the list.",
	}, nil
}
