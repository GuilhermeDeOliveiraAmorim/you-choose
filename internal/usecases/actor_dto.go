package usecases

import actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"

type ActorDto struct {
	ID        string `json:"actior_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
}

type InputCreateActorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateActorDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type InputDeleteActorDto struct {
	ID string `json:"id"`
}

type OutputDeleteActorDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindActorDto struct {
	ID     string        `json:"id"`
	Actors []actor.Actor `json:"actors"`
}

type OutputFindActorDto struct {
	Actor actor.Actor `json:"actor"`
}
