package entities

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"golang.org/x/crypto/bcrypt"
)

const (
	SpecialCharacters = "@#$%&*"
	UpperCaseLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCaseLetters  = "abcdefghijklmnopqrstuvwxyz"
	Digits            = "0123456789"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin(email, password string) (*Login, []exceptions.ProblemDetails) {
	validationErrors := ValidateLogin(email, password)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &Login{
		Email:    email,
		Password: password,
	}, nil
}

func ValidateLogin(email, password string) []exceptions.ProblemDetails {
	var validationErrors []exceptions.ProblemDetails

	if !isValidEmail(email) {
		validationErrors = append(validationErrors, exceptions.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Invalid email",
			Status:   400,
			Detail:   "Email is invalid",
			Instance: exceptions.RFC400,
		})
	}

	if !isValidPassword(password) {
		validationErrors = append(validationErrors, exceptions.ProblemDetails{
			Type:     "Validation Error",
			Title:    "Invalid password",
			Status:   400,
			Detail:   "Password must be at least 6 characters long, contain at least one uppercase letter, one lowercase letter, one digit, and one special character",
			Instance: exceptions.RFC400,
		})
	}

	return validationErrors
}

func isValidEmail(email string) bool {
	emailPattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func isValidPassword(password string) bool {
	return hasMinimumLength(password, 6) &&
		hasUpperCaseLetter(password) &&
		hasLowerCaseLetter(password) &&
		hasDigit(password) &&
		hasSpecialCharacter(password)
}

func hasMinimumLength(password string, length int) bool {
	return len(password) >= length
}

func hasUpperCaseLetter(password string) bool {
	return strings.ContainsAny(password, UpperCaseLetters)
}

func hasLowerCaseLetter(password string) bool {
	return strings.ContainsAny(password, LowerCaseLetters)
}

func hasDigit(password string) bool {
	return strings.ContainsAny(password, Digits)
}

func hasSpecialCharacter(password string) bool {
	return strings.ContainsAny(password, SpecialCharacters)
}

func hashWithBcrypt(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func hashEmail(email string) string {
	key := []byte(config.SECRETS_VAR.JWT_SECRET)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}

func compareBcryptHash(hashedData, data string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data))
	return err == nil
}

func (lo *Login) HashEmail() error {
	lo.Email = hashEmail(lo.Email)
	return nil
}

func (lo *Login) HashPassword() error {
	hashedPassword, err := hashWithBcrypt(lo.Password)
	if err != nil {
		return err
	}

	lo.Password = string(hashedPassword)

	return nil
}

func (lo *Login) VerifyEmail(email string) bool {
	return lo.Email == hashEmail(email)
}

func (lo *Login) VerifyPassword(password string) bool {
	return compareBcryptHash(lo.Password, password)
}

func (lo *Login) ChangeEmail(newEmail string) {
	lo.Email = newEmail
}

func (lo *Login) ChangePassword(newPassword string) {
	lo.Password = newPassword
}

func (lo *Login) Equals(other *Login) bool {
	return lo.Email == other.Email
}
