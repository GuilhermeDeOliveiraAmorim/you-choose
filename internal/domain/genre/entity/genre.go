package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID        string
	Name      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewGenre(name string, picture string) (*Genre, error) {
	g := &Genre{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: time.Now(),
	}

	err := g.Validate()

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Genre) Validate() error {
	if g.Name == "" || g.Picture == "" {
		return errors.New("invalid entity")
	}
	return nil
}
