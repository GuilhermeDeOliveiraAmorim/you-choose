package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WriterUseCase struct {
	WriterRepository entity.WriterRepositoryInterface
	MovieRepository  entity.MovieRepositoryInterface
	FileRepository   entity.FileRepositoryInterface
}

func NewWriterUseCase(
	writerRepository entity.WriterRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *WriterUseCase {
	return &WriterUseCase{
		WriterRepository: writerRepository,
		MovieRepository:  movieRepository,
		FileRepository:   fileRepository,
	}
}

func (writerUseCase *WriterUseCase) Create(input InputCreateWriterDto) (OutputCreateWriterDto, error) {
	output := OutputCreateWriterDto{}

	writer, err := entity.NewWriter(input.Name)
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

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	inputFindPicture := InputFindWriterPictureToBase64Dto{
		WriterId: writer.ID,
	}

	picture, err := writerUseCase.FindWriterPictureToBase64(inputFindPicture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	file, err := writerUseCase.FileRepository.Find(writer.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fileDto := FileDto{
		ID:           file.ID,
		EntityId:     file.EntityId,
		Name:         file.Name,
		Size:         file.Size,
		Extension:    file.Extension,
		AverageColor: file.AverageColor,
		IsDeleted:    file.IsDeleted,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
		DeletedAt:    file.DeletedAt,
	}

	output.Writer.ID = writer.ID
	output.Writer.Name = writer.Name
	output.Writer.Picture = picture.Writer.Picture
	output.Writer.IsDeleted = writer.IsDeleted
	output.Writer.CreatedAt = writer.CreatedAt
	output.Writer.UpdatedAt = writer.UpdatedAt
	output.Writer.DeletedAt = writer.DeletedAt
	output.Writer.File = fileDto

	return output, nil
}

func (writerUseCase *WriterUseCase) Delete(input InputDeleteWriterDto) (OutputDeleteWriterDto, error) {
	output := OutputDeleteWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if writer.IsDeleted {
		return output, errors.New("writer previously deleted")
	}

	writer.IsDeleted = true
	writer.DeletedAt = time.Now().Local().String()

	err = writerUseCase.WriterRepository.Update(&writer)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = writer.IsDeleted

	return output, nil
}

func (writerUseCase *WriterUseCase) Update(input InputUpdateWriterDto) (OutputUpdateWriterDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
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

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
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

		inputFindPicture := InputFindWriterPictureToBase64Dto{
			WriterId: writer.ID,
		}

		picture, err := writerUseCase.FindWriterPictureToBase64(inputFindPicture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		file, err := writerUseCase.FileRepository.Find(writer.Picture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		fileDto := FileDto{
			ID:           file.ID,
			EntityId:     file.EntityId,
			Name:         file.Name,
			Size:         file.Size,
			Extension:    file.Extension,
			AverageColor: file.AverageColor,
			IsDeleted:    file.IsDeleted,
			CreatedAt:    file.CreatedAt,
			UpdatedAt:    file.UpdatedAt,
			DeletedAt:    file.DeletedAt,
		}

		output.Writers = append(output.Writers, WriterDto{
			ID:        writer.ID,
			Name:      writer.Name,
			Picture:   picture.Writer.Picture,
			IsDeleted: writer.IsDeleted,
			CreatedAt: writer.CreatedAt,
			UpdatedAt: writer.UpdatedAt,
			DeletedAt: writer.DeletedAt,
			File:      fileDto,
		})
	}

	return output, nil
}

func (writerUseCase *WriterUseCase) AddPictureToWriter(input InputAddPictureToWriterDto) (OutputAddPictureToWriterDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToWriterDto{}

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
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

	picture, err := entity.NewFile(name, writer.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := writerUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	writer.Picture = picture.ID

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

func (writerUseCase *WriterUseCase) FindWriterPictureToBase64(input InputFindWriterPictureToBase64Dto) (OutputFindWriterPictureToBase64Dto, error) {
	output := OutputFindWriterPictureToBase64Dto{}

	writer, err := writerUseCase.WriterRepository.Find(input.WriterId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := writerUseCase.FileRepository.Find(writer.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64("/home/guilhermeamorim/Workspace/estudo/you-choose/cmd/upload/", picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Writer.ID = writer.ID
	output.Writer.Name = writer.Name
	output.Writer.Picture = pictureToBase64
	output.Writer.IsDeleted = writer.IsDeleted
	output.Writer.CreatedAt = writer.CreatedAt
	output.Writer.UpdatedAt = writer.UpdatedAt
	output.Writer.DeletedAt = writer.DeletedAt

	return output, nil
}
