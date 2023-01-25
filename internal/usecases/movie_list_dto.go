package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type InputAddChooserToMovieListDto struct {
	MovieList entity.MovieList `json:"movieList"`
	Chooser   entity.Chooser   `json:"chooser"`
}

type OutputAddChooserToMovieListDto struct {
	IDMovieList string `json:"id_movie_list"`
	IDChooser   string `json:"id_chooser"`
}

type InputCreateMovieListDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputCreateMovieListDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type InputFindMovieListDto struct {
	ID string `json:"id"`
}

type OutputFindMovieListDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputFindAllMovieListDto struct {
	MovieLists []OutputFindMovieListDto `json:"movie_lists"`
}
