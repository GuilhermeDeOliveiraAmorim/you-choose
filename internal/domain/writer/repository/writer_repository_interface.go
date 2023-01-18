package domain

import (
	writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
)

type WriterRepositoryInterface interface {
	Create(a *writer.Writer) (*writer.Writer, error)
	Update(a *writer.Writer) (*writer.Writer, error)
	FindById(id string) (*writer.Writer, error)
	DeleteById(id string) (*writer.Writer, error)
	FindAll() ([]*writer.Writer, error)
}
