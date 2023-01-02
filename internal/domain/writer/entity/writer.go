package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Writer struct {
	ID        string
	Name      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWriter(name string, picture string) (*Writer, error) {
	w := &Writer{
		ID:        uuid.New().String(),
		Name:      name,
		Picture:   picture,
		CreatedAt: time.Now(),
	}

	err := w.Validate()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Writer) Validate() error {
	if w.Name == "" || w.Picture == "" {
		return errors.New("invalid entity")
	}
	return nil
}
