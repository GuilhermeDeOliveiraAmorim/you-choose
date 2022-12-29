package entity

import "github.com/google/uuid"

type Writer struct {
	ID      string
	Name    string
	Picture string
}

func NewWriter(name string, picture string) *Writer {
	return &Writer{
		ID:      uuid.New().String(),
		Name:    name,
		Picture: picture,
	}
}
