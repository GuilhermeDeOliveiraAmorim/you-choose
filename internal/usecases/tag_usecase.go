package usecases

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type TagUseCase struct {
	TagRepository       entity.TagRepositoryInterface
	MovieListRepository entity.MovieListRepositoryInterface
	FileRepository      entity.FileRepositoryInterface
}

func NewTagUseCase(
	tagRepository entity.TagRepositoryInterface,
	movieListRepository entity.MovieListRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *TagUseCase {
	return &TagUseCase{
		TagRepository:       tagRepository,
		MovieListRepository: movieListRepository,
		FileRepository:      fileRepository,
	}
}

func (tagUseCase *TagUseCase) Create(input InputCreateTagDto) (OutputCreateTagDto, error) {
	output := OutputCreateTagDto{}

	doesThisTagAlreadyExist, err := tagUseCase.TagRepository.FindTagByName(input.Name)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if doesThisTagAlreadyExist.Name != "" {
		return output, errors.New("this tag already exists")
	}

	tag, err := entity.NewTag(input.Name)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := tagUseCase.TagRepository.Create(tag); err != nil {
		return output, errors.New(err.Error())
	}

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = tag.Picture
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt

	return output, nil
}

func (tagUseCase *TagUseCase) Find(input InputFindTagDto) (OutputFindTagDto, error) {
	output := OutputFindTagDto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if tag.ID == "" {
		return output, errors.New("tag not found")
	}

	inputFindPicture := InputFindTagPictureToBase64Dto{
		TagId: tag.ID,
	}

	picture, err := tagUseCase.FindTagPictureToBase64(inputFindPicture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	file, err := tagUseCase.FileRepository.Find(tag.Picture)
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

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = picture.Tag.Picture
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt
	output.Tag.File = fileDto

	return output, nil
}

func (tagUseCase *TagUseCase) Delete(input InputDeleteTagDto) (OutputDeleteTagDto, error) {
	output := OutputDeleteTagDto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if tag.IsDeleted {
		return output, errors.New("tag previously deleted")
	}

	tag.IsDeleted = true
	tag.DeletedAt = time.Now().Local().String()

	err = tagUseCase.TagRepository.Delete(&tag)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = tag.IsDeleted

	return output, nil
}

func (tagUseCase *TagUseCase) Update(input InputUpdateTagDto) (OutputUpdateTagDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputUpdateTagDto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	tag.Name = input.Name
	tag.Picture = input.Picture

	isValid, err := tag.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	tag.UpdatedAt = timeNow

	err = tagUseCase.TagRepository.Update(&tag)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = tag.Picture
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt

	return output, nil
}

func (tagUseCase *TagUseCase) IsDeleted(input InputIsDeletedTagDto) (OutputIsDeletedTagDto, error) {
	output := OutputIsDeletedTagDto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.IsDeleted = false

	if tag.IsDeleted {
		output.IsDeleted = true
	}

	return output, nil
}

func (tagUseCase *TagUseCase) FindAll() (OutputFindAllTagDto, error) {
	output := OutputFindAllTagDto{}

	tags, err := tagUseCase.TagRepository.FindAll()
	if err != nil {
		return output, errors.New(err.Error())
	}

	for _, tag := range tags {

		inputFindPicture := InputFindTagPictureToBase64Dto{
			TagId: tag.ID,
		}

		picture, err := tagUseCase.FindTagPictureToBase64(inputFindPicture)
		if err != nil {
			return output, errors.New(err.Error())
		}

		file, err := tagUseCase.FileRepository.Find(tag.Picture)
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

		output.Tags = append(output.Tags, TagDto{
			ID:        tag.ID,
			Name:      tag.Name,
			Picture:   picture.Tag.Picture,
			IsDeleted: tag.IsDeleted,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
			DeletedAt: tag.DeletedAt,
			File:      fileDto,
		})
	}

	return output, nil
}

func (tagUseCase *TagUseCase) FindTagByName(input InputFindTagByNameDto) (OutputFindTagByNameDto, error) {
	output := OutputFindTagByNameDto{}

	tag, err := tagUseCase.TagRepository.FindTagByName(input.TagName)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = tag.Picture
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt

	return output, nil
}

func (tagUseCase *TagUseCase) AddPictureToTag(input InputAddPictureToTagDto) (OutputAddPictureToTagDto, error) {
	timeNow := time.Now().Local().String()
	output := OutputAddPictureToTagDto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
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

	picture, err := entity.NewFile(name, tag.ID, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := tagUseCase.FileRepository.Create(picture); err != nil {
		return output, errors.New(err.Error())
	}

	tag.Picture = picture.ID

	isValid, err := tag.Validate()
	if !isValid {
		return output, errors.New(err.Error())
	}

	tag.UpdatedAt = timeNow

	err = tagUseCase.TagRepository.Update(&tag)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = tag.Picture
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt

	return output, nil
}

func (tagUseCase *TagUseCase) FindTagPictureToBase64(input InputFindTagPictureToBase64Dto) (OutputFindTagPictureToBase64Dto, error) {
	output := OutputFindTagPictureToBase64Dto{}

	tag, err := tagUseCase.TagRepository.Find(input.TagId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	picture, err := tagUseCase.FileRepository.Find(tag.Picture)
	if err != nil {
		return output, errors.New(err.Error())
	}

	pictureToBase64, err := PictureToBase64(dotenv, picture.Name, picture.Extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.Tag.ID = tag.ID
	output.Tag.Name = tag.Name
	output.Tag.Picture = pictureToBase64
	output.Tag.IsDeleted = tag.IsDeleted
	output.Tag.CreatedAt = tag.CreatedAt
	output.Tag.UpdatedAt = tag.UpdatedAt
	output.Tag.DeletedAt = tag.DeletedAt

	return output, nil
}
