package createdirector

import (
	director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
)

func CreateDirectorUseCase(input *InputCreateDirectorDto) *OutputCreateDirectorDto {
	if input.Name == "" {
		return nil
	}

	if input.Picture == "" {
		return nil
	}

	directorOutput, _ := director.NewDirector(input.Name, input.Picture)

	output := OutputCreateDirectorDto{
		directorOutput.ID, directorOutput.Name, directorOutput.Picture,
	}

	return &output
}
