package usecases

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/google/uuid"
)

type FileUseCase struct {
	FileRepository entity.FileRepositoryInterface
}

func NewFileUseCase(fileRepository entity.FileRepositoryInterface) *FileUseCase {
	return &FileUseCase{
		FileRepository: fileRepository,
	}
}

func (fileUseCase *FileUseCase) Create(input InputCreateFileDto) (OutputCreateFileDto, error) {
	output := OutputCreateFileDto{}

	_, name, size, extension, err := MoveFile(input.File, input.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fileEntity, err := entity.NewFile(name, input.EntityId, size, extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := fileUseCase.FileRepository.Create(fileEntity); err != nil {
		return output, errors.New(err.Error())
	}

	output.File.ID = fileEntity.ID
	output.File.EntityId = fileEntity.EntityId
	output.File.Name = fileEntity.Name
	output.File.Size = fileEntity.Size
	output.File.Extension = fileEntity.Extension
	output.File.IsDeleted = fileEntity.IsDeleted
	output.File.CreatedAt = fileEntity.CreatedAt
	output.File.UpdatedAt = fileEntity.UpdatedAt
	output.File.DeletedAt = fileEntity.DeletedAt

	return output, nil
}

// func (fileUseCase *FileUseCase) Find(input InputFindFileDto) (OutputFindFileDto, error) {
// 	output := OutputFindFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	output.File.ID = file.ID
// 	output.File.EntityId = file.EntityId
// 	output.File.Name = file.Name
// 	output.File.Size = file.Size
// 	output.File.Extension = file.Extension
// 	output.File.IsDeleted = file.IsDeleted
// 	output.File.CreatedAt = file.CreatedAt
// 	output.File.UpdatedAt = file.UpdatedAt
// 	output.File.DeletedAt = file.DeletedAt

// 	return output, nil
// }

// func (fileUseCase *FileUseCase) Delete(input InputDeleteFileDto) (OutputDeleteFileDto, error) {
// 	timeNow := time.Now().Local().String()
// 	output := OutputDeleteFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	if file.IsDeleted {
// 		return output, errors.New("file previously deleted")
// 	}

// 	file.IsDeleted = true
// 	file.DeletedAt = timeNow

// 	output.IsDeleted = file.IsDeleted

// 	return output, errors.New(err.Error())
// }

// func (fileUseCase *FileUseCase) Update(input InputUpdateFileDto) (OutputUpdateFileDto, error) {
// 	timeNow := time.Now().Local().String()
// 	output := OutputUpdateFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	file.Name = input.Name

// 	isValid, err := file.Validate()
// 	if !isValid {
// 		return output, errors.New(err.Error())
// 	}

// 	file.UpdatedAt = timeNow

// 	err = fileUseCase.FileRepository.Update(&file)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	output.File.ID = file.ID
// 	output.File.EntityId = file.EntityId
// 	output.File.Name = file.Name
// 	output.File.Size = file.Size
// 	output.File.Extension = file.Extension
// 	output.File.IsDeleted = file.IsDeleted
// 	output.File.CreatedAt = file.CreatedAt
// 	output.File.UpdatedAt = file.UpdatedAt
// 	output.File.DeletedAt = file.DeletedAt

// 	return output, nil
// }

// func (fileUseCase *FileUseCase) IsDeleted(input InputIsDeletedFileDto) (OutputIsDeletedFileDto, error) {
// 	output := OutputIsDeletedFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	output.IsDeleted = false

// 	if file.IsDeleted {
// 		output.IsDeleted = true
// 	}

// 	return output, nil
// }

// func (fileUseCase *FileUseCase) FindAll() (OutputFindAllFileDto, error) {
// 	output := OutputFindAllFileDto{}

// 	files, err := fileUseCase.FileRepository.FindAll()
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	for _, file := range files {
// 		output.Files = append(output.Files, FileDto{
// 			ID:        file.ID,
// 			EntityId:  file.EntityId,
// 			Name:      file.Name,
// 			Size:      file.Size,
// 			Extension: file.Extension,
// 			IsDeleted: file.IsDeleted,
// 			CreatedAt: file.CreatedAt,
// 			UpdatedAt: file.UpdatedAt,
// 			DeletedAt: file.DeletedAt,
// 		})
// 	}

// 	return output, nil
// }

func MoveFile(file multipart.File, handler *multipart.FileHeader) (int64, string, int64, string, error) {
	path := "upload/"

	extension := filepath.Ext(handler.Filename)

	name := uuid.New().String()

	size := handler.Size

	fileCreate, err := os.Create(path + name + extension)

	defer file.Close()
	defer fileCreate.Close()

	if err != nil {
		return 0, "", 0, "", errors.New(err.Error())
	}

	fileWritten, err := io.Copy(fileCreate, file)
	if err != nil {
		return 0, "", 0, "", errors.New(err.Error())
	}

	extension = strings.Replace(filepath.Ext(handler.Filename), ".", "", -1)

	return fileWritten, name, size, extension, nil
}
