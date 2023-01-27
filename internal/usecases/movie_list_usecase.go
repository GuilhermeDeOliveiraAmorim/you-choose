package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieListUseCase struct {
	MovieListRepository entity.MovieListRepositoryInterface
	ChooserRepository   entity.ChooserRepositoryInterface
}

func NewMovieListUseCase(movieListRepository entity.MovieListRepositoryInterface, chooserRepository entity.ChooserRepositoryInterface) *MovieListUseCase {
	return &MovieListUseCase{
		MovieListRepository: movieListRepository,
		ChooserRepository:   chooserRepository,
	}
}

func (movieListUseCase *MovieListUseCase) Create(input InputCreateMovieListDto) (OutputCreateMovieListDto, error) {
	movieList, err := entity.NewMovieList(input.Title, input.Description, input.Picture)

	output := OutputCreateMovieListDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := movieListUseCase.MovieListRepository.Create(movieList); err != nil {
		return output, errors.New(err.Error())
	}

	output.ID = movieList.ID
	output.Title = movieList.Title
	output.Description = movieList.Description
	output.Picture = movieList.Picture

	return output, nil
}

func (movieListUseCase *MovieListUseCase) FindAll() (OutputFindAllMovieListDto, error) {
	movieLists, err := movieListUseCase.MovieListRepository.FindAll()

	output := OutputFindAllMovieListDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	movieListsOutput := []OutputFindMovieListDto{}

	for _, movieList := range movieLists {
		movieListsOutput = append(movieListsOutput, OutputFindMovieListDto{movieList.ID, movieList.Title, movieList.Description, movieList.Picture})
	}

	output = OutputFindAllMovieListDto{
		MovieLists: movieListsOutput,
	}

	return output, nil
}

func (movieListUseCase *MovieListUseCase) Find(input InputFindMovieListDto) (OutputFindMovieListDto, error) {
	movieLists, err := movieListUseCase.MovieListRepository.FindAll()

	output := OutputFindMovieListDto{}

	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, movieList := range movieLists {
		if input.ID == movieList.ID {
			output.ID = movieList.ID
			output.Title = movieList.Title
			output.Description = movieList.Description
			output.Picture = movieList.Picture
			return output, nil
		}
	}

	return output, errors.New(err.Error())
}

func (movieListUseCase *MovieListUseCase) AddChooserToMovieList(input InputAddChooserToMovieListDto) (OutputAddChooserToMovieListDto, error) {
	output := OutputAddChooserToMovieListDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	chooser, err := movieListUseCase.ChooserRepository.Find(input.ChooserId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	timeNow := time.Now().Local().String()

	err = movieListUseCase.MovieListRepository.AddChooserToMovieList(&movieList, &chooser, timeNow, timeNow, timeNow)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Chooser = OutputFindChooserDto{
		ID:       chooser.ID,
		UserName: chooser.UserName,
		Picture:  chooser.Picture,
	}

	output.MovieList = OutputFindMovieListDto{
		ID:          movieList.ID,
		Title:       movieList.Title,
		Description: movieList.Description,
		Picture:     movieList.Picture,
	}

	return output, nil
}
