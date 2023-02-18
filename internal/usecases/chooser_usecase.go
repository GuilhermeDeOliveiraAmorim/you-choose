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
	TagRepository       entity.TagRepositoryInterface
}

func NewChooserUseCase(chooserRepository entity.ChooserRepositoryInterface, movieListRepository entity.MovieListRepositoryInterface, movieRepository entity.MovieRepositoryInterface, tagRepository entity.TagRepositoryInterface) *ChooserUseCase {
	return &ChooserUseCase{
		ChooserRepository:   chooserRepository,
		MovieListRepository: movieListRepository,
		MovieRepository:     movieRepository,
		TagRepository:       tagRepository,
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

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if chooser.ID == "" {
		return output, errors.New("chooser not found")
	}

	output.ID = chooser.ID
	output.UserName = chooser.UserName
	output.Picture = chooser.Picture

	return output, errors.New(err.Error())
}

func (chooserUseCase *ChooserUseCase) Update(input InputUpdateChooserDto) (OutputUpdateChooserDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateChooserDto{}

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
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

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
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

	chooser, err := chooserUseCase.ChooserRepository.Find(input.ChooserId)
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
	output := OutputFindAllChooserDto{}

	choosers, err := chooserUseCase.ChooserRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	if len(choosers) == 0 {
		return output, errors.New("no chooser found")
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

func (chooserUseCase *ChooserUseCase) CreateMovieList(input InputCreateMovieListDto) (OutputCreateMovieListDto, error) {
	output := OutputCreateMovieListDto{}

	movieList, err := entity.NewMovieList(input.Title, input.Description, input.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := chooserUseCase.MovieListRepository.Create(movieList); err != nil {
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

func (chooserUseCase *ChooserUseCase) AddChoosersToMovieList(input InputAddChoosersToMovieListDto) (OutputAddChoosersToMovieListDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddChoosersToMovieListDto{}

	movieList, err := chooserUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	choosersIds, err := chooserUseCase.MovieListRepository.FindMovieListChoosers(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var choosersInMovieList []entity.Chooser

	for _, chooserId := range choosersIds {
		chooser, err := chooserUseCase.ChooserRepository.Find(chooserId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		choosersInMovieList = append(choosersInMovieList, chooser)
	}

	var choosersAdded []entity.Chooser

	for _, chooserId := range input.ChoosersIds {
		chooser, err := chooserUseCase.ChooserRepository.Find(chooserId.ChooserId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		choosersAdded = append(choosersAdded, chooser)
	}

	for _, chooserInMovieList := range choosersInMovieList {
		for position, chooserAdded := range choosersAdded {
			if chooserInMovieList.ID == chooserAdded.ID {
				choosersAdded = append(choosersAdded[:position], choosersAdded[position+1:]...)
			}
		}
	}

	err = chooserUseCase.ChooserRepository.AddChoosersToMovieList(movieList, choosersAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputChoosers []ChooserDto

	for _, chooserId := range choosersIds {
		chooser, err := chooserUseCase.ChooserRepository.Find(chooserId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputChoosers = append(outputChoosers, ChooserDto{
			ID:        chooser.ID,
			FirstName: chooser.FirstName,
			LastName:  chooser.LastName,
			UserName:  chooser.UserName,
			Picture:   chooser.Picture,
			IsDeleted: chooser.IsDeleted,
			CreatedAt: chooser.CreatedAt,
			UpdatedAt: chooser.UpdatedAt,
			DeletedAt: chooser.DeletedAt,
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
	output.MovieList.Choosers = outputChoosers

	return output, nil
}

func (chooserUseCase *ChooserUseCase) AddTagsToMovieList(input InputAddTagsToMovieListDto) (OutputAddTagsToMovieListDto, error) {
	dateNow := time.Now().Local().String()

	output := OutputAddTagsToMovieListDto{}

	movieList, err := chooserUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	tagsIds, err := chooserUseCase.MovieListRepository.FindMovieListTags(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var tagsInMovieList []entity.Tag

	for _, chooserId := range tagsIds {
		chooser, err := chooserUseCase.TagRepository.Find(chooserId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		tagsInMovieList = append(tagsInMovieList, chooser)
	}

	var tagsAdded []entity.Tag

	for _, chooserId := range input.TagsIds {
		chooser, err := chooserUseCase.TagRepository.Find(chooserId.TagId)
		if err != nil {
			return output, errors.New(err.Error())
		}
		tagsAdded = append(tagsAdded, chooser)
	}

	for _, chooserInMovieList := range tagsInMovieList {
		for position, chooserAdded := range tagsAdded {
			if chooserInMovieList.ID == chooserAdded.ID {
				tagsAdded = append(tagsAdded[:position], tagsAdded[position+1:]...)
			}
		}
	}

	err = chooserUseCase.ChooserRepository.AddTagsToMovieList(movieList, tagsAdded)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputTags []TagDto

	for _, tagId := range tagsIds {
		tag, err := chooserUseCase.TagRepository.Find(tagId)
		if err != nil {
			return output, errors.New(err.Error())
		}

		outputTags = append(outputTags, TagDto{
			ID:        tag.ID,
			Name:      tag.Name,
			Picture:   tag.Picture,
			IsDeleted: tag.IsDeleted,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
			DeletedAt: tag.DeletedAt,
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
	output.MovieList.Tags = outputTags

	return output, nil
}
