package usecases

import (
	"errors"
	"fmt"
	"time"

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

func (chooserUseCase *ChooserUseCase) Create(input InputCreateChooserDto) (OutputCreateChooserDto, error) {
	chooser, err := entity.NewChooser(input.FirstName, input.LastName, input.UserName, input.Picture, input.Password)

	output := OutputCreateChooserDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	choosers, err := chooserUseCase.ChooserRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, existingChooser := range choosers {
		if input.UserName == existingChooser.UserName {
			return output, errors.New("username already exists")
		}
	}

	if err := chooserUseCase.ChooserRepository.Create(chooser); err != nil {
		return output, errors.New(err.Error())
	}

	output.ID = chooser.ID
	output.UserName = chooser.UserName
	output.Picture = chooser.Picture

	return output, nil
}

func (chooserUseCase *ChooserUseCase) FindAll() (OutputFindAllChooserDto, error) {
	choosers, err := chooserUseCase.ChooserRepository.FindAll()

	output := OutputFindAllChooserDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	choosersOutput := []OutputFindChooserDto{}

	for _, chooser := range choosers {
		choosersOutput = append(choosersOutput, OutputFindChooserDto{chooser.ID, chooser.UserName, chooser.Picture})
	}

	output = OutputFindAllChooserDto{
		Choosers: choosersOutput,
	}

	return output, nil
}

func (chooserUseCase *ChooserUseCase) Find(input InputFindChooserDto) (OutputFindChooserDto, error) {
	choosers, err := chooserUseCase.ChooserRepository.FindAll()

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

func (chooserUseCase *ChooserUseCase) Delete(input InputDeleteChooserDto) (OutputDeleteChooserDto, error) {
	output := OutputDeleteChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	chooser.IsDeleted = true
	chooser.DeletedAt = time.Now().Local().String()

	err = chooserUseCase.ChooserRepository.Delete(&chooser)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = true

	return output, nil
}

func (chooserUseCase *ChooserUseCase) Update(input InputUpdateChooserDto) (OutputUpdateChooserDto, error) {
	output := OutputUpdateChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fmt.Println(chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Picture)
	fmt.Println(input.FirstName, input.LastName, input.UserName, input.Picture)

	chooser.FirstName = input.FirstName
	chooser.LastName = input.LastName
	chooser.UserName = input.UserName
	chooser.Picture = input.Picture

	isValidChooser, err := chooser.Validate()
	if !isValidChooser {
		return output, errors.New(err.Error())
	}

	fmt.Println(chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Picture, chooser.UpdatedAt)

	chooser.UpdatedAt = time.Now().Local().String()

	fmt.Println(input.FirstName, input.LastName, input.UserName, input.Picture, chooser.UpdatedAt)

	err = chooserUseCase.ChooserRepository.Update(&chooser)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.ID = chooser.ID
	output.FirstName = chooser.FirstName
	output.LastName = chooser.LastName
	output.UserName = chooser.UserName
	output.Picture = chooser.Picture

	return output, nil
}
