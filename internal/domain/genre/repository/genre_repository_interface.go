package domain

import (
	"context"

	genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
)

type GenreRepositoryInterface interface {
	Create(ctx context.Context, a *genre.Genre) (*genre.Genre, error)
	Update(ctx context.Context, a *genre.Genre) (*genre.Genre, error)
	FindById(ctx context.Context, id string) (*genre.Genre, error)
	DeleteById(ctx context.Context, id string) (*genre.Genre, error)
	FindAll(ctx context.Context) ([]*genre.Genre, error)
}
