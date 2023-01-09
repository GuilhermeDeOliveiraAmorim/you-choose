package findmoviebyid

import movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"

type InputFindMovieByIdDto struct {
	MovieID string `json:"movie_id"`
}

type OutputFindMovieByIdDto struct {
	Movie movie.Movie `json:"movie"`
}
