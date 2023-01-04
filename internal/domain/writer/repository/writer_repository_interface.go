package domain

import (
	"context"

	writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
)

type WriterRepositoryInterface interface {
	Create(ctx context.Context, w *writer.Writer) (*writer.Writer, error)
	Find(ctx context.Context, id string) (*writer.Writer, error)
	FindAll(ctx context.Context) ([]*writer.Writer, error)
}
