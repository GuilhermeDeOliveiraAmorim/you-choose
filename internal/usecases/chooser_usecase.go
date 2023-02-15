package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ChooserUseCase struct {
	ChooserRepository   entity.ChooserRepositoryInterface
	MovieListRepository entity.MovieListRepositoryInterface
	MovieRepository     entity.MovieRepositoryInterface
}

func NewChooserUseCase(chooserRepository entity.ChooserRepositoryInterface, movieListRepository entity.MovieListRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *ChooserUseCase {
	return &ChooserUseCase{
		ChooserRepository:   chooserRepository,
		MovieListRepository: movieListRepository,
		MovieRepository:     movieRepository,
	}
}

func (chooserUseCase *ChooserUseCase) Create(input InputCreateChooserDto) (OutputCreateChooserDto, error) {
	output := OutputCreateChooserDto{}

	chooser, err := entity.NewChooser(input.FirstName, input.LastName, input.UserName, input.Picture)
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

func (chooserUseCase *ChooserUseCase) Find(input InputFindChooserDto) (OutputFindChooserDto, error) {
	output := OutputFindChooserDto{}

	choosers, err := chooserUseCase.ChooserRepository.FindAll()
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

func (chooserUseCase *ChooserUseCase) Update(input InputUpdateChooserDto) (OutputUpdateChooserDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	chooser.FirstName = input.FirstName
	chooser.LastName = input.LastName
	chooser.UserName = input.UserName
	chooser.Picture = input.Picture

	isValid, err := chooser.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	chooser.UpdatedAt = timeNow

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

func (chooserUseCase *ChooserUseCase) Delete(input InputDeleteChooserDto) (OutputDeleteChooserDto, error) {
	output := OutputDeleteChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if chooser.IsDeleted {
		return output, errors.New("chooser previously deleted")
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

func (chooserUseCase *ChooserUseCase) IsDeleted(input InputIsDeletedChooserDto) (OutputIsDeletedChooserDto, error) {
	output := OutputIsDeletedChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if chooser.IsDeleted {
		output.IsDeleted = true
	}

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
