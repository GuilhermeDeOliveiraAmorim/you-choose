package findactorbyid

import (
	"context"
	"errors"

	actorRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type ActorRepository struct {
	actorRepositoryInterface actorRepository.ActorRepositoryInterface
}

func FindActorByIdUseCase(input *InputFindActorByIdDto, ctx context.Context, repository *ActorRepository) (*OutputFindActorByIdDto, error) {
	if input.ID == "" {
		return nil, errors.New("id cannot be empty")
	}

	if input.Actors == nil {
		return nil, errors.New("actor list not loaded")
	}

	actorId := input.ID
	actorsArr := input.Actors

	for position, value := range actorsArr {
		if value.ID == actorId {
			return actorsArr[position], nil
		}
	}

	return output, err
}
