package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type DirectorUseCase struct {
	DirectorRepository entity.DirectorRepositoryInterface
	MovieRepository    entity.MovieRepositoryInterface
}

func NewDirectorUseCase(directorRepository entity.DirectorRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *DirectorUseCase {
	return &DirectorUseCase{
		DirectorRepository: directorRepository,
		MovieRepository:    movieRepository,
	}
}

func (directorUseCase *DirectorUseCase) Create(input InputCreateDirectorDto) (OutputCreateDirectorDto, error) {
	output := OutputCreateDirectorDto{}

	director, err := entity.NewDirector(input.Name, input.Picture)
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

	output.Director.ID = director.ID
	output.Director.Name = director.Name
	output.Director.Picture = director.Picture
	output.Director.IsDeleted = director.IsDeleted
	output.Director.CreatedAt = director.CreatedAt
	output.Director.UpdatedAt = director.UpdatedAt
	output.Director.DeletedAt = director.DeletedAt

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
		output.Directors = append(output.Directors, DirectorDto{
			ID:        director.ID,
			Name:      director.Name,
			Picture:   director.Picture,
			IsDeleted: director.IsDeleted,
			CreatedAt: director.CreatedAt,
			UpdatedAt: director.UpdatedAt,
			DeletedAt: director.DeletedAt,
		})
	}

	return output, nil
}
