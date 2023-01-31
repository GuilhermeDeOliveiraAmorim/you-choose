package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ActorUseCase struct {
	ActorRepository entity.ActorRepositoryInterface
}

func NewActorUseCase(ActorRepository entity.ActorRepositoryInterface) *ActorUseCase {
	return &ActorUseCase{
		ActorRepository: ActorRepository,
	}
}

func (a *ActorUseCase) Create(input InputCreateActorDto) (OutputCreateActorDto, error) {
	actor, err := entity.NewActor(input.Name, input.Picture)

	output := OutputCreateActorDto{}

	if err != nil {
		return output, err
	}

	a.ActorRepository.Create(actor)

	output.ID = actor.ID
	output.Name = actor.Name
	output.Picture = actor.Picture

	return output, nil
}

func (a *ActorUseCase) Delete(input InputDeleteActorDto) (OutputDeleteActorDto, error) {
	actor, err := a.ActorRepository.Find(input.ID)

	output := OutputDeleteActorDto{}

	if err != nil {
		return output, errors.New("actor not found")
	}

	actor.IsDeleted = true
	actor.DeletedAt = time.Now().Local().String()

	output.IsDeleted = actor.IsDeleted

	return output, err
}

func (a *ActorUseCase) Find(input InputFindActorDto) (OutputFindActorDto, error) {
	actors, err := a.ActorRepository.FindAll()

	output := OutputFindActorDto{}

	if err != nil {
		return output, err
	}

	actorIdToFind := input.ID

	for position, registeredActor := range actors {
		if registeredActor.ID == actorIdToFind {
			return OutputFindActorDto{
				Actor: *actors[position],
			}, nil
		}
	}

	return output, err
}
