package addchoosertolist

import (
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

type InputAddChooserToMovieListDto struct {
	MovieList *movieList.MovieList `json:"movieList"`
	Chooser   *chooser.Chooser     `json:"chooser"`
}

type OutputAddChooserToMovieListDto struct {
	IDMovieList string `json:"id_movie_list"`
	IDChooser   string `json:"id_chooser"`
}
