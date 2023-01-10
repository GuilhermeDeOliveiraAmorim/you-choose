package domain

import (
	"errors"
	"time"
	"unicode"

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

	validCharacters := []rune("!@#$%&*?")

	countUpper := 0
	countLower := 0
	countNumber := 0
	countValidCharacters := 0

	if len(chars) < 8 {
		return false, nil
	}

	for i := 0; i < len(chars); i++ {
		if unicode.IsUpper(chars[i]) {
			countUpper = countUpper + 1
		}

		if unicode.IsLower(chars[i]) {
			countLower = countLower + 1
		}

		if unicode.IsNumber(chars[i]) {
			countNumber = countNumber + 1
		}
		for y := 0; y < len(validCharacters); y++ {
			if chars[i] == validCharacters[y] {
				countValidCharacters = countValidCharacters + 1
			}
		}
	}

	if countUpper < 2 {
		return false, nil
	}

	if countLower < 2 {
		return false, nil
	}

	if countNumber < 2 {
		return false, nil
	}

	if countValidCharacters < 2 {
		return false, nil
	}

	return true, nil
}
