package usecases

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
}

func NewMovieUseCase(movieRepository entity.MovieRepositoryInterface) *MovieUseCase {
	return &MovieUseCase{
		MovieRepository: movieRepository,
	}
}

func (movieUseCase *MovieUseCase) Create(input InputCreateMovieDto) (OutputCreateMovieDto, error) {
	output := OutputCreateMovieDto{}

	movie, err := entity.NewMovie(input.Title, input.Synopsis, input.ImdbRating, input.Poster)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := movieUseCase.MovieRepository.Create(movie); err != nil {
		return output, errors.New(err.Error())
	}

	movies, err := movieUseCase.MovieRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, existingMovie := range movies {
		if input.Title == existingMovie.Title {
			return output, errors.New("movie already exists")
		}
	}

	output.Movie.ID = movie.ID
	output.Movie.Title = movie.Title
	output.Movie.Synopsis = movie.Synopsis
	output.Movie.ImdbRating = movie.ImdbRating
	output.Movie.Votes = movie.Votes
	output.Movie.YouChooseRating = movie.YouChooseRating
	output.Movie.Poster = movie.Poster
	output.Movie.CreatedAt = movie.CreatedAt
	output.Movie.UpdatedAt = movie.UpdatedAt
	output.Movie.DeletedAt = movie.DeletedAt
	output.Movie.IsDeleted = movie.IsDeleted

	return output, nil
}

func (movieUseCase *MovieUseCase) FindAll() (OutputFindAllMoviesDto, error) {
	output := OutputFindAllMoviesDto{}

	movies, err := movieUseCase.MovieRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	moviesOutput := []MovieDto{}

	for _, movie := range movies {
		moviesOutput = append(moviesOutput, MovieDto{
			ID:              movie.ID,
			Title:           movie.Title,
			Synopsis:        movie.Synopsis,
			ImdbRating:      movie.ImdbRating,
			Votes:           movie.Votes,
			YouChooseRating: movie.YouChooseRating,
			Poster:          movie.Poster,
			CreatedAt:       movie.CreatedAt,
			UpdatedAt:       movie.UpdatedAt,
			DeletedAt:       movie.DeletedAt,
			IsDeleted:       movie.IsDeleted,
		})
	}

	output.Movies = moviesOutput

	return output, nil
}