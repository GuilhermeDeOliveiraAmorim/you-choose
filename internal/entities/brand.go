package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Brand struct {
	SharedEntity
	Votable
	Name       string `json:"name"`
	Logo       string `json:"logo"`
}

func NewBrand(name, logo string) (*Brand, []util.ProblemDetails) {
	return &Brand{
		SharedEntity: *NewSharedEntity(),
		Votable:      *NewVotable(),
		Name:         name,
		Logo:         logo,
	}, nil
}

func (b *Brand) AddLogo(logo string) {
	b.Logo = logo
}

func (b *Brand) UpdateLogo(logo string) {
	timeNow := time.Now()
	b.UpdatedAt = &timeNow

	b.Logo = logo
}

func (b *Brand) Equals(brand Brand) bool {
	return b.Name == brand.Name
}

