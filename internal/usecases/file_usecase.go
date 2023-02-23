package usecases

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
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

	file, err := MoveFile(input.File, input.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fmt.Println("file", file)

	fileEntity, err := entity.NewFile(input.Name, input.EntityId, input.Handler.Filename, "")
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

func MoveFile(file multipart.File, handler *multipart.FileHeader) (int64, error) {
	fileCreate, err := os.Create(handler.Filename)
	defer file.Close()
	defer fileCreate.Close()

	if err != nil {
		return 0, errors.New(err.Error())
	}

	fileWritten, err := io.Copy(fileCreate, file)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return fileWritten, nil
}
