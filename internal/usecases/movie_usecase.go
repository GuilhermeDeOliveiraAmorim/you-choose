package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieUseCase struct {
	MovieRepository  entity.MovieRepositoryInterface
	ActorRepository  entity.ActorRepositoryInterface
	WriterRepository entity.WriterRepositoryInterface
}

func NewMovieUseCase(movieRepository entity.MovieRepositoryInterface, actorRepository entity.ActorRepositoryInterface, writerRepository entity.WriterRepositoryInterface) *MovieUseCase {
	return &MovieUseCase{
		MovieRepository:  movieRepository,
		ActorRepository:  actorRepository,
		WriterRepository: writerRepository,
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

	actorsIds, err := movieUseCase.MovieRepository.FindMovieActors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var actorsMovie []entity.Actor

	for _, actorId := range actorsIds {
		actor, err := movieUseCase.ActorRepository.Find(actorId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		actorsMovie = append(actorsMovie, actor)
	}

	var actorsAdded []entity.Actor

	for _, actorId := range input.ActorsIds {
		actor, err := movieUseCase.ActorRepository.Find(actorId.ActorId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		actorsAdded = append(actorsAdded, actor)
	}

	for _, actorMovie := range actorsMovie {
		for position, actorAdded := range actorsAdded {
			if actorMovie.ID == actorAdded.ID {
				actorsAdded = append(actorsAdded[:position], actorsAdded[position+1:]...)
			}
		}
	}

	err = movieUseCase.MovieRepository.AddActorsToMovie(movie, actorsAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputActors []ActorDto

	for _, actor := range actorsAdded {
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

	actorsIds, err := movieUseCase.MovieRepository.FindMovieActors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputActors []ActorDto

	for _, actorId := range actorsIds {
		actor, err := movieUseCase.ActorRepository.Find(actorId)
		if err != nil {
			return output, errors.New(err.Error())
		}

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

func (movieUseCase *MovieUseCase) AddWritersToMovie(input InputAddWritersToMovieDto) (OutputAddWritersToMovieDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddWritersToMovieDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	writersIds, err := movieUseCase.MovieRepository.FindMovieWriters(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var writersMovie []entity.Writer

	for _, writerId := range writersIds {
		writer, err := movieUseCase.WriterRepository.Find(writerId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		writersMovie = append(writersMovie, writer)
	}

	var writersAdded []entity.Writer

	for _, writerId := range input.WritersIds {
		writer, err := movieUseCase.WriterRepository.Find(writerId.WriterId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		writersAdded = append(writersAdded, writer)
	}

	for _, writerMovie := range writersMovie {
		for position, writerAdded := range writersAdded {
			if writerMovie.ID == writerAdded.ID {
				writersAdded = append(writersAdded[:position], writersAdded[position+1:]...)
			}
		}
	}

	err = movieUseCase.MovieRepository.AddWritersToMovie(movie, writersAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputWriters []WriterDto

	for _, writer := range writersAdded {
		outputWriters = append(outputWriters, WriterDto{
			ID:        writer.ID,
			Name:      writer.Name,
			Picture:   writer.Picture,
			IsDeleted: writer.IsDeleted,
			CreatedAt: writer.CreatedAt,
			UpdatedAt: writer.UpdatedAt,
			DeletedAt: writer.DeletedAt,
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
	output.Movie.UpdatedAt = dateNow
	output.Movie.DeletedAt = movie.DeletedAt
	output.Movie.Writers = outputWriters

	return output, nil
}

func (movieUseCase *MovieUseCase) FindMovieWriters(input InputFindMovieWritersDto) (OutputFindMovieWritersDto, error) {
	output := OutputFindMovieWritersDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	writersIds, err := movieUseCase.MovieRepository.FindMovieWriters(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputWriters []WriterDto

	for _, writerId := range writersIds {
		writer, err := movieUseCase.WriterRepository.Find(writerId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputWriters = append(outputWriters, WriterDto{
			ID:        writer.ID,
			Name:      writer.Name,
			Picture:   writer.Picture,
			IsDeleted: writer.IsDeleted,
			CreatedAt: writer.CreatedAt,
			UpdatedAt: writer.UpdatedAt,
			DeletedAt: writer.DeletedAt,
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
	output.Movie.Writers = outputWriters

	return output, nil
}
