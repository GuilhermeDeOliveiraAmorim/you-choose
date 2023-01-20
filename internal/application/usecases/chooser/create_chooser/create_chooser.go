package createchooser

import (
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	chooserRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/repository"
)

type CreateChooserUseCase struct {
	ChooserRepository chooserRepository.ChooserRepositoryInterface
}

func NewCreateChooserUseCase(chooserRepository chooserRepository.ChooserRepositoryInterface) *CreateChooserUseCase {
	return &CreateChooserUseCase{
		ChooserRepository: chooserRepository,
	}
}

func (c *CreateChooserUseCase) Execute(input InputCreateChooserDto) (OutputCreateChooserDto, error) {
	newChooser, err := chooser.NewChooser(input.FirstName, input.LastName, input.UserName, input.Picture, input.Password)

	output := OutputCreateChooserDto{}

	if err != nil {
		return output, err
	}

	c.ChooserRepository.Create(newChooser)

	output.ID = newChooser.ID
	output.UserName = newChooser.UserName
	output.Picture = newChooser.Picture

	return output, nil
}
