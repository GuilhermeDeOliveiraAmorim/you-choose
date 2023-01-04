package domain

import (
	"context"

	director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
)

type DirectorRepositoryInterface interface {
	Add(ctx context.Context, d *director.Director) (*director.Director, error)
	Find(ctx context.Context, id string) (*director.Director, error)
	FindAll(ctx context.Context) ([]*director.Director, error)
}
