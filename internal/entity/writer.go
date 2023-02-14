package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Writer struct {
	ID        string
	Name      string
	Picture   string
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewWriter(name string, picture string) (*Writer, error) {
	dateNow := time.Now()
	writer := &Writer{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
	}

	isValid, err := writer.Validate()
	if !isValid {
		return nil, err
	}

	return writer, nil
}

func (writer *Writer) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["name"] = writer.Name
	inputs["picture"] = writer.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
