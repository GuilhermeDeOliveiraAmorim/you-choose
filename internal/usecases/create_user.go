package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/presenters"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type CreateUserInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserUseCase(
	UserRepository repositories.UserRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: UserRepository,
	}
}

func (c *CreateUserUseCase) Execute(input CreateUserInputDto) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	email, hashEmailWithHMACErr := c.UserRepository.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error hashing email",
				Status:   500,
				Detail:   "An error occurred while hashing the email with HMAC.",
				Instance: exceptions.RFC500,
			},
		}
	}

	userEmailExists, userEmailExistsErr := c.UserRepository.ThisUserEmailExists(email)
	if userEmailExists {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Conflict",
				Title:    "Email already exists",
				Status:   409,
				Detail:   "The provided email is already registered.",
				Instance: exceptions.RFC409,
			},
		}
	} else if strings.Compare(userEmailExistsErr.Error(), "email not found") != 0 {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error checking user email existence",
				Status:   500,
				Detail:   "An error occurred while checking if the email already exists.",
				Instance: exceptions.RFC500,
			},
		}
	}

	userNameExists, userNameExistsErr := c.UserRepository.ThisUserNameExists(input.Name)
	if userNameExists {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Conflict",
				Title:    "Username already exists",
				Status:   409,
				Detail:   "The provided username is already taken.",
				Instance: exceptions.RFC409,
			},
		}
	} else if strings.Compare(userNameExistsErr.Error(), "username not found") != 0 {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error checking user name existence",
				Status:   500,
				Detail:   "An error occurred while checking if the username already exists.",
				Instance: exceptions.RFC500,
			},
		}
	}

	newLogin, newLoginErr := entities.NewLogin(input.Email, input.Password)
	if newLoginErr != nil {
		return presenters.SuccessOutputDTO{}, newLoginErr
	}

	encryptEmailErr := newLogin.EncryptEmail()
	if encryptEmailErr != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error encrypting email",
				Status:   500,
				Detail:   "An error occurred while encrypting the email address.",
				Instance: exceptions.RFC500,
			},
		}
	}

	encryptPasswordErr := newLogin.EncryptPassword()
	if encryptPasswordErr != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error encrypting password",
				Status:   500,
				Detail:   "An error occurred while encrypting the password.",
				Instance: exceptions.RFC500,
			},
		}
	}

	newUser, newUserErr := entities.NewUser(input.Name, *newLogin)
	if newUserErr != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating user",
				Status:   500,
				Detail:   "An error occurred while creating the new user.",
				Instance: exceptions.RFC500,
			},
		}
	}

	createUserErr := c.UserRepository.CreateUser(*newUser)
	if createUserErr != nil {
		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating new user",
				Status:   500,
				Detail:   "An error occurred while saving the new user to the database.",
				Instance: exceptions.RFC500,
			},
		}
	}

	return presenters.SuccessOutputDTO{
		SuccessMessage: "User " + newUser.Name + " created successfully",
		ContentMessage: "Welcome, " + newUser.Name + "!",
	}, nil
}
