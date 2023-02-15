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
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewGenre(name string, picture string) (*Genre, error) {
	dateNow := time.Now()
	genre := &Genre{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
	}

	isValid, err := genre.Validate()
	if !isValid {
		return nil, err
	}

	return genre, nil
}

func (genre *Genre) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["name"] = genre.Name
	inputs["picture"] = genre.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
