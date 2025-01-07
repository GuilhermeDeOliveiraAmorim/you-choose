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
	UserID     string `json:"user_id"`
}

type CreateMovieOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateMovieUseCase struct {
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
}

func NewCreateMovieUseCase(
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
) *CreateMovieUseCase {
	return &CreateMovieUseCase{
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
	}
}

func (u *CreateMovieUseCase) Execute(input CreateMovieInputDTO) (CreateMovieOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	} else if !user.IsAdmin {
		return CreateMovieOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not an admin",
				Status:   403,
				Detail:   "User is not an admin",
				Instance: util.RFC403,
			},
		}
	}

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
