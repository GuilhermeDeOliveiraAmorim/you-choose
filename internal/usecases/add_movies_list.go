package usecases

import (
	"strings"

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
		if strings.Compare(err.Error(), "user not found") == 0 {
			return AddMoviesListOutputDTO{}, []util.ProblemDetails{
				util.NewForbiddenError(
					util.GetErrorMessage("AddMoviesListUseCase", "UserNotFound", "Title"),
					util.GetErrorMessage("AddMoviesListUseCase", "UserNotFound", "Detail"),
				),
			}
		}

		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			util.NewNotFoundError(
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotFound", "Title"),
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotFound", "Detail"),
			),
		}
	}

	if !user.Active {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			util.NewForbiddenError(
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotActive", "Title"),
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotActive", "Detail"),
			),
		}
	}

	if !user.IsAdmin {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			util.NewForbiddenError(
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotAdmin", "Title"),
				util.GetErrorMessage("AddMoviesListUseCase", "UserNotAdmin", "Detail"),
			),
		}
	}

	var problems []util.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.Movies.ListID)
	if errGetList != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			util.NewInternalServerError(
				util.GetErrorMessage("AddMoviesListUseCase", "ListNotFound", "Title"),
				errGetList.Error(),
			),
		}
	}

	if list.ListType != entities.MOVIE_TYPE {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			util.NewBadRequestError(
				util.GetErrorMessage("AddMoviesListUseCase", "InvalidListType", "Title"),
				util.GetErrorMessage("AddMoviesListUseCase", "InvalidListType", "Detail"),
			),
		}
	}

	for _, movieID := range input.Movies.Movies {
		for _, item := range list.Items {
			switch item := item.(type) {
			case entities.Movie:
				if item.ID == movieID {
					problems = append(problems,
						util.NewBadRequestError(
							util.GetErrorMessage("AddMoviesListUseCase", "MovieAlreadyInList", "Title"),
							util.GetErrorMessage("AddMoviesListUseCase", "MovieAlreadyInList", "Detail"),
						),
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
			util.NewInternalServerError(
				util.GetErrorMessage("AddMoviesListUseCase", "ErrorFetchingMovies", "Title"),
				errGetMoviesByID.Error(),
			),
		}
	}

	movieIDs := []string{}

	getOldMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return AddMoviesListOutputDTO{}, errGetMovieIDs
	}
	movieIDs = append(movieIDs, getOldMovieIDs...)

	var items []interface{}
	for _, movie := range movies {
		items = append(items, movie)
	}

	list.AddItems(items)

	getNewMovieIDs, errGetMovieIDs := list.GetItemIDs()
	if len(errGetMovieIDs) > 0 {
		return AddMoviesListOutputDTO{}, errGetMovieIDs
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
			util.NewInternalServerError(
				util.GetErrorMessage("AddMoviesListUseCase", "ErrorAddingMovies", "Title"),
				errAddMovies.Error(),
			),
		}
	}

	return AddMoviesListOutputDTO{
		SuccessMessage: "Movies added successfully.",
		ContentMessage: "The movies were successfully added to the list.",
	}, nil
}
