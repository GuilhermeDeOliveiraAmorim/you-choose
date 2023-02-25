package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ActorUseCase struct {
	ActorRepository entity.ActorRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewActorUseCase(
	actorRepository entity.ActorRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *ActorUseCase {
	return &ActorUseCase{
		ActorRepository: actorRepository,
		MovieRepository: movieRepository,
		FileRepository:  fileRepository,
	}
}

func (actorUseCase *ActorUseCase) Create(input InputCreateActorDto) (OutputCreateActorDto, error) {
	output := OutputCreateActorDto{}

	actor, err := entity.NewActor(input.Name)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := actorUseCase.ActorRepository.Create(actor); err != nil {
		return output, errors.New(err.Error())
	}

	output.Actor.ID = actor.ID
	output.Actor.Name = actor.Name
	output.Actor.Picture = actor.Picture
	output.Actor.IsDeleted = actor.IsDeleted
	output.Actor.CreatedAt = actor.CreatedAt
	output.Actor.UpdatedAt = actor.UpdatedAt
	output.Actor.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *ActorUseCase) Find(input InputFindActorDto) (OutputFindActorDto, error) {
	output := OutputFindActorDto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Actor.ID = actor.ID
	output.Actor.Name = actor.Name
	output.Actor.Picture = actor.Picture
	output.Actor.IsDeleted = actor.IsDeleted
	output.Actor.CreatedAt = actor.CreatedAt
	output.Actor.UpdatedAt = actor.UpdatedAt
	output.Actor.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *ActorUseCase) Delete(input InputDeleteActorDto) (OutputDeleteActorDto, error) {
	output := OutputDeleteActorDto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if actor.IsDeleted {
		return output, errors.New("actor previously deleted")
	}

	actor.IsDeleted = true
	actor.DeletedAt = time.Now().Local().String()

	output.IsDeleted = actor.IsDeleted

	return output, errors.New(err.Error())
}

func (actorUseCase *ActorUseCase) Update(input InputUpdateActorDto) (OutputUpdateActorDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateActorDto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	actor.Name = input.Name
	actor.Picture = input.Picture

	isValid, err := actor.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	actor.UpdatedAt = timeNow

	err = actorUseCase.ActorRepository.Update(&actor)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Actor.ID = actor.ID
	output.Actor.Name = actor.Name
	output.Actor.Picture = actor.Picture
	output.Actor.IsDeleted = actor.IsDeleted
	output.Actor.CreatedAt = actor.CreatedAt
	output.Actor.UpdatedAt = actor.UpdatedAt
	output.Actor.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *ActorUseCase) IsDeleted(input InputIsDeletedActorDto) (OutputIsDeletedActorDto, error) {
	output := OutputIsDeletedActorDto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if actor.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (actorUseCase *ActorUseCase) FindAll() (OutputFindAllActorDto, error) {
	output := OutputFindAllActorDto{}

	actors, err := actorUseCase.ActorRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, actor := range actors {
		output.Actors = append(output.Actors, ActorDto{
			ID:        actor.ID,
			Name:      actor.Name,
			Picture:   actor.Picture,
			IsDeleted: actor.IsDeleted,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
			DeletedAt: actor.DeletedAt,
		})
	}

	return output, nil
}

func (actorUseCase *ActorUseCase) AddPictureToActor(input InputAddPictureToActorDto) (OutputAddPictureToActorDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToActorDto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	extension := strings.Replace(filepath.Ext(input.File.Handler.Filename), ".", "", -1)
	if strings.ToLower(extension) != "jpeg" && strings.ToLower(extension) != "jpg" {
		return output, errors.New("format not allowed")
	}

	_, name, size, extension, err := MoveFile(input.File.File, input.File.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	colorAverage, err := PictureAverageColor(name, extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := entity.NewFile(name, actor.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := actorUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	actor.Picture = picture.ID

	isValid, err := actor.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	actor.UpdatedAt = timeNow

	err = actorUseCase.ActorRepository.Update(&actor)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Actor.ID = actor.ID
	output.Actor.Name = actor.Name
	output.Actor.Picture = actor.Picture
	output.Actor.IsDeleted = actor.IsDeleted
	output.Actor.CreatedAt = actor.CreatedAt
	output.Actor.UpdatedAt = actor.UpdatedAt
	output.Actor.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *ActorUseCase) FindActorPictureToBase64(input InputFindActorPictureToBase64Dto) (OutputFindActorPictureToBase64Dto, error) {
	output := OutputFindActorPictureToBase64Dto{}

	actor, err := actorUseCase.ActorRepository.Find(input.ActorId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := actorUseCase.FileRepository.Find(actor.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64("/home/guilherme/Workspace/you-choose/cmd/upload/", picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Actor.ID = actor.ID
	output.Actor.Name = actor.Name
	output.Actor.Picture = pictureToBase64
	output.Actor.IsDeleted = actor.IsDeleted
	output.Actor.CreatedAt = actor.CreatedAt
	output.Actor.UpdatedAt = actor.UpdatedAt
	output.Actor.DeletedAt = actor.DeletedAt

	return output, nil
}
