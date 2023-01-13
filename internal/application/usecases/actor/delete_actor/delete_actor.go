package deleteactor

import (
	"context"
	"errors"
	"time"

	actorRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type ActorRepository struct {
	actorRepositoryInterface actorRepository.ActorRepositoryInterface
}

func DeleteActorUseCase(input *InputDeleteActorDto, ctx context.Context, repository *ActorRepository) (*OutputDeleteActorDto, error) {
	if input.ID == "" {
		return nil, errors.New("id cannot be empty")
	}

	actorFound, err := repository.actorRepositoryInterface.FindById(ctx, input.ID)
	if err != nil {
		return nil, errors.New("actor not found")
	}

	actorFound.IsDeleted = true
	actorFound.DeletedAt = time.Now()

	output := OutputDeleteActorDto{
		IsDeleted: actorFound.IsDeleted,
	}

	return &output, err
}
