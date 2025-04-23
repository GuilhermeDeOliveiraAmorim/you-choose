package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
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

type CreateMovieOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
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

func (u *CreateMovieUseCase) Execute(input CreateMovieInputDTO) (CreateMovieOutputDTO, []exceptions.ProblemDetails) {
	movieExists, errThisMovieExist := u.MovieRepository.ThisMovieExist(input.Movie.ExternalID)
	if errThisMovieExist != nil && strings.Compare(errThisMovieExist.Error(), "movie not found") > 0 {
		return CreateMovieOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing movie",
				Status:   500,
				Detail:   "An error occurred while checking if the movie already exists.",
				Instance: exceptions.RFC500,
			},
		}
	}

	if movieExists {
		return CreateMovieOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Conflict",
				Title:    "Movie already exists",
				Status:   409,
				Detail:   "A movie with the same external ID already exists. Please check the external ID and try again.",
				Instance: exceptions.RFC409,
			},
		}
	}

	movie, problems := entities.NewMovie(
		input.Movie.Name,
		input.Movie.Year,
		input.Movie.ExternalID,
	)

	if len(problems) > 0 {
		return CreateMovieOutputDTO{}, problems
	}

	poster, errSaveImage := u.ImageRepository.SaveImage(input.Movie.Poster)
	if errSaveImage != nil {
		return CreateMovieOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error saving poster",
				Status:   500,
				Detail:   "An error occurred while saving the movie poster.",
				Instance: exceptions.RFC500,
			},
		}
	}

	movie.AddPoster(poster)

	errCreateMovie := u.MovieRepository.CreateMovie(*movie)
	if errCreateMovie != nil {
		return CreateMovieOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating movie",
				Status:   500,
				Detail:   "An error occurred while creating the movie in the database.",
				Instance: exceptions.RFC500,
			},
		}
	}

	return CreateMovieOutputDTO{
		SuccessMessage: "Movie created successfully!",
		ContentMessage: "The movie '" + movie.Name + "' was created successfully.",
	}, nil
}
