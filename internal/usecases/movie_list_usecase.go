package usecases

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"

type MovieListUseCase struct {
	MovieListRepository entity.MovieListRepositoryInterface
}

func NewMovieListUseCase(movieListRepository entity.MovieListRepositoryInterface) *MovieListUseCase {
	return &MovieListUseCase{
		MovieListRepository: movieListRepository,
	}
}

func (ml *MovieListUseCase) CreateMovieList(input InputCreateMovieListDto) (OutputCreateMovieListDto, error) {
	movieList, err := entity.NewMovieList(input.Title, input.Description, input.Picture)

	output := OutputCreateMovieListDto{}

	if err != nil {
		return output, err
	}

	if err := ml.MovieListRepository.Create(movieList); err != nil {
		return output, err
	}

	output.ID = movieList.ID
	output.Title = movieList.Title
	output.Description = movieList.Description
	output.Picture = movieList.Picture

	return output, nil
}

func AddChooserToMovieList(input InputAddChooserToMovieListDto) (OutputAddChooserToMovieListDto, error) {
	chooser := input.Chooser
	movieList := input.MovieList

	movieList.AddChooser(chooser)

	output := &OutputAddChooserToMovieListDto{
		IDMovieList: movieList.ID,
		IDChooser:   chooser.ID,
	}

	return output, nil
}
