package entity

import "github.com/google/uuid"

type Chooser struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Picture   string
}

func NewChooser(firstName string, lastName string, userName string, picture string) *Chooser {
	return &Chooser{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Picture:   picture,
	}
}
