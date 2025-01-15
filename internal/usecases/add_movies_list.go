package usecases

import (
	"fmt"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Movies struct {
	ListID string   `json:"list_id"`
	Movies []string `json:"movies"`
}

type AddMoviesListInputDTO struct {
	UserID string `json:"user_id"`
	Movies Movies `json:"add_movies_list"`
}

type AddMoviesListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type AddMoviesListUseCase struct {
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
}

func NewAddMoviesListUseCase(
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
) *AddMoviesListUseCase {
	return &AddMoviesListUseCase{
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
	}
}

func (u *AddMoviesListUseCase) Execute(input AddMoviesListInputDTO) (AddMoviesListOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	} else if !user.IsAdmin {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
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

	list, errGetList := u.ListRepository.GetListByID(input.Movies.ListID)
	if errGetList != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	} else if list.ListType != entities.MOVIE_TYPE {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   "List type must be 'movie'",
				Instance: util.RFC400,
			},
		}
	}

	for _, movieID := range input.Movies.Movies {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Movie:
				if item.ID == movieID {
					problems = append(problems,
						util.ProblemDetails{
							Type:     "Validation Error",
							Title:    "Movie already in list",
							Status:   400,
							Detail:   fmt.Sprintf("Movie with ID %s already exists in the list.", movieID),
							Instance: util.RFC400,
						},
					)
				}
			}
		}
	}

	if len(problems) > 0 {
		return AddMoviesListOutputDTO{}, problems
	}

	movies, errGetMoviesByID := u.MovieRepository.GetMoviesByIDs(input.Movies.Movies)
	if errGetMoviesByID != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching movies",
				Status:   500,
				Detail:   errGetMoviesByID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	movieIDs := []string{}

	getOldMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return AddMoviesListOutputDTO{}, problems
	}

	movieIDs = append(movieIDs, getOldMovieIDs...)

	var items []interface{}
	for _, movie := range movies {
		items = append(items, movie)
	}

	list.AddItems(items)

	getNewMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return AddMoviesListOutputDTO{}, problems
	}

	movieIDs = append(movieIDs, getNewMovieIDs...)

	combinations, errGetCombinations := list.GetCombinations(movieIDs)
	if len(errGetCombinations) > 0 {
		return AddMoviesListOutputDTO{}, errGetCombinations
	}

	list.AddCombinations(combinations)

	errAddMovies := u.ListRepository.AddMovies(list)
	if errAddMovies != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error adding movies to list",
				Status:   500,
				Detail:   errAddMovies.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return AddMoviesListOutputDTO{
		SuccessMessage: "Movies added successfully.",
		ContentMessage: "The movies were successfully added to the list.",
	}, nil
}
