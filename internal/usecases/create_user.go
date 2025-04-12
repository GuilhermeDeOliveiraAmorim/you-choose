package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type CreateUserInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutputDto struct {
	Name           string `json:"name"`
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
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

func (c *CreateUserUseCase) Execute(input CreateUserInputDto) (CreateUserOutputDto, []util.ProblemDetails) {
	email, hashEmailWithHMACErr := util.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error hashing email",
				Status:   500,
				Detail:   "An error occurred while hashing the email with HMAC.",
				Instance: util.RFC500,
			},
		}
	}

	userEmailExists, userEmailExistsErr := c.UserRepository.ThisUserEmailExists(email)
	if userEmailExists {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Conflict",
				Title:    "Email already exists",
				Status:   409,
				Detail:   "The provided email is already registered.",
				Instance: util.RFC409,
			},
		}
	} else if strings.Compare(userEmailExistsErr.Error(), "not found") != 0 {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error checking user email existence",
				Status:   500,
				Detail:   "An error occurred while checking if the email already exists.",
				Instance: util.RFC500,
			},
		}
	}

	userNameExists, userNameExistsErr := c.UserRepository.ThisUserNameExists(input.Name)
	if userNameExists {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Conflict",
				Title:    "Username already exists",
				Status:   409,
				Detail:   "The provided username is already taken.",
				Instance: util.RFC409,
			},
		}
	} else if strings.Compare(userNameExistsErr.Error(), "not found") != 0 {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error checking user name existence",
				Status:   500,
				Detail:   "An error occurred while checking if the username already exists.",
				Instance: util.RFC500,
			},
		}
	}

	newLogin, newLoginErr := entities.NewLogin(input.Email, input.Password)
	if newLoginErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Bad Request",
				Title:    "Error creating login",
				Status:   400,
				Detail:   "An error occurred while creating the user login information.",
				Instance: util.RFC400,
			},
		}
	}

	encryptEmailErr := newLogin.EncryptEmail()
	if encryptEmailErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error encrypting email",
				Status:   500,
				Detail:   "An error occurred while encrypting the email address.",
				Instance: util.RFC500,
			},
		}
	}

	encryptPasswordErr := newLogin.EncryptPassword()
	if encryptPasswordErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error encrypting password",
				Status:   500,
				Detail:   "An error occurred while encrypting the password.",
				Instance: util.RFC500,
			},
		}
	}

	newUser, newUserErr := entities.NewUser(input.Name, *newLogin)
	if newUserErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating user",
				Status:   500,
				Detail:   "An error occurred while creating the new user.",
				Instance: util.RFC500,
			},
		}
	}

	createUserErr := c.UserRepository.CreateUser(*newUser)
	if createUserErr != nil {
		return CreateUserOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating new user",
				Status:   500,
				Detail:   "An error occurred while saving the new user to the database.",
				Instance: util.RFC500,
			},
		}
	}

	return CreateUserOutputDto{
		Name:           newUser.Name,
		SuccessMessage: "User created successfully",
		ContentMessage: "Welcome, " + newUser.Name + "!",
	}, nil
}
