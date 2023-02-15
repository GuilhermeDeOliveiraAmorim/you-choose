package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieUseCase struct {
	MovieRepository    entity.MovieRepositoryInterface
	ActorRepository    entity.ActorRepositoryInterface
	WriterRepository   entity.WriterRepositoryInterface
	DirectorRepository entity.DirectorRepositoryInterface
	GenreRepository    entity.GenreRepositoryInterface
}

func NewMovieUseCase(movieRepository entity.MovieRepositoryInterface, actorRepository entity.ActorRepositoryInterface, writerRepository entity.WriterRepositoryInterface, directorRepository entity.DirectorRepositoryInterface, genreRepository entity.GenreRepositoryInterface) *MovieUseCase {
	return &MovieUseCase{
		MovieRepository:    movieRepository,
		ActorRepository:    actorRepository,
		WriterRepository:   writerRepository,
		DirectorRepository: directorRepository,
		GenreRepository:    genreRepository,
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

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, err
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

	directorsIds, err := movieUseCase.MovieRepository.FindMovieDirectors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputDirectors []DirectorDto

	for _, directorId := range directorsIds {
		director, err := movieUseCase.DirectorRepository.Find(directorId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputDirectors = append(outputDirectors, DirectorDto{
			ID:        director.ID,
			Name:      director.Name,
			Picture:   director.Picture,
			IsDeleted: director.IsDeleted,
			CreatedAt: director.CreatedAt,
			UpdatedAt: director.UpdatedAt,
			DeletedAt: director.DeletedAt,
		})
	}

	genresIds, err := movieUseCase.MovieRepository.FindMovieGenres(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputGenres []GenreDto

	for _, genreId := range genresIds {
		genre, err := movieUseCase.GenreRepository.Find(genreId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputGenres = append(outputGenres, GenreDto{
			ID:        genre.ID,
			Name:      genre.Name,
			Picture:   genre.Picture,
			IsDeleted: genre.IsDeleted,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
			DeletedAt: genre.DeletedAt,
		})
	}

	output.Movie.Actors = outputActors
	output.Movie.Writers = outputWriters
	output.Movie.Directors = outputDirectors
	output.Movie.Genres = outputGenres

	return output, nil
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

func (movieUseCase *MovieUseCase) AddDirectorsToMovie(input InputAddDirectorsToMovieDto) (OutputAddDirectorsToMovieDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddDirectorsToMovieDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	directorsIds, err := movieUseCase.MovieRepository.FindMovieDirectors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var directorsMovie []entity.Director

	for _, directorId := range directorsIds {
		director, err := movieUseCase.DirectorRepository.Find(directorId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		directorsMovie = append(directorsMovie, director)
	}

	var directorsAdded []entity.Director

	for _, directorId := range input.DirectorsIds {
		director, err := movieUseCase.DirectorRepository.Find(directorId.DirectorId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		directorsAdded = append(directorsAdded, director)
	}

	for _, directorMovie := range directorsMovie {
		for position, directorAdded := range directorsAdded {
			if directorMovie.ID == directorAdded.ID {
				directorsAdded = append(directorsAdded[:position], directorsAdded[position+1:]...)
			}
		}
	}

	err = movieUseCase.MovieRepository.AddDirectorsToMovie(movie, directorsAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputDirectors []DirectorDto

	for _, director := range directorsAdded {
		outputDirectors = append(outputDirectors, DirectorDto{
			ID:        director.ID,
			Name:      director.Name,
			Picture:   director.Picture,
			IsDeleted: director.IsDeleted,
			CreatedAt: director.CreatedAt,
			UpdatedAt: director.UpdatedAt,
			DeletedAt: director.DeletedAt,
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
	output.Movie.Directors = outputDirectors

	return output, nil
}

func (movieUseCase *MovieUseCase) FindMovieDirectors(input InputFindMovieDirectorsDto) (OutputFindMovieDirectorsDto, error) {
	output := OutputFindMovieDirectorsDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	directorsIds, err := movieUseCase.MovieRepository.FindMovieDirectors(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputDirectors []DirectorDto

	for _, directorId := range directorsIds {
		director, err := movieUseCase.DirectorRepository.Find(directorId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputDirectors = append(outputDirectors, DirectorDto{
			ID:        director.ID,
			Name:      director.Name,
			Picture:   director.Picture,
			IsDeleted: director.IsDeleted,
			CreatedAt: director.CreatedAt,
			UpdatedAt: director.UpdatedAt,
			DeletedAt: director.DeletedAt,
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
	output.Movie.Directors = outputDirectors

	return output, nil
}

func (movieUseCase *MovieUseCase) AddGenresToMovie(input InputAddGenresToMovieDto) (OutputAddGenresToMovieDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddGenresToMovieDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	genresIds, err := movieUseCase.MovieRepository.FindMovieGenres(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var genresMovie []entity.Genre

	for _, genreId := range genresIds {
		genre, err := movieUseCase.GenreRepository.Find(genreId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		genresMovie = append(genresMovie, genre)
	}

	var genresAdded []entity.Genre

	for _, genreId := range input.GenresIds {
		genre, err := movieUseCase.GenreRepository.Find(genreId.GenreId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		genresAdded = append(genresAdded, genre)
	}

	for _, genreMovie := range genresMovie {
		for position, genreAdded := range genresAdded {
			if genreMovie.ID == genreAdded.ID {
				genresAdded = append(genresAdded[:position], genresAdded[position+1:]...)
			}
		}
	}

	err = movieUseCase.MovieRepository.AddGenresToMovie(movie, genresAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputGenres []GenreDto

	for _, genre := range genresAdded {
		outputGenres = append(outputGenres, GenreDto{
			ID:        genre.ID,
			Name:      genre.Name,
			Picture:   genre.Picture,
			IsDeleted: genre.IsDeleted,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
			DeletedAt: genre.DeletedAt,
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
	output.Movie.Genres = outputGenres

	return output, nil
}

func (movieUseCase *MovieUseCase) FindMovieGenres(input InputFindMovieGenresDto) (OutputFindMovieGenresDto, error) {
	output := OutputFindMovieGenresDto{}

	movie, err := movieUseCase.MovieRepository.Find(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	genresIds, err := movieUseCase.MovieRepository.FindMovieGenres(input.MovieId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputGenres []GenreDto

	for _, genreId := range genresIds {
		genre, err := movieUseCase.GenreRepository.Find(genreId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputGenres = append(outputGenres, GenreDto{
			ID:        genre.ID,
			Name:      genre.Name,
			Picture:   genre.Picture,
			IsDeleted: genre.IsDeleted,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
			DeletedAt: genre.DeletedAt,
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
	output.Movie.Genres = outputGenres

	return output, nil
}
