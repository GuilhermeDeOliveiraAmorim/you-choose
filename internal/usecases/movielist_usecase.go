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
	output := OutputCreateMovieListDto{}

	movieList, err := entity.NewMovieList(input.Title, input.Description, input.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := movieListUseCase.MovieListRepository.Create(movieList); err != nil {
		return output, errors.New(err.Error())
	}

	output.MovieList.ID = movieList.ID
	output.MovieList.Title = movieList.Title
	output.MovieList.Description = movieList.Description
	output.MovieList.Picture = movieList.Picture
	output.MovieList.IsDeleted = movieList.IsDeleted
	output.MovieList.CreatedAt = movieList.CreatedAt
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt

	return output, nil
}

func (movieListUseCase *MovieListUseCase) Find(input InputFindMovieListDto) (OutputFindMovieListDto, error) {
	output := OutputFindMovieListDto{}

	movieLists, err := movieListUseCase.MovieListRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, movieList := range movieLists {
		if input.MovieListId == movieList.ID {
			output.MovieList.ID = movieList.ID
			output.MovieList.Title = movieList.Title
			output.MovieList.Description = movieList.Description
			output.MovieList.Picture = movieList.Picture
			output.MovieList.IsDeleted = movieList.IsDeleted
			output.MovieList.CreatedAt = movieList.CreatedAt
			output.MovieList.UpdatedAt = movieList.UpdatedAt
			output.MovieList.DeletedAt = movieList.DeletedAt
			return output, nil
		}
	}

	return output, nil
}

func (movieListUseCase *MovieListUseCase) Delete(input InputDeleteMovieListDto) (OutputDeleteMovieListDto, error) {
	output := OutputDeleteMovieListDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if movieList.IsDeleted {
		return output, errors.New("movieList previously deleted")
	}

	movieList.IsDeleted = true
	movieList.DeletedAt = time.Now().Local().String()

	output.IsDeleted = movieList.IsDeleted

	return output, nil
}

func (movieListUseCase *MovieListUseCase) Update(input InputUpdateMovieListDto) (OutputUpdateMovieListDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateMovieListDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	movieList.Title = input.Title
	movieList.Description = input.Description
	movieList.Picture = input.Picture

	isValid, err := movieList.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	movieList.UpdatedAt = timeNow

	err = movieListUseCase.MovieListRepository.Update(&movieList)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.MovieList.ID = movieList.ID
	output.MovieList.Title = movieList.Title
	output.MovieList.Picture = movieList.Picture
	output.MovieList.IsDeleted = movieList.IsDeleted
	output.MovieList.CreatedAt = movieList.CreatedAt
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt

	return output, nil
}

func (movieListUseCase *MovieListUseCase) IsDeleted(input InputIsDeletedMovieListDto) (OutputIsDeletedMovieListDto, error) {
	output := OutputIsDeletedMovieListDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if movieList.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (movieListUseCase *MovieListUseCase) FindAll() (OutputFindAllMovieListDto, error) {
	output := OutputFindAllMovieListDto{}

	movieLists, err := movieListUseCase.MovieListRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, movieList := range movieLists {
		output.MovieLists = append(output.MovieLists, MovieListDto{
			ID:          movieList.ID,
			Title:       movieList.Title,
			Description: movieList.Description,
			Picture:     movieList.Picture,
			IsDeleted:   movieList.IsDeleted,
			CreatedAt:   movieList.CreatedAt,
			UpdatedAt:   movieList.UpdatedAt,
			DeletedAt:   movieList.DeletedAt,
		})
	}

	return output, nil
}
