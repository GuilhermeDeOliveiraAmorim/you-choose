package usecases

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieListUseCase struct {
	MovieListRepository entity.MovieListRepositoryInterface
	ChooserRepository   entity.ChooserRepositoryInterface
}

func NewMovieListUseCase(movieListRepository entity.MovieListRepositoryInterface) *MovieListUseCase {
	return &MovieListUseCase{
		MovieListRepository: movieListRepository,
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
