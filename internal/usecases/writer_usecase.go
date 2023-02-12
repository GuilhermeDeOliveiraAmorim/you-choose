package usecases

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WriterUseCase struct {
	WriterRepository entity.WriterRepositoryInterface
	MovieRepository  entity.MovieRepositoryInterface
}

func NewWriterUseCase(writerRepository entity.WriterRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *WriterUseCase {
	return &WriterUseCase{
		WriterRepository: writerRepository,
		MovieRepository:  movieRepository,
	}
}

func (writerUseCase *WriterUseCase) Create(input InputCreateWriterDto) (OutputCreateWriterDto, error) {
	output := OutputCreateWriterDto{}

	writer, err := entity.NewWriter(input.Name, input.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := writerUseCase.WriterRepository.Create(writer); err != nil {
		return output, errors.New(err.Error())
	}

	output.Writer.ID = writer.ID
	output.Writer.Name = writer.Name
	output.Writer.Picture = writer.Picture
	output.Writer.IsDeleted = writer.IsDeleted
	output.Writer.CreatedAt = writer.CreatedAt
	output.Writer.UpdatedAt = writer.UpdatedAt
	output.Writer.DeletedAt = writer.DeletedAt

	return output, nil
}

func (writerUseCase *WriterUseCase) Find(input InputFindWriterDto) (OutputFindWriterDto, error) {
	output := OutputFindWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Writer.ID = writer.ID
	output.Writer.Name = writer.Name
	output.Writer.Picture = writer.Picture
	output.Writer.IsDeleted = writer.IsDeleted
	output.Writer.CreatedAt = writer.CreatedAt
	output.Writer.UpdatedAt = writer.UpdatedAt
	output.Writer.DeletedAt = writer.DeletedAt

	return output, nil
}

func (writerUseCase *WriterUseCase) Delete(input InputDeleteWriterDto) (OutputDeleteWriterDto, error) {
	output := OutputDeleteWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if writer.IsDeleted {
		return output, errors.New("writer previously deleted")
	}

	writer.IsDeleted = true
	writer.DeletedAt = time.Now().Local().String()

	output.IsDeleted = writer.IsDeleted

	return output, errors.New(err.Error())
}

func (writerUseCase *WriterUseCase) Update(input InputUpdateWriterDto) (OutputUpdateWriterDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	writer.Name = input.Name
	writer.Picture = input.Picture

	isValid, err := writer.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	writer.UpdatedAt = timeNow

	err = writerUseCase.WriterRepository.Update(&writer)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Writer.ID = writer.ID
	output.Writer.Name = writer.Name
	output.Writer.Picture = writer.Picture
	output.Writer.IsDeleted = writer.IsDeleted
	output.Writer.CreatedAt = writer.CreatedAt
	output.Writer.UpdatedAt = writer.UpdatedAt
	output.Writer.DeletedAt = writer.DeletedAt

	return output, nil
}

func (writerUseCase *WriterUseCase) IsDeleted(input InputIsDeletedWriterDto) (OutputIsDeletedWriterDto, error) {
	output := OutputIsDeletedWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if writer.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (writerUseCase *WriterUseCase) FindAll() (OutputFindAllWriterDto, error) {
	output := OutputFindAllWriterDto{}

	writers, err := writerUseCase.WriterRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, writer := range writers {
		output.Writers = append(output.Writers, WriterDto{
			ID:        writer.ID,
			Name:      writer.Name,
			Picture:   writer.Picture,
			IsDeleted: writer.IsDeleted,
			CreatedAt: writer.CreatedAt,
			UpdatedAt: writer.UpdatedAt,
			DeletedAt: writer.DeletedAt,
		})
	}

	return output, nil
}

func (writerUseCase *WriterUseCase) FindAllWriterMovies(input InputFindAllWriterMoviesDto) (OutputFindAllWriterMoviesDto, error) {
	output := OutputFindAllWriterMoviesDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Writer = WriterDto{
		ID:        writer.ID,
		Name:      writer.Name,
		Picture:   writer.Picture,
		IsDeleted: writer.IsDeleted,
		CreatedAt: writer.CreatedAt,
		UpdatedAt: writer.UpdatedAt,
		DeletedAt: writer.DeletedAt,
	}

	movies, err := writerUseCase.WriterRepository.FindAllWriterMovies(input.ID)
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, movie := range movies {
		output.Movies = append(output.Movies, MovieDto{
			ID:              movie.ID,
			Title:           movie.Title,
			Synopsis:        movie.Synopsis,
			ImdbRating:      movie.ImdbRating,
			Votes:           movie.Votes,
			YouChooseRating: movie.YouChooseRating,
			Poster:          movie.Poster,
			CreatedAt:       movie.CreatedAt,
			UpdatedAt:       movie.UpdatedAt,
			DeletedAt:       movie.DeletedAt,
			IsDeleted:       movie.IsDeleted,
		})
	}

	return output, nil
}
