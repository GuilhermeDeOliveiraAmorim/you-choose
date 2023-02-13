package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieUseCase struct {
	MovieRepository entity.MovieRepositoryInterface
	ActorRepository entity.ActorRepositoryInterface
}

func NewMovieUseCase(movieRepository entity.MovieRepositoryInterface, actorRepository entity.ActorRepositoryInterface) *MovieUseCase {
	return &MovieUseCase{
		MovieRepository: movieRepository,
		ActorRepository: actorRepository,
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

	output.Movie.ID = movie.ID
	output.Movie.Title = movie.Title
	output.Movie.Synopsis = movie.Synopsis
	output.Movie.ImdbRating = movie.ImdbRating
	output.Movie.Votes = movie.Votes
	output.Movie.YouChooseRating = movie.YouChooseRating
	output.Movie.Poster = movie.Poster
	output.Movie.IsDeleted = movie.IsDeleted
	output.Movie.CreatedAt = movie.CreatedAt
	output.Movie.UpdatedAt = movie.UpdatedAt
	output.Movie.DeletedAt = movie.DeletedAt

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

func (movieUseCase *MovieUseCase) Find(input InputFindMovieDto) (OutpuFindMovieDto, error) {
	output := OutpuFindMovieDto{}

	movies, err := movieUseCase.FindAll()
	if err != nil {
		return output, err
	}

	for _, movie := range movies.Movies {
		if input.ID == movie.ID {
			output.ID = movie.ID
			output.Title = movie.Title
			output.Synopsis = movie.Synopsis
			output.ImdbRating = movie.ImdbRating
			output.Votes = movie.Votes
			output.YouChooseRating = movie.YouChooseRating
			output.Poster = movie.Poster
			output.IsDeleted = movie.IsDeleted
			output.CreatedAt = movie.CreatedAt
			output.UpdatedAt = movie.UpdatedAt
			output.DeletedAt = movie.DeletedAt
			return output, nil
		}
	}

	return output, errors.New(err.Error())
}

func (movieUseCase *MovieUseCase) AddActorsToMovie(input InputAddActorsToMovieDto) (OutputAddActorsToMovieDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddActorsToMovieDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	actors, err := movieUseCase.ActorRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, actorId := range input.ActorsIds {
		for _, actor := range actors {
			if actorId.MovieId == actor.ID {
				movie.AddActor(&actor)
			}
		}
	}

	var outputActors []ActorDto

	for _, actor := range movie.Actors {
		outputActors = append(outputActors, ActorDto{
			ID:        actor.ID,
			Name:      actor.Name,
			Picture:   actor.Picture,
			IsDeleted: actor.IsDeleted,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
			DeletedAt: actor.DeletedAt,
		})
	}

	err = movieUseCase.MovieRepository.AddActorsToMovie(movie, movie.Actors)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Movie.ID = movie.ID
	output.Movie.Title = movie.Title
	output.Movie.Synopsis = movie.Synopsis
	output.Movie.ImdbRating = movie.ImdbRating
	output.Movie.Votes = movie.Votes
	output.Movie.YouChooseRating = movie.YouChooseRating
	output.Movie.Poster = movie.Poster
	output.Movie.IsDeleted = movie.IsDeleted
	output.Movie.CreatedAt = movie.CreatedAt
	output.Movie.UpdatedAt = dateNow
	output.Movie.DeletedAt = movie.DeletedAt
	output.Movie.Actors = outputActors

	return output, nil
}

func (movieUseCase *MovieUseCase) FindMovieActors(input InputFindMovieActorsDto) (OutputFindMovieActorsDto, error) {
	output := OutputFindMovieActorsDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	actors, err := movieUseCase.MovieRepository.FindMovieActors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputActors []ActorDto

	for _, actor := range actors {
		outputActors = append(outputActors, ActorDto{
			ID:        actor.ID,
			Name:      actor.Name,
			Picture:   actor.Picture,
			IsDeleted: actor.IsDeleted,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
			DeletedAt: actor.DeletedAt,
		})
	}

	output.Movie.ID = movie.ID
	output.Movie.Title = movie.Title
	output.Movie.Synopsis = movie.Synopsis
	output.Movie.ImdbRating = movie.ImdbRating
	output.Movie.Votes = movie.Votes
	output.Movie.YouChooseRating = movie.YouChooseRating
	output.Movie.Poster = movie.Poster
	output.Movie.IsDeleted = movie.IsDeleted
	output.Movie.CreatedAt = movie.CreatedAt
	output.Movie.UpdatedAt = movie.UpdatedAt
	output.Movie.DeletedAt = movie.DeletedAt
	output.Movie.Actors = outputActors

	return output, nil
}
