package domain

import (
	"context"

	movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"
)

type MovieRepositoryInterface interface {
	Create(ctx context.Context, a *movie.Movie) (*movie.Movie, error)
	Update(ctx context.Context, a *movie.Movie) (*movie.Movie, error)
	FindById(ctx context.Context, id string) (*movie.Movie, error)
	DeleteById(ctx context.Context, id string) (*movie.Movie, error)
	FindAll(ctx context.Context) ([]*movie.Movie, error)
	UpdateYouChooseRating(ctx context.Context, id string) (*movie.Movie, error)
}
