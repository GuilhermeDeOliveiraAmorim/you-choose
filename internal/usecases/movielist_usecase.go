package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieListUseCase struct {
	MovieListRepository entity.MovieListRepositoryInterface
	ChooserRepository   entity.ChooserRepositoryInterface
	MovieRepository     entity.MovieRepositoryInterface
	TagRepository       entity.TagRepositoryInterface
	FileRepository      entity.FileRepositoryInterface
}

func NewMovieListUseCase(
	movieListRepository entity.MovieListRepositoryInterface,
	chooserRepository entity.ChooserRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	tagRepository entity.TagRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *MovieListUseCase {
	return &MovieListUseCase{
		MovieListRepository: movieListRepository,
		ChooserRepository:   chooserRepository,
		MovieRepository:     movieRepository,
		TagRepository:       tagRepository,
		FileRepository:      fileRepository,
	}
}

func (movieListUseCase *MovieListUseCase) Create(input InputCreateMovieListDto) (OutputCreateMovieListDto, error) {
	output := OutputCreateMovieListDto{}

	movieList, err := entity.NewMovieList(input.Title, input.Description)
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

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if movieList.ID == "" {
		return output, errors.New("movie list not found")
	}

	inputFindPicture := InputFindMovieListPictureToBase64Dto{
		MovieListId: movieList.ID,
	}

	picture, err := movieListUseCase.FindMovieListPictureToBase64(inputFindPicture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	file, err := movieListUseCase.FileRepository.Find(movieList.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fileDto := FileDto{
		ID:           file.ID,
		EntityId:     file.EntityId,
		Name:         file.Name,
		Size:         file.Size,
		Extension:    file.Extension,
		AverageColor: file.AverageColor,
		IsDeleted:    file.IsDeleted,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
		DeletedAt:    file.DeletedAt,
	}

	output.MovieList.ID = movieList.ID
	output.MovieList.Title = movieList.Title
	output.MovieList.Description = movieList.Description
	output.MovieList.Picture = picture.MovieList.Picture
	output.MovieList.IsDeleted = movieList.IsDeleted
	output.MovieList.CreatedAt = movieList.CreatedAt
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt
	output.MovieList.File = fileDto

	choosersIds, err := movieListUseCase.MovieListRepository.FindMovieListChoosers(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputChoosers []ChooserDto

	for _, chooserId := range choosersIds {
		chooser, err := movieListUseCase.ChooserRepository.Find(chooserId)
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

	moviesIds, err := movieListUseCase.MovieListRepository.FindMovieListMovies(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputMovies []MovieDto

	for _, movieId := range moviesIds {
		movie, err := movieListUseCase.MovieRepository.Find(movieId)
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

	tagsIds, err := movieListUseCase.MovieListRepository.FindMovieListTags(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputTags []TagDto

	for _, tagId := range tagsIds {
		tag, err := movieListUseCase.TagRepository.Find(tagId)
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

	output.MovieList.Choosers = outputChoosers
	output.MovieList.Movies = outputMovies
	output.MovieList.Tags = outputTags

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

	err = movieListUseCase.MovieListRepository.Update(&movieList)
	if err != nil {
		return output, errors.New(err.Error())
	}

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

func (movieListUseCase *MovieListUseCase) FindMovieListMovies(input InputFindMovieListMoviesDto) (OutputFindMovieListMoviesDto, error) {
	output := OutputFindMovieListMoviesDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	moviesIds, err := movieListUseCase.MovieListRepository.FindMovieListMovies(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputMovies []MovieDto

	for _, movieId := range moviesIds {
		movie, err := movieListUseCase.MovieRepository.Find(movieId)
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
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt
	output.MovieList.Movies = outputMovies

	return output, nil
}

func (movieListUseCase *MovieListUseCase) FindMovieListChoosers(input InputFindMovieListChoosersDto) (OutputFindMovieListChoosersDto, error) {
	output := OutputFindMovieListChoosersDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	choosersIds, err := movieListUseCase.MovieListRepository.FindMovieListChoosers(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputChoosers []ChooserDto

	for _, chooserId := range choosersIds {
		chooser, err := movieListUseCase.ChooserRepository.Find(chooserId)
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
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt
	output.MovieList.Choosers = outputChoosers

	return output, nil
}

func (movieListUseCase *MovieListUseCase) FindMovieListTags(input InputFindMovieListTagsDto) (OutputFindMovieListTagsDto, error) {
	output := OutputFindMovieListTagsDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	tagsIds, err := movieListUseCase.MovieListRepository.FindMovieListTags(movieList.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	var outputTags []TagDto

	for _, tagId := range tagsIds {
		tag, err := movieListUseCase.TagRepository.Find(tagId)
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
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt
	output.MovieList.Tags = outputTags

	return output, nil
}

func (movieListUseCase *MovieListUseCase) AddPictureToMovieList(input InputAddPictureToMovieListDto) (OutputAddPictureToMovieListDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToMovieListDto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	extension := strings.Replace(filepath.Ext(input.File.Handler.Filename), ".", "", -1)
	if strings.ToLower(extension) != "jpeg" && strings.ToLower(extension) != "jpg" {
		return output, errors.New("format not allowed")
	}

	_, name, size, extension, err := MoveFile(input.File.File, input.File.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	colorAverage, err := PictureAverageColor(name, extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := entity.NewFile(name, movieList.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := movieListUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	movieList.Picture = picture.ID

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

func (movieListUseCase *MovieListUseCase) FindMovieListPictureToBase64(input InputFindMovieListPictureToBase64Dto) (OutputFindMovieListPictureToBase64Dto, error) {
	output := OutputFindMovieListPictureToBase64Dto{}

	movieList, err := movieListUseCase.MovieListRepository.Find(input.MovieListId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := movieListUseCase.FileRepository.Find(movieList.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64("/home/guilhermeamorim/Workspace/estudo/you-choose/cmd/upload/", picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.MovieList.ID = movieList.ID
	output.MovieList.Title = movieList.Title
	output.MovieList.Picture = pictureToBase64
	output.MovieList.IsDeleted = movieList.IsDeleted
	output.MovieList.CreatedAt = movieList.CreatedAt
	output.MovieList.UpdatedAt = movieList.UpdatedAt
	output.MovieList.DeletedAt = movieList.DeletedAt

	return output, nil
}
