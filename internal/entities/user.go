package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
)

type User struct {
	SharedEntity
	Name    string `json:"name"`
	Login   Login  `json:"login"`
	IsAdmin bool   `json:"is_admin"`
}

func NewUser(name string, login Login) (*User, []exceptions.ProblemDetails) {
	validationErrors := ValidateUser(name)

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &User{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Login:        login,
		IsAdmin:      false,
	}, nil
}

func ValidateUser(name string) []exceptions.ProblemDetails {
	var validationErrors []exceptions.ProblemDetails

	if name == "" {
		validationErrors = append(validationErrors, exceptions.ProblemDetails{
			Type:     "Validation Error",
			Title:    "User name cannot be empty",
			Status:   400,
			Detail:   "User name is required",
			Instance: exceptions.RFC400,
		})
	}

	if len(name) > 100 {
		validationErrors = append(validationErrors, exceptions.ProblemDetails{
			Type:     "Validation Error",
			Title:    "User name is too long",
			Status:   400,
			Detail:   "User name must not exceed 100 characters",
			Instance: exceptions.RFC400,
		})
	}

	return validationErrors
}

func (u *User) ChangeName(newName string) []exceptions.ProblemDetails {
	validationErrors := ValidateUser(newName)

	if len(validationErrors) > 0 {
		return validationErrors
	}

	timeNow := time.Now()

	u.UpdatedAt = &timeNow
	u.Name = newName

	return validationErrors
}
