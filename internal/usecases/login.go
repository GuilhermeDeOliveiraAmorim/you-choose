package usecases

import (
	"context"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/dgrijalva/jwt-go"
)

type LoginInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDto struct {
	Name           string `json:"name"`
	AccessToken    string `json:"access_token"`
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type LoginUseCase struct {
	UserRepository repositories.UserRepository
}

func NewLoginUseCase(
	UserRepository repositories.UserRepository,
) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: UserRepository,
	}
}

func (c *LoginUseCase) Execute(ctx context.Context, input LoginInputDto) (LoginOutputDto, []exceptions.ProblemDetails) {
	problems := []exceptions.ProblemDetails{}

	logging.NewLogger(logging.Logger{
		Context: ctx,
		TypeLog: logging.LoggerTypes.INFO,
		Layer:   logging.LoggerLayers.USECASES,
		Code:    exceptions.RFC200_CODE,
		From:    "LoginUseCase",
		Message: "starting login process",
	})

	email, hashEmailWithHMACErr := c.UserRepository.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("LoginUseCase", "HMAC")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "LoginUseCase",
			Message:  "error hashing email: " + input.Email,
			Error:    hashEmailWithHMACErr,
			Problems: problems,
		})

		return LoginOutputDto{}, problems
	}

	user, getUserByEmailErr := c.UserRepository.GetUserByEmail(email)
	if getUserByEmailErr != nil {
		if strings.Compare(getUserByEmailErr.Error(), "user not found") == 0 {
			problems = append(problems, exceptions.NewProblemDetails(exceptions.Forbidden, language.GetErrorMessage("LoginUseCase", "UserNotFound")))

			logging.NewLogger(logging.Logger{
				Context:  ctx,
				TypeLog:  logging.LoggerTypes.ERROR,
				Layer:    logging.LoggerLayers.USECASES,
				Code:     exceptions.RFC403_CODE,
				From:     "LoginUseCase",
				Message:  "user not found: " + input.Email,
				Error:    getUserByEmailErr,
				Problems: problems,
			})

			return LoginOutputDto{}, problems
		}

		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("LoginUseCase", "ErrorGettingUser")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "LoginUseCase",
			Message:  "error retrieving user: " + getUserByEmailErr.Error(),
			Error:    getUserByEmailErr,
			Problems: problems,
		})

		return LoginOutputDto{}, problems
	}

	if !user.Login.VerifyPassword(input.Password) {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("LoginUseCase", "InvalidCredentials")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC401_CODE,
			From:     "LoginUseCase",
			Message:  "invalid credentials: " + input.Email,
			Error:    getUserByEmailErr,
			Problems: problems,
		})

		return LoginOutputDto{}, problems
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.SECRETS_VAR.JWT_SECRET)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		problems = append(problems, exceptions.NewProblemDetails(exceptions.InternalServerError, language.GetErrorMessage("LoginUseCase", "JWTError")))

		logging.NewLogger(logging.Logger{
			Context:  ctx,
			TypeLog:  logging.LoggerTypes.ERROR,
			Layer:    logging.LoggerLayers.USECASES,
			Code:     exceptions.RFC500_CODE,
			From:     "LoginUseCase",
			Message:  "error signing JWT: " + err.Error(),
			Error:    err,
			Problems: problems,
		})

		return LoginOutputDto{}, problems
	}

	logging.NewLogger(logging.Logger{
		Context: ctx,
		TypeLog: logging.LoggerTypes.INFO,
		Layer:   logging.LoggerLayers.USECASES,
		Code:    exceptions.RFC200_CODE,
		From:    "LoginUseCase",
		Message: "login successful: " + user.ID,
	})

	return LoginOutputDto{
		Name:           user.Name,
		AccessToken:    tokenString,
		SuccessMessage: "Logged in successfully",
		ContentMessage: "Welcome, " + user.Name + "!",
	}, nil
}
