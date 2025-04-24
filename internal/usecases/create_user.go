package usecases

import (
	"context"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
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

func (c *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInputDto) (presenters.SuccessOutputDTO, []exceptions.ProblemDetails) {
	logging.NewLogger(logging.Logger{
		Context: ctx,
		TypeLog: logging.LoggerTypes.INFO,
		Layer:   logging.LoggerLayers.USECASES,
		Code:    exceptions.RFC200_CODE,
		From:    "CreateUserUseCase",
		Message: "starting create user process",
	})

	email, hashEmailWithHMACErr := c.UserRepository.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error hashing email: " + input.Email,
			Error:   hashEmailWithHMACErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC409_CODE,
			From:    "CreateUserUseCase",
			Message: "email already exists: " + input.Email,
			Error:   userEmailExistsErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error checking email existence: " + input.Email,
			Error:   userEmailExistsErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC409_CODE,
			From:    "CreateUserUseCase",
			Message: "username already exists: " + input.Name,
			Error:   userNameExistsErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error checking username existence: " + input.Name,
			Error:   userNameExistsErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error creating new login: " + input.Email,
			Problems: newLoginErr,
		})

		return presenters.SuccessOutputDTO{}, newLoginErr
	}

	encryptEmailErr := newLogin.EncryptEmail()
	if encryptEmailErr != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error encrypting email: " + input.Email,
			Error:   encryptEmailErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error encrypting password: " + input.Password,
			Error:   encryptPasswordErr,
		})

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
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "CreateUserUseCase",
			Message:  "error creating new user: " + input.Name,
			Problems: newUserErr,
		})

		return presenters.SuccessOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Bad Request",
				Title:    "Error creating user",
				Status:   400,
				Detail:   "An error occurred while creating the new user.",
				Instance: exceptions.RFC400,
			},
		}
	}

	createUserErr := c.UserRepository.CreateUser(*newUser)
	if createUserErr != nil {
		logging.NewLogger(logging.Logger{
			Context: ctx,
			TypeLog: logging.LoggerTypes.ERROR,
			Layer:   logging.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "CreateUserUseCase",
			Message: "error creating new user in database: " + input.Name,
			Error:   createUserErr,
		})

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

	logging.NewLogger(logging.Logger{
		Context: ctx,
		TypeLog: logging.LoggerTypes.INFO,
		Layer:   logging.LoggerLayers.USECASES,
		Code:    exceptions.RFC201_CODE,
		From:    "CreateUserUseCase",
		Message: "user created successfully: " + input.Name,
	})

	return presenters.SuccessOutputDTO{
		SuccessMessage: "User " + newUser.Name + " created successfully",
		ContentMessage: "Welcome, " + newUser.Name + "!",
	}, nil
}
