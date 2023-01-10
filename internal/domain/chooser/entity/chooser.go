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

	isValidUserName, err := UserNameValidator(userName)
	if !isValidUserName {
		return nil, err
	}

	userNameEncrypt, err := EncryptString(userName)
	if err != nil {
		return nil, err
	}

	c.UserName = userNameEncrypt

	isValidPassword, err := PasswordValidator(password)
	if !isValidPassword {
		return nil, err
	}

	passwordEncrypt, err := EncryptString(password)
	if err != nil {
		return nil, err
	}

	c.Password = passwordEncrypt

	return c, nil
}

func (c *Chooser) Validate() error {
	inputs := make(map[string]string)

	inputs["first name"] = c.FirstName
	inputs["last name"] = c.LastName
	inputs["username"] = c.UserName
	inputs["picture"] = c.Picture
	inputs["password"] = c.Password

	for key, value := range inputs {
		if value == "" {
			message := "input " + key + " cannot be empty"
			return errors.New(message)
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

	if len(chars) < 10 {
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

		if (len(chars) - 1) != i {
			if chars[i] == chars[i+1] {
				return false, nil
			}
		}

		for y := 0; y < len(validCharacters); y++ {
			if chars[i] == validCharacters[y] {
				countValidCharacters = countValidCharacters + 1
			}
		}
	}

	if countUpper < 3 {
		return false, nil
	}

	if countLower < 2 {
		return false, nil
	}

	if countNumber < 2 {
		return false, nil
	}

	if countValidCharacters < 3 {
		return false, nil
	}

	return true, nil
}

func UserNameValidator(username string) (bool, error) {
	chars := []rune(username)

	if len(chars) < 4 {
		return false, nil
	}

	countUpper := 0
	countInvalidCharacters := 0

	for i := 0; i < len(chars); i++ {
		if unicode.IsUpper(chars[i]) {
			countUpper = countUpper + 1
		}

		if !unicode.IsLetter(chars[i]) && !unicode.IsNumber(chars[i]) {
			countInvalidCharacters = countInvalidCharacters + 1
		}
	}

	if countUpper != 0 {
		return false, nil
	}

	if countInvalidCharacters != 0 {
		return false, nil
	}

	return true, nil
}
