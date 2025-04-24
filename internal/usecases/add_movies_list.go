package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type Movies struct {
	ListID string   `json:"list_id"`
	Movies []string `json:"movies"`
}

type AddMoviesListInputDTO struct {
	UserID string `json:"user_id"`
	Movies Movies `json:"add_movies_list"`
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

func (u *AddMoviesListUseCase) Execute(input AddMoviesListInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	var problems []exceptions.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.Movies.ListID)
	if errGetList != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(
				exceptions.InternalServerError,
				language.GetErrorMessage("AddMoviesListUseCase", "ListNotFound"),
			),
		}
	}

	if list.ListType != entities.MOVIE_TYPE {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(
				exceptions.BadRequest,
				language.GetErrorMessage("AddMoviesListUseCase", "InvalidListType"),
			),
		}
	}

	for _, movieID := range input.Movies.Movies {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Movie:
				if item.ID == movieID {
					problems = append(problems,
						exceptions.NewProblemDetails(
							exceptions.BadRequest,
							language.GetErrorMessage("AddMoviesListUseCase", "MovieAlreadyInList"),
						),
					)
				}
			}
		}
	}

	if len(problems) > 0 {
		return presenters.SuccessOutputDTO{}, problems
	}

	movies, errGetMoviesByID := u.MovieRepository.GetMoviesByIDs(input.Movies.Movies)
	if errGetMoviesByID != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(
				exceptions.InternalServerError,
				language.GetErrorMessage("AddMoviesListUseCase", "ErrorFetchingMovies"),
			),
		}
	}

	movieIDs := []string{}

	getOldMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return presenters.SuccessOutputDTO{}, errGetMovieIDs
	}
	movieIDs = append(movieIDs, getOldMovieIDs...)

	var items []interface{}
	for _, movie := range movies {
		items = append(items, movie)
	}

	list.AddItems(items)

	getNewMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return presenters.SuccessOutputDTO{}, errGetMovieIDs
	}

	movieIDs = append(movieIDs, getNewMovieIDs...)

	combinations := list.GetCombinations(movieIDs)

	list.AddCombinations(combinations)

	errAddMovies := u.ListRepository.AddMovies(list)
	if errAddMovies != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(
				exceptions.InternalServerError,
				language.GetErrorMessage("AddMoviesListUseCase", "ErrorAddingMovies"),
			),
		}
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Movies added successfully.",
		ContentMessage: "The movies were successfully added to the list.",
	}, nil
}
