package entity

import "github.com/google/uuid"

type Director struct {
	ID      string
	Name    string
	Picture string
}

func NewDirector(name string, picture string) *Director {
	return &Director{
		ID:      uuid.New().String(),
		Name:    name,
		Picture: picture,
	}
}
