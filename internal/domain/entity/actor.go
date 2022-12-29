package entity

import "github.com/google/uuid"

type Actor struct {
	ID      string
	Name    string
	Picture string
}

func NewActor(name string, picture string) *Actor {
	return &Actor{
		ID:      uuid.New().String(),
		Name:    name,
		Picture: picture,
	}
}
