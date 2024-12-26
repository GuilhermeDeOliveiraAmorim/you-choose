package usecases

import (
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
				Type:     "https://httpstatuses.com/500",
				Title:    "Internal Server Error",
				Status:   500,
				Detail:   "An unexpected error occurred while saving the poster.",
				Instance: util.RFC500,
			},
		}
	}

	movie.UpdatePoster(posterID)

	errCreateMovie := u.MovieRepository.CreateMovie(*movie)
	if errCreateMovie != nil {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "https://httpstatuses.com/500",
				Title:    "Internal Server Error",
				Status:   500,
				Detail:   "An unexpected error occurred while creating the movie.",
				Instance: util.RFC500,
			},
		}
	}

	return CreateMovieOutputDTO{
		SuccessMessage: "Movie created successfully!",
		ContentMessage: movie.Name,
	}, nil
}
