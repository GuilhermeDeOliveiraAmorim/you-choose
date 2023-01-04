package createactor

import (
	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
)

func CreateActorUseCase(input *InputCreateActorDto) *OutputCreateActorDto {
	if input.Name == "" {
		return nil
	}

	if input.Picture == "" {
		return nil
	}

	actorOutput, _ := actor.NewActor(input.Name, input.Picture)

	output := OutputCreateActorDto{
		actorOutput.ID,
		actorOutput.Name,
		actorOutput.Picture,
	}

	return &output
}
