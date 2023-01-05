package createchooser

import (
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
)

func CreateChooserUseCase(input *InputCreateChooserDto) *OutputCreateChooserDto {
	if input.FirstName == "" {
		return nil
	}

	if input.LastName == "" {
		return nil
	}

	if input.UserName == "" {
		return nil
	}

	if input.Picture == "" {
		return nil
	}

	if input.Password == "" {
		return nil
	}

	chooserOutput, _ := chooser.NewChooser(input.FirstName, input.LastName, input.UserName, input.Picture, input.Password)

	output := OutputCreateChooserDto{
		chooserOutput.ID,
		chooserOutput.UserName,
		chooserOutput.Picture,
	}

	return &output
}
