package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type GenreUseCase struct {
	GenreRepository entity.GenreRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewGenreUseCase(
	genreRepository entity.GenreRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *GenreUseCase {
	return &GenreUseCase{
		GenreRepository: genreRepository,
		MovieRepository: movieRepository,
		FileRepository:  fileRepository,
	}
}

func (genreUseCase *GenreUseCase) Create(input InputCreateGenreDto) (OutputCreateGenreDto, error) {
	output := OutputCreateGenreDto{}

	genre, err := entity.NewGenre(input.Name)
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

	inputFindPicture := InputFindGenrePictureToBase64Dto{
		GenreId: genre.ID,
	}

	picture, err := genreUseCase.FindGenrePictureToBase64(inputFindPicture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	file, err := genreUseCase.FileRepository.Find(genre.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fileDto := FileDto{
		ID:           file.ID,
		EntityId:     file.EntityId,
		Name:         file.Name,
		Size:         file.Size,
		Extension:    file.Extension,
		AverageColor: file.AverageColor,
		IsDeleted:    file.IsDeleted,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
		DeletedAt:    file.DeletedAt,
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = picture.Genre.Picture
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt
	output.Genre.File = fileDto

	return output, nil
}

func (genreUseCase *GenreUseCase) Delete(input InputDeleteGenreDto) (OutputDeleteGenreDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputDeleteGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if genre.IsDeleted {
		return output, errors.New("genre previously deleted")
	}

	genre.IsDeleted = true
	genre.DeletedAt = timeNow

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

		inputFindPicture := InputFindGenrePictureToBase64Dto{
			GenreId: genre.ID,
		}

		picture, err := genreUseCase.FindGenrePictureToBase64(inputFindPicture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		file, err := genreUseCase.FileRepository.Find(genre.Picture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		fileDto := FileDto{
			ID:           file.ID,
			EntityId:     file.EntityId,
			Name:         file.Name,
			Size:         file.Size,
			Extension:    file.Extension,
			AverageColor: file.AverageColor,
			IsDeleted:    file.IsDeleted,
			CreatedAt:    file.CreatedAt,
			UpdatedAt:    file.UpdatedAt,
			DeletedAt:    file.DeletedAt,
		}

		output.Genres = append(output.Genres, GenreDto{
			ID:        genre.ID,
			Name:      genre.Name,
			Picture:   picture.Genre.Picture,
			IsDeleted: genre.IsDeleted,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
			DeletedAt: genre.DeletedAt,
			File:      fileDto,
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

func (genreUseCase *GenreUseCase) AddPictureToGenre(input InputAddPictureToGenreDto) (OutputAddPictureToGenreDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToGenreDto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	extension := strings.Replace(filepath.Ext(input.File.Handler.Filename), ".", "", -1)
	if strings.ToLower(extension) != "jpeg" && strings.ToLower(extension) != "jpg" {
		return output, errors.New("format not allowed")
	}

	_, name, size, extension, err := MoveFile(input.File.File, input.File.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	colorAverage, err := PictureAverageColor(name, extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := entity.NewFile(name, genre.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := genreUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	genre.Picture = picture.ID

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

func (genreUseCase *GenreUseCase) FindGenrePictureToBase64(input InputFindGenrePictureToBase64Dto) (OutputFindGenrePictureToBase64Dto, error) {
	output := OutputFindGenrePictureToBase64Dto{}

	genre, err := genreUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := genreUseCase.FileRepository.Find(genre.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64(dotenv, picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = genre.ID
	output.Genre.Name = genre.Name
	output.Genre.Picture = pictureToBase64
	output.Genre.IsDeleted = genre.IsDeleted
	output.Genre.CreatedAt = genre.CreatedAt
	output.Genre.UpdatedAt = genre.UpdatedAt
	output.Genre.DeletedAt = genre.DeletedAt

	return output, nil
}
