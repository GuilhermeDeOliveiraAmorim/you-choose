package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type DirectorUseCase struct {
	DirectorRepository entity.DirectorRepositoryInterface
	MovieRepository    entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewDirectorUseCase(
	directorRepository entity.DirectorRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository  entity.FileRepositoryInterface) *DirectorUseCase {
	return &DirectorUseCase{
		DirectorRepository: directorRepository,
		MovieRepository:    movieRepository,
		FileRepository: fileRepository,
	}
}

func (directorUseCase *DirectorUseCase) Create(input InputCreateDirectorDto) (OutputCreateDirectorDto, error) {
	output := OutputCreateDirectorDto{}

	director, err := entity.NewDirector(input.Name)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := directorUseCase.DirectorRepository.Create(director); err != nil {
		return output, errors.New(err.Error())
	}

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = director.Picture
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt

	return output, nil
}

func (directorUseCase *DirectorUseCase) Find(input InputFindDirectorDto) (OutputFindDirectorDto, error) {
	output := OutputFindDirectorDto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	inputFindPicture := InputFindDirectorPictureToBase64Dto{
		DirectorId: director.ID,
	}

	picture, err := directorUseCase.FindDirectorPictureToBase64(inputFindPicture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	file, err := directorUseCase.FileRepository.Find(director.Picture)
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

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = picture.Director.Picture
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt
	output.Director.File = fileDto

	return output, nil
}

func (directorUseCase *DirectorUseCase) Delete(input InputDeleteDirectorDto) (OutputDeleteDirectorDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputDeleteDirectorDto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if director.IsDeleted {
		return output, errors.New("director previously deleted")
	}

	director.IsDeleted = true
	director.DeletedAt = timeNow

	err = directorUseCase.DirectorRepository.Update(&director)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = director.IsDeleted

	return output, nil
}

func (directorUseCase *DirectorUseCase) Update(input InputUpdateDirectorDto) (OutputUpdateDirectorDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateDirectorDto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	director.Name = input.Name
	director.Picture = input.Picture

	isValid, err := director.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	director.UpdatedAt = timeNow

	err = directorUseCase.DirectorRepository.Update(&director)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = director.Picture
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt

	return output, nil
}

func (directorUseCase *DirectorUseCase) IsDeleted(input InputIsDeletedDirectorDto) (OutputIsDeletedDirectorDto, error) {
	output := OutputIsDeletedDirectorDto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if director.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (directorUseCase *DirectorUseCase) FindAll() (OutputFindAllDirectorDto, error) {
	output := OutputFindAllDirectorDto{}

	directors, err := directorUseCase.DirectorRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, director := range directors {

		inputFindPicture := InputFindDirectorPictureToBase64Dto{
			DirectorId: director.ID,
		}

		picture, err := directorUseCase.FindDirectorPictureToBase64(inputFindPicture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		file, err := directorUseCase.FileRepository.Find(director.Picture)
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

		output.Directors = append(output.Directors, DirectorDto{
			ID:        director.ID,
			Name:      director.Name,
			Picture:   picture.Director.Picture,
			IsDeleted: director.IsDeleted,
			CreatedAt: director.CreatedAt,
			UpdatedAt: director.UpdatedAt,
			DeletedAt: director.DeletedAt,
			File: fileDto,
		})
	}

	return output, nil
}

func (directorUseCase *DirectorUseCase) AddPictureToDirector(input InputAddPictureToDirectorDto) (OutputAddPictureToDirectorDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToDirectorDto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
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

	picture, err := entity.NewFile(name, director.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := directorUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	director.Picture = picture.ID

	isValid, err := director.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	director.UpdatedAt = timeNow

	err = directorUseCase.DirectorRepository.Update(&director)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = director.Picture
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt

	return output, nil
}

func (directorUseCase *DirectorUseCase) FindDirectorPictureToBase64(input InputFindDirectorPictureToBase64Dto) (OutputFindDirectorPictureToBase64Dto, error) {
	output := OutputFindDirectorPictureToBase64Dto{}

	director, err := directorUseCase.DirectorRepository.Find(input.DirectorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := directorUseCase.FileRepository.Find(director.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64("/home/guilherme/Workspace/you-choose/cmd/upload/", picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = pictureToBase64
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt

	return output, nil
}
