package domain

import (
	"context"

	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
)

type ActorRepositoryInterface interface {
	Create(ctx context.Context, a *actor.Actor) (*actor.Actor, error)
	Update(ctx context.Context, a *actor.Actor) (*actor.Actor, error)
	FindById(ctx context.Context, id string) (*actor.Actor, error)
	DeleteById(ctx context.Context, id string) (*actor.Actor, error)
	FindAll(ctx context.Context) ([]*actor.Actor, error)
}
