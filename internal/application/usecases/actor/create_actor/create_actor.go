package createactor

import (
	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
	actorRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type CreateActorUseCase struct {
	ActorRepository actorRepository.ActorRepositoryInterface
}

func NewCreateActorUseCase(ActorRepository actorRepository.ActorRepositoryInterface) *CreateActorUseCase {
	return &CreateActorUseCase{
		ActorRepository: ActorRepository,
	}
}

func (c *CreateActorUseCase) Execute(input InputCreateActorDto) (OutputCreateActorDto, error) {
	actor, err := actor.NewActor(input.Name, input.Picture)

	output := OutputCreateActorDto{}

	if err != nil {
		return output, err
	}

	c.ActorRepository.Create(actor)

	output.ID = actor.ID
	output.Name = actor.Name
	output.Picture = actor.Picture

	return output, nil
}
