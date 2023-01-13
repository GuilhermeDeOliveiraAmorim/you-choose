package findactorbyid

import (
	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
)

type InputFindActorByIdDto struct {
	ID     string        `json:"id"`
	Actors []actor.Actor `json:"actors"`
}

type OutputFindActorByIdDto struct {
	Actor actor.Actor `json:"actor"`
}
