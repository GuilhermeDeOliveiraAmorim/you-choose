package updateyouchooserating

import movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"

type InputUpdateYouChooseRatingMovieDto struct {
	MovieID string `json:"movie_id"`
}

type OutputUpdateYouChooseRatingMovieDto struct {
	Movie movie.Movie `json:"movie"`
}
