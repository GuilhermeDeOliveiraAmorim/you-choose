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

func (chooserUseCase *ChooserUseCase) AddMoviesToMovieList(input InputAddMoviesToMovieListDto) (OutputAddMoviesToMovieListDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddMoviesToMovieListDto{}

	movieList, err := chooserUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	moviesIds, err := chooserUseCase.MovieListRepository.FindMovieListMovies(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var moviesMovieList []entity.Movie

	for _, movieId := range moviesIds {
		movie, err := chooserUseCase.MovieRepository.Find(movieId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		moviesMovieList = append(moviesMovieList, movie)
	}

	var moviesAdded []entity.Movie

	for _, movieId := range input.MoviesIds {
		movie, err := chooserUseCase.MovieRepository.Find(movieId.MovieId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		moviesAdded = append(moviesAdded, movie)
	}

	for _, movieMovieList := range moviesMovieList {
		for position, movieAdded := range moviesAdded {
			if movieMovieList.ID == movieAdded.ID {
				moviesAdded = append(moviesAdded[:position], moviesAdded[position+1:]...)
			}
		}
	}

	err = chooserUseCase.ChooserRepository.AddMoviesToMovieList(movieList, moviesAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputMovies []MovieDto

	for _, movieId := range moviesIds {
		movie, err := chooserUseCase.MovieRepository.Find(movieId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputMovies = append(outputMovies, MovieDto{
			ID:              movie.ID,
			Title:           movie.Title,
			Synopsis:        movie.Synopsis,
			ImdbRating:      movie.ImdbRating,
			Votes:           movie.Votes,
			YouChooseRating: movie.YouChooseRating,
			Poster:          movie.Poster,
			IsDeleted:       movie.IsDeleted,
			CreatedAt:       movie.CreatedAt,
			UpdatedAt:       movie.UpdatedAt,
			DeletedAt:       movie.DeletedAt,
		})
	}

	output.MovieList.ID = movieList.ID
	output.MovieList.Title = movieList.Title
	output.MovieList.Description = movieList.Description
	output.MovieList.Picture = movieList.Picture
	output.MovieList.IsDeleted = movieList.IsDeleted
	output.MovieList.CreatedAt = movieList.CreatedAt
	output.MovieList.UpdatedAt = dateNow
	output.MovieList.DeletedAt = movieList.DeletedAt
	output.MovieList.Movies = outputMovies

	return output, nil
}
