package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Chooser struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Picture   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewChooser(firstName string, lastName string, userName string, picture string, password string) (*Chooser, error) {
	c := &Chooser{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Picture:   picture,
		CreatedAt: time.Now(),
	}

	err := c.Validate()

	if err != nil {
		return nil, err
	}

	pwd, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}

	c.Password = pwd

	return c, nil
}

func (c *Chooser) Validate() error {
	if c.FirstName == "" || c.LastName == "" || c.UserName == "" || c.Picture == "" || c.Password == "" {
		return errors.New("invalid entity")
	}
	return nil
}

func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
