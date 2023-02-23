package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        string
	EntityId  string
	Name      string
	Size      string
	Extension string
	IsDeleted bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewFile(name string, entityId string, size string, extension string) (*File, error) {
	dateNow := time.Now()
	file := &File{
		ID:        uuid.New().String(),
		EntityId:  entityId,
		Name:      name,
		Size:      size,
		Extension: extension,
		IsDeleted: false,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
	}

	isValid, err := file.Validate()
	if !isValid {
		return nil, err
	}

	return file, nil
}

func (file *File) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["entity_id"] = file.EntityId
	inputs["name"] = file.Name
	inputs["size"] = file.Size
	inputs["extension"] = file.Extension

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
