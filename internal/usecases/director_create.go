package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type DirectorUseCase struct {
	DirectorRepository entity.DirectorRepositoryInterface
}

func NewDirectorUseCase(directorRepository entity.DirectorRepositoryInterface) *DirectorUseCase {
	return &DirectorUseCase{
		DirectorRepository: directorRepository,
	}
}

func (d *DirectorUseCase) Create(input InputCreateDirectorDto) (OutputCreateDirectorDto, error) {
	director, err := entity.NewDirector(input.Name, input.Picture)

	output := OutputCreateDirectorDto{}

	if err != nil {
		return output, err
	}

	if err := d.DirectorRepository.Create(director); err != nil {
		return output, err
	}

	output.ID = director.ID
	output.Name = director.Name
	output.Picture = director.Picture

	return output, nil
}
