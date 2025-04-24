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

func (u *AddMoviesListUseCase) Execute(ctx context.Context, input AddMoviesListInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	var problems []exceptions.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.Movies.ListID)
	if errGetList != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddMoviesListUseCase", "ListNotFound")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddMoviesListUseCase",
			Message:  "error getting list by ID",
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	if list.ListType != entities.MOVIE_TYPE {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.BadRequest, language.GetErrorMessage("AddMoviesListUseCase", "InvalidListType")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "AddMoviesListUseCase",
			Message:  "error adding movies to list",
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
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
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "AddMoviesListUseCase",
			Message:  "error adding movies to list",
			Error:    errGetList,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	movies, errGetMoviesByID := u.MovieRepository.GetMoviesByIDs(input.Movies.Movies)
	if errGetMoviesByID != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddMoviesListUseCase", "ErrorFetchingMovies")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddMoviesListUseCase",
			Message:  "error getting movies by IDs",
			Error:    errGetMoviesByID,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	movieIDs := []string{}

	getOldMovieIDs := list.GetItemIDs()

	movieIDs = append(movieIDs, getOldMovieIDs...)

	var items []interface{}
	for _, movie := range movies {
		items = append(items, movie)
	}

	list.AddItems(items)

	getNewMovieIDs := list.GetItemIDs()

	movieIDs = append(movieIDs, getNewMovieIDs...)

	combinations := list.GetCombinations(movieIDs)

	list.AddCombinations(combinations)

	errAddMovies := u.ListRepository.AddMovies(list)
	if errAddMovies != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("AddMoviesListUseCase", "ErrorAddingMovies")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "AddMoviesListUseCase",
			Message:  "error adding movies to list",
			Error:    errAddMovies,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Movies added successfully.",
		ContentMessage: "The movies were successfully added to the list.",
	}, nil
}
