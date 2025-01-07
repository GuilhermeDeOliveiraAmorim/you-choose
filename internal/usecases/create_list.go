package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type CreateListInputDTO struct {
	Name   string   `json:"name"`
	Movies []string `json:"movies"`
	UserID string   `json:"user_id"`
}

type CreateListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateListUseCase struct {
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
}

func NewCreateListUseCase(
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
) *CreateListUseCase {
	return &CreateListUseCase{
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
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
	}

	if len(input.Movies) < 2 {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "At least two movies must be provided.",
				Instance: util.RFC400,
			},
		}
	}

	listExists, errThisListExist := u.ListRepository.ThisListExistByName(input.Name)
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

	movies, errGetMoviesByID := u.MovieRepository.GetMoviesByIDs(input.Movies)
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
	}

	list, problems := entities.NewList(input.Name)
	if len(problems) > 0 {
		return CreateListOutputDTO{}, problems
	}

	list.AddMovies(movies)

	combinations, errGetCombinations := list.GetCombinations(input.Movies)
	if len(errGetCombinations) > 0 {
		return CreateListOutputDTO{}, errGetCombinations
	}

	list.AddCombinations(combinations)

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
