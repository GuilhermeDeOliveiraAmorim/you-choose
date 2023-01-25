package usecases

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ChooserUseCase struct {
	ChooserRepository entity.ChooserRepositoryInterface
}

func NewChooserUseCase(chooserRepository entity.ChooserRepositoryInterface) *ChooserUseCase {
	return &ChooserUseCase{
		ChooserRepository: chooserRepository,
	}
}

func (c *ChooserUseCase) Create(input InputCreateChooserDto) (OutputCreateChooserDto, error) {
	chooser, err := entity.NewChooser(input.FirstName, input.LastName, input.UserName, input.Picture, input.Password)

	output := OutputCreateChooserDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	choosers, err := c.ChooserRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, existingChooser := range choosers {
		if input.UserName == existingChooser.UserName {
			return output, errors.New("username already exists")
		}
	}

	if err := c.ChooserRepository.Create(chooser); err != nil {
		return output, errors.New(err.Error())
	}

	output.ID = chooser.ID
	output.UserName = chooser.UserName
	output.Picture = chooser.Picture

	return output, nil
}

func (c *ChooserUseCase) FindAll() (OutputFindAllChooserDto, error) {
	choosers, err := c.ChooserRepository.FindAll()

	output := OutputFindAllChooserDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	choosersOutput := []OutputFindChooserDto{}

	for _, chooser := range choosers {
		choosersOutput = append(choosersOutput, OutputFindChooserDto{chooser.ID, chooser.Picture, chooser.UserName})
	}

	output = OutputFindAllChooserDto{
		Choosers: choosersOutput,
	}

	return output, nil
}

func (c *ChooserUseCase) Find(input InputFindChooserDto) (OutputFindChooserDto, error) {
	choosers, err := c.ChooserRepository.FindAll()

	output := OutputFindChooserDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, chooser := range choosers {
		if input.ID == chooser.ID {
			output.ID = chooser.ID
			output.UserName = chooser.UserName
			output.Picture = chooser.Picture
			return output, nil
		}
	}

	return output, errors.New(err.Error())
}
