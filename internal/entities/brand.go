package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Brand struct {
	SharedEntity
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	VotesCount int    `json:"votes_count"`
}

func NewBrand(name, logo string) (*Brand, []util.ProblemDetails) {
	return &Brand{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Logo:         logo,
		VotesCount:   0,
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

func (b *Brand) IncrementVotesCount() {
	timeNow := time.Now()
	b.UpdatedAt = &timeNow

	b.VotesCount++
}

func (b *Brand) DecrementVotesCount() {
	if b.VotesCount > 0 {
		timeNow := time.Now()
		b.UpdatedAt = &timeNow

		b.VotesCount--
	}
}
