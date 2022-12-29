package repository

import (
	"context"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
)

type DirectorRepositoryInterface interface {
	Create(ctx context.Context, director *entity.Director) error
	Update(ctx context.Context, director *entity.Director) error
	FindByID(ctx context.Context, id string) (*entity.Director, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.Director, error)
	FindAll(ctx context.Context) ([]*entity.Director, error)
	FindAllByIDs(ctx context.Context, ids []string) ([]entity.Director, error)
}
