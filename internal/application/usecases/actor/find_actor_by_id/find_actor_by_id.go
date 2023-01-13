package findactorbyid

import (
	actorRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type FindActorByIdUseCase struct {
	ActorRepository actorRepository.ActorRepositoryInterface
}

func NewFindActorByIdUseCase(ActorRepository actorRepository.ActorRepositoryInterface) *FindActorByIdUseCase {
	return &FindActorByIdUseCase{
		ActorRepository: ActorRepository,
	}
}

func (actorUseCase *FindActorByIdUseCase) Execute(input InputFindActorByIdDto) (OutputFindActorByIdDto, error) {
	actors, err := actorUseCase.ActorRepository.FindAll()

	output := OutputFindActorByIdDto{}

	if err != nil {
		return output, err
	}

	actorIdToFind := input.ID

	for position, registeredActor := range actors {
		if registeredActor.ID == actorIdToFind {
			return OutputFindActorByIdDto{
				Actor: *actors[position],
			}, nil
		}
	}

	return output, err
}
