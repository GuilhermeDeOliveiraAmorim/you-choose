package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type CreateMovieInputDTO struct {
	Name       string `json:"name"`
	Year       int64  `json:"year"`
	Poster     string `json:"poster"`
	ExternalID string `json:"external_id"`
}

type CreateMovieOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateMovieUseCase struct {
	MovieRepository repositories.MovieRepository
}

func NewCreateMovieUseCase(
	MovieRepository repositories.MovieRepository,
) *CreateMovieUseCase {
	return &CreateMovieUseCase{
		MovieRepository: MovieRepository,
	}
}

func (u *CreateMovieUseCase) Execute(input CreateMovieInputDTO) (CreateMovieOutputDTO, []util.ProblemDetails) {
	movieExists, errThisMovieExist := u.MovieRepository.ThisMovieExist(input.ExternalID)
	if errThisMovieExist != nil && strings.Compare(errThisMovieExist.Error(), "movie not found") > 0 {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing movie",
				Status:   500,
				Detail:   errThisMovieExist.Error(),
				Instance: util.RFC500,
			},
		}
	}

	if movieExists {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A movie with the same external ID already exists.",
				Instance: util.RFC409,
			},
		}
	}

	movie, problems := entities.NewMovie(
		input.Name,
		input.Year,
		input.ExternalID,
	)

	if len(problems) > 0 {
		return CreateMovieOutputDTO{}, problems
	}

	posterID, errSavePoster := u.MovieRepository.SavePoster(input.Poster)
	if errSavePoster != nil {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving poster",
				Status:   500,
				Detail:   errSavePoster.Error(),
				Instance: util.RFC500,
			},
		}
	}

	movie.UpdatePoster(posterID)

	errCreateMovie := u.MovieRepository.CreateMovie(*movie)
	if errCreateMovie != nil {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating movie",
				Status:   500,
				Detail:   errCreateMovie.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return CreateMovieOutputDTO{
		SuccessMessage: "Movie created successfully!",
		ContentMessage: movie.Name,
	}, nil
}
