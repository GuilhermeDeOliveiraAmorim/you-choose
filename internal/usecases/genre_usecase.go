package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type GenreUseCase struct {
	GenreRepository entity.GenreRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
}

func NewGenreUseCase(genreRepository entity.GenreRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *GenreUseCase {
	return &GenreUseCase{
		GenreRepository: genreRepository,
		MovieRepository: movieRepository,
	}
}

func (genreUseCase *GenreUseCase) Create(input InputCreateGenreDto) (OutputCreateGenreDto, error) {
	output := OutputCreateGenreDto{}

	genre, err := entity.NewGenre(input.Name, input.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	doesThisGenreAlreadyExist, err := genreUseCase.GenreRepository.FindGenreByName(input.Name)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if doesThisGenreAlreadyExist.Name != "" {
		return output, errors.New("this genre already exists")
	}

	if err := genreUseCase.GenreRepository.Create(genre); err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = genre.Picture
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt

	return output, nil
}

func (genreUseCase *GenreUseCase) Find(input InputFindGenreDto) (OutputFindGenreDto, error) {
	output := OutputFindGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = genre.Picture
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt

	return output, nil
}

func (genreUseCase *GenreUseCase) Delete(input InputDeleteGenreDto) (OutputDeleteGenreDto, error) {
	output := OutputDeleteGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if genre.IsDeleted {
		return output, errors.New("genre previously deleted")
	}

	genre.IsDeleted = true
	genre.DeletedAt = time.Now().Local().String()

	err = genreUseCase.GenreRepository.Update(&genre)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = genre.IsDeleted

	return output, nil
}

func (genreUseCase *GenreUseCase) Update(input InputUpdateGenreDto) (OutputUpdateGenreDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	genre.Name = input.Name
	genre.Picture = input.Picture

	isValid, err := genre.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	genre.UpdatedAt = timeNow

	err = genreUseCase.GenreRepository.Update(&genre)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = genre.Picture
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt

	return output, nil
}

func (genreUseCase *GenreUseCase) IsDeleted(input InputIsDeletedGenreDto) (OutputIsDeletedGenreDto, error) {
	output := OutputIsDeletedGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if genre.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (genreUseCase *GenreUseCase) FindAll() (OutputFindAllGenreDto, error) {
	output := OutputFindAllGenreDto{}

	genres, err := genreUseCase.GenreRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, genre := range genres {
		output.Genres = append(output.Genres, GenreDto{
			ID:        genre.ID,
			Name:      genre.Name,
			Picture:   genre.Picture,
			IsDeleted: genre.IsDeleted,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
			DeletedAt: genre.DeletedAt,
		})
	}

	return output, nil
}

func (genreUseCase *GenreUseCase) FindGenreByName(input InputFindGenreByNameDto) (OutputFindGenreByNameDto, error) {
	output := OutputFindGenreByNameDto{}

	genre, err := genreUseCase.GenreRepository.FindGenreByName(input.GenreName)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = genre.Picture
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt

	return output, nil
}
