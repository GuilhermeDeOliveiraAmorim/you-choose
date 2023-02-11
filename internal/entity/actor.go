package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Actor struct {
	ID        string
	Name      string
	Picture   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	IsDeleted bool
}

func NewActor(name string, picture string) (*Actor, error) {
	dateNow := time.Now()
	actor := &Actor{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
	}

	isValid, err := actor.Validate()
	if !isValid {
		return nil, err
	}

	return actor, nil
}

func (actor *Actor) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["name"] = actor.Name
	inputs["picture"] = actor.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
