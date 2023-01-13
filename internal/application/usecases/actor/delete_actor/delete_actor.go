package deleteactor

import (
	"errors"
	"time"

	actorRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type DeleteActorUseCase struct {
	ActorRepository actorRepository.ActorRepositoryInterface
}

func NewDeleteActorUseCase(ActorRepository actorRepository.ActorRepositoryInterface) *DeleteActorUseCase {
	return &DeleteActorUseCase{
		ActorRepository: ActorRepository,
	}
}

func (actorUseCase *DeleteActorUseCase) Execute(input InputDeleteActorDto) (OutputDeleteActorDto, error) {
	actor, err := actorUseCase.ActorRepository.FindById(input.ID)

	output := OutputDeleteActorDto{}

	if err != nil {
		return output, errors.New("actor not found")
	}

	actor.IsDeleted = true
	actor.DeletedAt = time.Now()

	output.IsDeleted = actor.IsDeleted

	return output, err
}
