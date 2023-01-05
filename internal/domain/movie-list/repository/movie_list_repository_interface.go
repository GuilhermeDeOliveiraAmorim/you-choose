package domain

import (
	"context"

	movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

type MovieListRepositoryInterface interface {
	Create(ctx context.Context, ml *movieList.MovieList) (*movieList.MovieList, error)
	FindById(ctx context.Context, id string) (*movieList.MovieList, error)
	FindAll(ctx context.Context) ([]*movieList.MovieList, error)
}
