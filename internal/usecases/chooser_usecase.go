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

	isValidChooser, err := chooser.Validate()
	if !isValidChooser {
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

func (chooserUseCase *ChooserUseCase) IsDeleted(input InputIsDeletedChooserDto) (OutputIsDeletedChooserDto, error) {
	output := OutputIsDeletedChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output = OutputIsDeletedChooserDto{
		IsDeleted: false,
	}

	if chooser.IsDeleted {
		output = OutputIsDeletedChooserDto{
			IsDeleted: true,
		}
	}

	return output, nil
}

func (chooserUseCase *ChooserUseCase) CreateChooserAndMovieList(input InputCreateChooserAndMovieListDto) (OutputCreateChooserAndMovieListDto, error) {
	output := OutputCreateChooserAndMovieListDto{}

	chooser, err := entity.NewChooser(input.Chooser.FirstName, input.Chooser.LastName, input.Chooser.UserName, input.Chooser.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	movieList, err := entity.NewMovieList(input.MovieList.Title, input.MovieList.Description, input.MovieList.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	outputChooser := OutputChooser{
		ChooserId: chooser.ID,
		UserName:  chooser.UserName,
		Picture:   chooser.Picture,
	}

	outputMovieList := OutputMovieList{
		ID:          movieList.ID,
		Title:       movieList.Title,
		Description: movieList.Description,
		Picture:     movieList.Picture,
	}

	outputMovieList.Choosers = append(outputMovieList.Choosers, outputChooser)

	output.MovieList = outputMovieList

	return output, nil
}

func (chooserUseCase *ChooserUseCase) ChooserCreateMovieList(input InputChooserCreateMovieListDto) (OutputChooserCreateMovieListDto, error) {
	output := OutputChooserCreateMovieListDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	movieList, err := entity.NewMovieList(input.MovieList.Title, input.MovieList.Description, input.MovieList.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := chooserUseCase.MovieListRepository.Create(movieList); err != nil {
		return output, errors.New(err.Error())
	}

	timeNow := time.Now().Local().String()

	err = chooserUseCase.MovieListRepository.AddChooserToMovieList(movieList, &chooser, timeNow, timeNow, timeNow)
	if err != nil {
		return output, errors.New(err.Error())
	}

	outputChooser := OutputChooser{
		ChooserId: chooser.ID,
		UserName:  chooser.UserName,
		Picture:   chooser.Picture,
	}

	output.ID = movieList.ID
	output.Title = movieList.Title
	output.Description = movieList.Description
	output.Picture = movieList.Picture
	output.Chooser = outputChooser

	return output, nil
}

func (chooserUseCase *ChooserUseCase) FindAllChooserMovieLists(input InputFindAllChooserMovieListsDto) (OutputFindAllChooserMovieListsDto, error) {
	output := OutputFindAllChooserMovieListsDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	movieLists, err := chooserUseCase.ChooserRepository.FindAllChooserMovieLists(chooser.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	outputChooser := OutputChooser{
		ChooserId: chooser.ID,
		UserName:  chooser.UserName,
		Picture:   chooser.Picture,
	}

	outputMovieLists := []OutputMovieList{}
	outputChoosers := []OutputChooser{}

	for _, movieList := range movieLists {
		outputMovieLists = append(outputMovieLists, OutputMovieList{
			ID:          movieList.ID,
			Title:       movieList.Title,
			Description: movieList.Description,
			Picture:     movieList.Picture,
		})

		for _, chooserMovieList := range movieList.Choosers {
			chooserInList := OutputChooser{
				ChooserId: chooserMovieList.ID,
				UserName:  chooserMovieList.UserName,
				Picture:   chooserMovieList.Picture,
			}
			outputChoosers = append(outputChoosers, chooserInList)
		}
		outputMovieLists = append(outputMovieLists, OutputMovieList{
			Choosers: outputChoosers,
		})
	}

	output.Chooser = outputChooser
	output.MovieLists = outputMovieLists

	return output, nil
}

func (chooserUseCase *ChooserUseCase) ChooserAddMovieToMovieList(input InputChooserAddMovieToMovieListDto) (OutputChooserAddMovieToMovieListDto, error) {
	output := OutputChooserAddMovieToMovieListDto{}

	movieList, err := chooserUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	movies, err := chooserUseCase.MovieRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, movieId := range input.MoviesIds {
		for _, movie := range movies {
			if movieId.MovieId == movie.ID {
				movieList.AddMovie(&movie)
			}
		}
	}

	var outputMovies []MovieDto
	var outputChoosers []ChooserDto

	for _, movie := range movieList.Movies {
		outputMovies = append(outputMovies, MovieDto{
			ID:              movie.ID,
			Title:           movie.Title,
			Synopsis:        movie.Synopsis,
			ImdbRating:      movie.ImdbRating,
			YouChooseRating: movie.YouChooseRating,
			Votes:           movie.Votes,
			Poster:          movie.Poster,
			IsDeleted:       movie.IsDeleted,
			CreatedAt:       movie.CreatedAt,
			UpdatedAt:       movie.UpdatedAt,
			DeletedAt:       movie.DeletedAt,
		})
	}

	for _, chooser := range movieList.Choosers {
		outputChoosers = append(outputChoosers, ChooserDto{
			ID:        chooser.ID,
			FirstName: chooser.FirstName,
			LastName:  chooser.LastName,
			UserName:  chooser.UserName,
			Picture:   chooser.Picture,
		})
	}

	err = chooserUseCase.ChooserRepository.ChooserAddMovieToMovieList(movieList, movieList.Movies)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.MovieList = MovieListDto{
		Movies:   outputMovies,
		Choosers: outputChoosers,
	}

	return output, nil
}
