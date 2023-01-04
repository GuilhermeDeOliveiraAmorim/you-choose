package domain

import (
	"context"

	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
)

type ChooserRepositoryInterface interface {
	Add(ctx context.Context, c *chooser.Chooser) (*chooser.Chooser, error)
	Find(ctx context.Context, id string) (*chooser.Chooser, error)
	FindAll(ctx context.Context) ([]*chooser.Chooser, error)
}
