package usecases

import (
	"fmt"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type List struct {
	Name  string   `json:"name"`
	Cover string   `json:"cover"`
	Type  string   `json:"type"`
	Items []string `json:"items"`
}

type CreateListInputDTO struct {
	List   List   `json:"list"`
	UserID string `json:"user_id"`
}

type CreateListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateListUseCase struct {
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
	ImageRepository repositories.ImageRepository
	BrandRepository repositories.BrandRepository
}

func NewCreateListUseCase(
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
	ImageRepository repositories.ImageRepository,
	BrandRepository repositories.BrandRepository,
) *CreateListUseCase {
	return &CreateListUseCase{
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
		ImageRepository: ImageRepository,
		BrandRepository: BrandRepository,
	}
}

func (u *CreateListUseCase) Execute(input CreateListInputDTO) (CreateListOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	} else if !user.IsAdmin {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not an admin",
				Status:   403,
				Detail:   "User is not an admin",
				Instance: util.RFC403,
			},
		}
	}

	if len(input.List.Items) < 2 {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "At least two items must be provided.",
				Instance: util.RFC400,
			},
		}
	}

	listExists, errThisListExist := u.ListRepository.ThisListExistByName(input.List.Name)
	if errThisListExist != nil && strings.Compare(errThisListExist.Error(), "list not found") > 0 {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing list",
				Status:   500,
				Detail:   errThisListExist.Error(),
				Instance: util.RFC500,
			},
		}
	}

	if listExists {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A list with the same name already exists.",
				Instance: util.RFC409,
			},
		}
	}

	list, problems := entities.NewList(input.List.Name, input.List.Cover)
	if len(problems) > 0 {
		return CreateListOutputDTO{}, problems
	}

	isValidType := false

	if contains(list.GetTypes(), input.List.Type) {
		isValidType = true
	}

	if !isValidType {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "Invalid type. Available types: " + strings.Join(list.GetTypes(), ", "),
				Instance: util.RFC400,
			},
		}
	}

	var movies []entities.Movie
	var brands []entities.Brand

	if input.List.Type == entities.TYPE_MOVIE {
		list.AddType(entities.TYPE_MOVIE)

		var errGetMoviesByID error

		movies, errGetMoviesByID = u.MovieRepository.GetMoviesByIDs(input.List.Items)
		if errGetMoviesByID != nil {
			return CreateListOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching movies",
					Status:   500,
					Detail:   errGetMoviesByID.Error(),
					Instance: util.RFC500,
				},
			}
		} else if len(movies) == 0 {
			return CreateListOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Not Found",
					Title:    "Movies not found",
					Status:   404,
					Detail:   "Movies not found",
					Instance: util.RFC404,
				},
			}
		}

		var items []interface{}
		for _, movie := range movies {
			items = append(items, movie)
		}

		list.AddItems(items)
	} else if input.List.Type == entities.TYPE_BRAND {
		list.AddType(entities.TYPE_BRAND)

		var errGetBrandsByID error

		brands, errGetBrandsByID = u.BrandRepository.GetBrandsByIDs(input.List.Items)
		if errGetBrandsByID != nil {
			return CreateListOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching brands",
					Status:   500,
					Detail:   errGetBrandsByID.Error(),
					Instance: util.RFC500,
				},
			}
		} else if len(brands) == 0 {
			return CreateListOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Not Found",
					Title:    "Brands not found",
					Status:   404,
					Detail:   "Brands not found",
					Instance: util.RFC404,
				},
			}
		}

		var items []interface{}
		for _, movie := range brands {
			items = append(items, movie)
		}

		list.AddItems(items)
	}

	combinations, errGetCombinations := list.GetCombinations(input.List.Items)
	if len(errGetCombinations) > 0 {
		return CreateListOutputDTO{}, errGetCombinations
	}

	list.AddCombinations(combinations)

	cover, errSaveImage := u.ImageRepository.SaveImage(input.List.Cover)
	if errSaveImage != nil {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving cover",
				Status:   500,
				Detail:   errSaveImage.Error(),
				Instance: util.RFC500,
			},
		}
	}

	list.AddCover(cover)

	fmt.Println("list.TypeList: ", list.TypeList)

	errCreateList := u.ListRepository.CreateList(*list)
	if errCreateList != nil {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating list",
				Status:   500,
				Detail:   errCreateList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return CreateListOutputDTO{
		SuccessMessage: "List created successfully!",
		ContentMessage: list.Name,
	}, nil
}

func contains(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}
