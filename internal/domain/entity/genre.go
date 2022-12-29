package entity

import "github.com/google/uuid"

type Genre struct {
	ID      string
	Name    string
	Picture string
}

func NewGenre(name string, picture string) *Genre {
	return &Genre{
		ID:      uuid.New().String(),
		Name:    name,
		Picture: picture,
	}
}
