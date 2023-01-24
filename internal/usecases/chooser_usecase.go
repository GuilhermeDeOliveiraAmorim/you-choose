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

	output = OutputFindAllChooserDto{
		Choosers: choosers,
	}

	return output, nil
}

func (c *ChooserUseCase) Find(input InputFindChooserDto) (entity.Chooser, error) {
	choosers, err := c.ChooserRepository.FindAll()

	var chooser entity.Chooser

	if err != nil {
		return chooser, errors.New(err.Error())
	}

	var chooserIdToFind = input.ID

	for _, chooser := range choosers {
		if chooserIdToFind == chooser.ID {
			return chooser, nil
		}
	}

	return chooser, errors.New(err.Error())
}
