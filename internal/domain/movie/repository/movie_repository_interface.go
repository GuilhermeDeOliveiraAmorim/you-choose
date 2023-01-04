package domain

import (
	"context"

	movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"
)

type MovieRepositoryInterface interface {
	Add(ctx context.Context, g *movie.Movie) (*movie.Movie, error)
	Find(ctx context.Context, id string) (*movie.Movie, error)
	FindAll(ctx context.Context) ([]*movie.Movie, error)
}
