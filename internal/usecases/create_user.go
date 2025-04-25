package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
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
	problems := []exceptions.ProblemDetails{}

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
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "EmailHMACError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error hashing email: " + input.Email,
			Error:    hashEmailWithHMACErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	userEmailExists, userEmailExistsErr := c.UserRepository.ThisUserEmailExists(email)
	if userEmailExists {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Conflict, language.GetErrorMessage("CreateUserUseCase", "EmailAlreadyExists")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC409_CODE,
			From:     "CreateUserUseCase",
			Message:  "email already exists: " + input.Email,
			Error:    userEmailExistsErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	} else if strings.Compare(userEmailExistsErr.Error(), "email not found") != 0 {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "EmailCheckError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error checking email existence: " + input.Email,
			Error:    userEmailExistsErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	userNameExists, userNameExistsErr := c.UserRepository.ThisUserNameExists(input.Name)
	if userNameExists {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Conflict, language.GetErrorMessage("CreateUserUseCase", "UsernameAlreadyExists")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC409_CODE,
			From:     "CreateUserUseCase",
			Message:  "username already exists: " + input.Name,
			Error:    userNameExistsErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	} else if strings.Compare(userNameExistsErr.Error(), "username not found") != 0 {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "UsernameCheckError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error checking username existence: " + input.Name,
			Error:    userNameExistsErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	newLogin, newLoginErr := entities.NewLogin(input.Email, input.Password)
	if newLoginErr != nil {
		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC400_CODE,
			From:     "CreateUserUseCase",
			Message:  "error creating new login: " + input.Email,
			Error:    errors.New("error creating new login"),
			Problems: newLoginErr,
		})

		return presenters.SuccessOutputDTO{}, newLoginErr
	}

	encryptEmailErr := newLogin.HashEmail()
	if encryptEmailErr != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "HashEmailError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error encrypting email: " + input.Email,
			Error:    encryptEmailErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
	}

	encryptPasswordErr := newLogin.HashPassword()
	if encryptPasswordErr != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "HashPasswordError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error encrypting password: " + input.Password,
			Error:    encryptPasswordErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
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
			Error:    errors.New("error creating new user"),
			Problems: newUserErr,
		})

		return presenters.SuccessOutputDTO{}, newUserErr
	}

	createUserErr := c.UserRepository.CreateUser(*newUser)
	if createUserErr != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("CreateUserUseCase", "SaveUserError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "CreateUserUseCase",
			Message:  "error creating new user in database: " + input.Name,
			Error:    createUserErr,
			Problems: problems,
		})

		return presenters.SuccessOutputDTO{}, problems
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
