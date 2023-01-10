package domain

import (
	"context"

	writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
)

type WriterRepositoryInterface interface {
	Create(ctx context.Context, a *writer.Writer) (*writer.Writer, error)
	Update(ctx context.Context, a *writer.Writer) (*writer.Writer, error)
	FindById(ctx context.Context, id string) (*writer.Writer, error)
	DeleteById(ctx context.Context, id string) (*writer.Writer, error)
	FindAll(ctx context.Context) ([]*writer.Writer, error)
}
