package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Actor struct {
	ID        string
	Name      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func NewActor(name string, picture string) (*Actor, error) {
	a := &Actor{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
		IsDeleted: false,
	}

	err := a.Validate()

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Actor) Validate() error {
	if a.Name == "" || a.Picture == "" {
		return errors.New("invalid entity")
	}

	return nil
}
