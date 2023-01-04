package domain

import (
	"context"

	genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
)

type GenreRepositoryInterface interface {
	Add(ctx context.Context, g *genre.Genre) (*genre.Genre, error)
	Find(ctx context.Context, id string) (*genre.Genre, error)
	FindAll(ctx context.Context) ([]*genre.Genre, error)
}
