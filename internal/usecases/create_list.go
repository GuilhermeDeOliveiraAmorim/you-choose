package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type List struct {
	Name     string   `json:"name"`
	Cover    string   `json:"cover"`
	ListType string   `json:"list_type"`
	Items    []string `json:"items"`
}

type CreateListInputDTO struct {
	List   List   `json:"list"`
	UserID string `json:"user_id"`
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

func (u *CreateListUseCase) Execute(input CreateListInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	if len(input.List.Items) < 2 {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "You must provide at least two items to create a list.",
				Instance: exceptions.RFC400,
			},
		}
	}

	listExists, errThisListExist := u.ListRepository.ThisListExistByName(input.List.Name)
	if errThisListExist != nil && strings.Compare(errThisListExist.Error(), "list not found") > 0 {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing list",
				Status:   500,
				Detail:   "There was a problem checking if the list already exists.",
				Instance: exceptions.RFC500,
			},
		}
	}

	if listExists {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A list with this name already exists. Please choose a different name.",
				Instance: exceptions.RFC409,
			},
		}
	}

	list, problems := entities.NewList(input.List.Name, input.List.Cover)
	if len(problems) > 0 {
		return presenters.SuccessOutputDTO{}, problems
	}

	isValidType := false

	if contains(list.GetTypes(), input.List.ListType) {
		isValidType = true
	}

	if !isValidType {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "The list type provided is not valid. Allowed types: " + strings.Join(list.GetTypes(), ", "),
				Instance: exceptions.RFC400,
			},
		}
	}

	var movies []entities.Movie
	var brands []entities.Brand

	if input.List.ListType == entities.MOVIE_TYPE {
		var err error

		list.AddType(entities.MOVIE_TYPE)

		movies, err = u.MovieRepository.GetMoviesByIDs(input.List.Items)
		if err != nil {
			return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching movies",
					Status:   500,
					Detail:   "We couldn't retrieve the movies at this time. Please try again later.",
					Instance: exceptions.RFC500,
				},
			}
		} else if len(movies) == 0 {
			return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
				{
					Type:     "Not Found",
					Title:    "Movies not found",
					Status:   404,
					Detail:   "No movies were found for the given IDs.",
					Instance: exceptions.RFC404,
				},
			}
		}

		var items []interface{}
		for _, movie := range movies {
			items = append(items, movie)
		}

		list.AddItems(items)

	} else if input.List.ListType == entities.BRAND_TYPE {
		var err error

		list.AddType(entities.BRAND_TYPE)

		brands, err = u.BrandRepository.GetBrandsByIDs(input.List.Items)
		if err != nil {
			return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching brands",
					Status:   500,
					Detail:   "We couldn't retrieve the brands at this moment. Please try again later.",
					Instance: exceptions.RFC500,
				},
			}
		} else if len(brands) == 0 {
			return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
				{
					Type:     "Not Found",
					Title:    "Brands not found",
					Status:   404,
					Detail:   "No brands were found for the given IDs.",
					Instance: exceptions.RFC404,
				},
			}
		}

		var items []interface{}
		for _, brand := range brands {
			items = append(items, brand)
		}

		list.AddItems(items)
	}

	combinations := list.GetCombinations(input.List.Items)

	list.AddCombinations(combinations)

	cover, err := u.ImageRepository.SaveImage(input.List.Cover)
	if err != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving cover",
				Status:   500,
				Detail:   "The cover image could not be saved at this time.",
				Instance: exceptions.RFC500,
			},
		}
	}

	list.AddCover(cover)

	err = u.ListRepository.CreateList(*list)
	if err != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating list",
				Status:   500,
				Detail:   "An error occurred while saving the list. Please try again later.",
				Instance: exceptions.RFC500,
			},
		}
	}

	return presenters.SuccessOutputDTO{
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
