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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func NewDirector(name string, picture string) (*Director, error) {
	d := &Director{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
		IsDeleted: false,
	}

	err := d.Validate()

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Director) Validate() error {
	if d.Name == "" || d.Picture == "" {
		return errors.New("invalid entity")
	}
	return nil
}
