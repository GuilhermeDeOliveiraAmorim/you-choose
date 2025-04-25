package usecases

import (
	"errors"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"golang.org/x/net/context"
)

type Movie struct {
	Name       string `json:"name"`
	Year       int64  `json:"year"`
	Poster     string `json:"poster"`
	ExternalID string `json:"external_id"`
}

type CreateMovieInputDTO struct {
	UserID string `json:"user_id"`
	Movie  Movie  `json:"movie"`
}

type CreateMovieUseCase struct {
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
	ImageRepository repositories.ImageRepository
}

func NewCreateMovieUseCase(
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
	ImageRepository repositories.ImageRepository,
) *CreateMovieUseCase {
	return &CreateMovieUseCase{
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
		ImageRepository: ImageRepository,
	}
}

func (u *CreateMovieUseCase) Execute(ctx context.Context, input CreateMovieInputDTO) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	problems := []exceptions.ProblemDetails{}

	movieExists, errThisMovieExist := u.MovieRepository.ThisMovieExist(input.Movie.ExternalID)
	if errThisMovieExist != nil && strings.Compare(errThisMovieExist.Error(), "movie not found") > 0 {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateMovieUseCase", "ErrorFetchingExistingMovie")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateMovieUseCase",
			Message:  "error checking if movie exists",
			Error:    errThisMovieExist,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	if movieExists {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Conflict, language.GetErrorMessage("CreateMovieUseCase", "MovieAlreadyExists")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC409_CODE,
			From:     "CreateMovieUseCase",
			Message:  "error checking if movie exists",
			Error:    errThisMovieExist,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	movie, problems := entities.NewMovie(
		input.Movie.Name,
		input.Movie.Year,
		input.Movie.ExternalID,
	)

	if len(problems) > 0 {
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "CreateMovieUseCase",
			Message:  "error creating movie",
			Error:    errors.New("error creating movie"),
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	poster, errSaveImage := u.ImageRepository.SaveImage(input.Movie.Poster)
	if errSaveImage != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateMovieUseCase", "ErrorSavingPoster")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateMovieUseCase",
			Message:  "error saving movie poster",
			Error:    errSaveImage,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	movie.AddPoster(poster)

	errCreateMovie := u.MovieRepository.CreateMovie(*movie)
	if errCreateMovie != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateMovieUseCase", "ErrorCreatingMovie")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateMovieUseCase",
			Message:  "error creating movie",
			Error:    errCreateMovie,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "Movie created successfully!",
		ContentMessage: "The movie '" + movie.Name + "' was created successfully.",
	}, nil
}
