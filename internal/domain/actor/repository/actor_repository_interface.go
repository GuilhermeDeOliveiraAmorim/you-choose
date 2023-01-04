package domain

import (
	"context"

	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
)

type ActorRepositoryInterface interface {
	Add(ctx context.Context, a *actor.Actor) (*actor.Actor, error)
	Find(ctx context.Context, id string) (*actor.Actor, error)
	FindAll(ctx context.Context) ([]*actor.Actor, error)
}
