package domain

import (
	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
)

type ActorRepositoryInterface interface {
	Create(a *actor.Actor) (*actor.Actor, error)
	Update(a *actor.Actor) (*actor.Actor, error)
	FindById(id string) (*actor.Actor, error)
	DeleteById(id string) (*actor.Actor, error)
	FindAll() ([]*actor.Actor, error)
}
