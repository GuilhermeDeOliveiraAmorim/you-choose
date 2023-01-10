package domain

import (
	"context"

	movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

type MovieListRepositoryInterface interface {
	Create(ctx context.Context, a *movieList.MovieList) (*movieList.MovieList, error)
	Update(ctx context.Context, a *movieList.MovieList) (*movieList.MovieList, error)
	FindById(ctx context.Context, id string) (*movieList.MovieList, error)
	DeleteById(ctx context.Context, id string) (*movieList.MovieList, error)
	FindAll(ctx context.Context) ([]*movieList.MovieList, error)
}
