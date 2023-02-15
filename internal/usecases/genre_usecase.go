package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type GenreUseCase struct {
	GenreRepository entity.GenreRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
}

func NewGenreUseCase(actorRepository entity.GenreRepositoryInterface, genreRepository entity.MovieRepositoryInterface) *GenreUseCase {
	return &GenreUseCase{
		GenreRepository: actorRepository,
		MovieRepository: genreRepository,
	}
}

func (actorUseCase *GenreUseCase) Create(input InputCreateGenreDto) (OutputCreateGenreDto, error) {
	output := OutputCreateGenreDto{}

	actor, err := entity.NewGenre(input.Name, input.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := actorUseCase.GenreRepository.Create(actor); err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = actor.ID
	output.Genre.Name = actor.Name
	output.Genre.Picture = actor.Picture
	output.Genre.IsDeleted = actor.IsDeleted
	output.Genre.CreatedAt = actor.CreatedAt
	output.Genre.UpdatedAt = actor.UpdatedAt
	output.Genre.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *GenreUseCase) Find(input InputFindGenreDto) (OutputFindGenreDto, error) {
	output := OutputFindGenreDto{}

	actor, err := actorUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = actor.ID
	output.Genre.Name = actor.Name
	output.Genre.Picture = actor.Picture
	output.Genre.IsDeleted = actor.IsDeleted
	output.Genre.CreatedAt = actor.CreatedAt
	output.Genre.UpdatedAt = actor.UpdatedAt
	output.Genre.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *GenreUseCase) Delete(input InputDeleteGenreDto) (OutputDeleteGenreDto, error) {
	output := OutputDeleteGenreDto{}

	actor, err := actorUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if actor.IsDeleted {
		return output, errors.New("actor previously deleted")
	}

	actor.IsDeleted = true
	actor.DeletedAt = time.Now().Local().String()

	output.IsDeleted = actor.IsDeleted

	return output, nil
}

func (actorUseCase *GenreUseCase) Update(input InputUpdateGenreDto) (OutputUpdateGenreDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateGenreDto{}

	actor, err := actorUseCase.GenreRepository.Find(input.GenreId)
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

	err = actorUseCase.GenreRepository.Update(&actor)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Genre.ID = actor.ID
	output.Genre.Name = actor.Name
	output.Genre.Picture = actor.Picture
	output.Genre.IsDeleted = actor.IsDeleted
	output.Genre.CreatedAt = actor.CreatedAt
	output.Genre.UpdatedAt = actor.UpdatedAt
	output.Genre.DeletedAt = actor.DeletedAt

	return output, nil
}

func (actorUseCase *GenreUseCase) IsDeleted(input InputIsDeletedGenreDto) (OutputIsDeletedGenreDto, error) {
	output := OutputIsDeletedGenreDto{}

	actor, err := actorUseCase.GenreRepository.Find(input.GenreId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if actor.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (actorUseCase *GenreUseCase) FindAll() (OutputFindAllGenreDto, error) {
	output := OutputFindAllGenreDto{}

	actors, err := actorUseCase.GenreRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, actor := range actors {
		output.Genres = append(output.Genres, GenreDto{
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
