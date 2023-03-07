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
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewDirector(name string) (*Director, error) {
	dateNow := time.Now()
	director := &Director{
		ID:        uuid.New().String(),
		Name:      name,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
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

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
