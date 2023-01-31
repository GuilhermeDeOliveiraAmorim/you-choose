package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Director struct {
	ID        string
	Name      string
	Picture   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	IsDeleted bool
}

func NewDirector(name string, picture string) (*Director, error) {
	dateNow := time.Now()
	director := &Director{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
		IsDeleted: false,
	}

	isValid, err := director.Validate()
	if !isValid {
		return nil, err
	}

	return director, nil
}

func (director *Director) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["name"] = director.Name
	inputs["picture"] = director.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
