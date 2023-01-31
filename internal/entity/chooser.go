package entity

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
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	IsDeleted bool
}

func NewChooser(firstName string, lastName string, userName string, picture string) (*Chooser, error) {
	dateNow := time.Now()
	chooser := &Chooser{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Picture:   picture,
		CreatedAt: dateNow.Local().String(),
		UpdatedAt: dateNow.Local().String(),
		DeletedAt: dateNow.Local().String(),
		IsDeleted: false,
	}

	isValid, err := chooser.Validate()
	if !isValid {
		return nil, err
	}

	return chooser, nil
}

func (chooser *Chooser) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["first name"] = chooser.FirstName
	inputs["last name"] = chooser.LastName
	inputs["username"] = chooser.UserName
	inputs["picture"] = chooser.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	isValidUserName, err := ValidateUserName(chooser.UserName)
	if !isValidUserName {
		return false, err
	}

	return true, nil
}

func EncryptString(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyWord(word string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(word))
	return err == nil
}

func ValidatePassword(password string) (bool, error) {
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
		return false, errors.New("your password must be have 3 or more uppercase letter")
	}

	if countLower < 2 {
		return false, errors.New("your password must be have 2 or more lowercase letter")
	}

	if countNumber < 2 {
		return false, errors.New("your password must be have 2 or more numbers")
	}

	if countValidCharacters < 3 {
		return false, errors.New("your password must be have 3 or more characters in sample: ! @ # $ % & * ?")
	}

	return true, nil
}

func ValidateUserName(username string) (bool, error) {
	chars := []rune(username)

	if len(chars) < 4 {
		return false, errors.New("your username must be more than 3 characters")
	}

	countUpper := 0
	countInvalidCharacters := 0
	countSpaces := 0

	for i := 0; i < len(chars); i++ {
		if unicode.IsSpace(chars[i]) {
			countSpaces = countSpaces + 1
		}

		if unicode.IsUpper(chars[i]) {
			countUpper = countUpper + 1
		}

		if !unicode.IsLetter(chars[i]) && !unicode.IsNumber(chars[i]) {
			countInvalidCharacters = countInvalidCharacters + 1
		}
	}

	if countSpaces != 0 {
		return false, errors.New("your username must not have any blank spaces")
	}

	if countUpper != 0 {
		return false, errors.New("your username must only contain lowercase letters")
	}

	if countInvalidCharacters != 0 {
		return false, errors.New("your username must only contain letters and numbers")
	}

	return true, nil
}
