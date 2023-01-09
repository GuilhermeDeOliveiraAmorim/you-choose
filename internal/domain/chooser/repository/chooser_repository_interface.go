package domain

import (
	"context"

	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
)

type ChooserRepositoryInterface interface {
	Create(ctx context.Context, c *chooser.Chooser) (*chooser.Chooser, error)
	Update(ctx context.Context, a *chooser.Chooser) (*chooser.Chooser, error)
	FindById(ctx context.Context, id string) (*chooser.Chooser, error)
	DeleteById(ctx context.Context, id string) (*chooser.Chooser, error)
	FindAll(ctx context.Context) ([]*chooser.Chooser, error)
}
