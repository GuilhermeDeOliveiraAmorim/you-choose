package domain

import (
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
)

type ChooserRepositoryInterface interface {
	Create(c *chooser.Chooser) (*chooser.Chooser, error)
	Update(c *chooser.Chooser) (*chooser.Chooser, error)
	FindById(id string) (*chooser.Chooser, error)
	DeleteById(id string) (*chooser.Chooser, error)
	FindAll() ([]*chooser.Chooser, error)
}
