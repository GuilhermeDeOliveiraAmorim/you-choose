package usecases

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"

type InputUpdateYouChooseRatingMovieDto struct {
	MovieID string `json:"movie_id"`
}

type OutputUpdateYouChooseRatingMovieDto struct {
	Movie entity.Movie `json:"movie"`
}

type InputFindMovieDto struct {
	MovieID string `json:"movie_id"`
}

type OutputFindMovieDto struct {
	Movie entity.Movie `json:"movie"`
}
