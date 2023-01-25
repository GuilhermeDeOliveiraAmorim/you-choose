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
	Password  string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewChooser(firstName string, lastName string, userName string, picture string, password string) (*Chooser, error) {
	dateNow := time.Now()
	chooser := &Chooser{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Picture:   picture,
		Password:  password,
		IsDeleted: false,
		CreatedAt: dateNow,
		UpdatedAt: dateNow,
		DeletedAt: dateNow,
	}

	isValidChooser, err := chooser.Validate()
	if !isValidChooser {
		return nil, err
	}

	isValidUserName, err := ValidateUserName(userName)
	if !isValidUserName {
		return nil, err
	}

	// userNameEncrypted, err := EncryptString(userName)
	// if err != nil {
	// 	return nil, err
	// }

	// chooser.UserName = userNameEncrypted

	isValidPassword, err := ValidatePassword(password)
	if !isValidPassword {
		return nil, err
	}

	passwordEncrypted, err := EncryptString(password)
	if err != nil {
		return nil, err
	}

	chooser.Password = passwordEncrypted

	return chooser, nil
}

func (c *Chooser) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["first name"] = c.FirstName
	inputs["last name"] = c.LastName
	inputs["username"] = c.UserName
	inputs["picture"] = c.Picture
	inputs["password"] = c.Password

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
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
