package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        string
	Name      string
	Picture   string
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewTag(name string, picture string) (*Tag, error) {
	dateNow := time.Now()
	tag := &Tag{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
	}

	isValid, err := tag.Validate()
	if !isValid {
		return nil, err
	}

	return tag, nil
}

func (tag *Tag) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["name"] = tag.Name
	inputs["picture"] = tag.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
