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
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.Validate()

	if err != nil {
		return nil, err
	}

	un, err := EncryptString(userName)
	if err != nil {
		return nil, err
	}

	c.UserName = un

	isAValidPassword, err := PasswordValidator(password)
	if !isAValidPassword {
		return nil, err
	}

	pwd, err := EncryptString(password)
	if err != nil {
		return nil, err
	}

	c.Password = pwd

	return c, nil
}

func (c *Chooser) Validate() error {
	inputs := make(map[string]string)

	inputs["first name"] = c.FirstName
	inputs["last name"] = c.LastName
	inputs["username"] = c.UserName
	inputs["picture"] = c.Picture
	inputs["password"] = c.Password

	for k, v := range inputs {
		if v == "" {
			msg := "input " + k + " cannot be empty"
			return errors.New(msg)
		}
	}

	return nil
}

func EncryptString(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func PasswordValidator(password string) (bool, error) {
	chars := []rune(password)
	if len(chars) > 8 {
		return false, errors.New(">")
	}
	if len(chars) < 4 {
		return false, errors.New("<")
	}
	return true, nil
}
