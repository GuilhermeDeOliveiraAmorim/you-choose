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
	actor := &Actor{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
		IsDeleted: false,
	}

	err := actor.Validate()

	if err != nil {
		return nil, err
	}

	return actor, nil
}

func (actor *Actor) Validate() error {
	if actor.Name == "" || actor.Picture == "" {
		return errors.New("invalid entity")
	}

	return nil
}
