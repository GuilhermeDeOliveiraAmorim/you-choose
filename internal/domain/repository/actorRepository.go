package repository

import (
	"context"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/entity"
)

type ActorRepositoryInterface interface {
	Create(ctx context.Context, actor *entity.Actor) error
	Update(ctx context.Context, actor *entity.Actor) error
	FindByID(ctx context.Context, id string) (*entity.Actor, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.Actor, error)
	FindAll(ctx context.Context) ([]*entity.Actor, error)
	FindAllByIDs(ctx context.Context, ids []string) ([]entity.Actor, error)
}
