package usecases

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"

type FileUseCase struct {
	FileRepository entity.FileRepositoryInterface
}

func NewFileUseCase(fileRepository entity.FileRepositoryInterface) *FileUseCase {
	return &FileUseCase{
		FileRepository: fileRepository,
	}
}
