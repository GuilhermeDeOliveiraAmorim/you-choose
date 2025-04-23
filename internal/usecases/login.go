package usecases

import (
	"context"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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
	util.NewLogger(util.Logger{
		Context: ctx,
		TypeLog: util.LoggerTypes.INFO,
		Layer:   util.LoggerLayers.USECASES,
		Code:    exceptions.RFC200_CODE,
		From:    "LoginUseCase",
		Message: "starting login process",
	})

	email, hashEmailWithHMACErr := util.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		util.NewLogger(util.Logger{
			Context: ctx,
			TypeLog: util.LoggerTypes.ERROR,
			Layer:   util.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "LoginUseCase",
			Message: "error hashing email: " + input.Email,
		})

		return LoginOutputDto{}, hashEmailWithHMACErr
	}

	user, getUserByEmailErr := c.UserRepository.GetUserByEmail(email)
	if getUserByEmailErr != nil {
		if strings.Compare(getUserByEmailErr.Error(), "user not found") == 0 {
			util.NewLogger(util.Logger{
				Context: ctx,
				TypeLog: util.LoggerTypes.INFO,
				Layer:   util.LoggerLayers.USECASES,
				Code:    exceptions.RFC400_CODE,
				From:    "LoginUseCase",
				Message: "user not found: " + input.Email,
			})

			return LoginOutputDto{}, []exceptions.ProblemDetails{
				exceptions.NewProblemDetails(exceptions.Forbidden, util.GetErrorMessage("LoginUseCase", "UserNotFound")),
			}
		}

		util.NewLogger(util.Logger{
			Context: ctx,
			TypeLog: util.LoggerTypes.ERROR,
			Layer:   util.LoggerLayers.USECASES,
			Code:    exceptions.RFC500_CODE,
			From:    "LoginUseCase",
			Message: "error retrieving user: " + getUserByEmailErr.Error(),
		})

		return LoginOutputDto{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(exceptions.InternalServerError, util.GetErrorMessage("LoginUseCase", "ErrorGettingUser")),
		}
	} else if !user.Active {
		util.NewLogger(util.Logger{
			Context: ctx,
			TypeLog: util.LoggerTypes.INFO,
			Layer:   util.LoggerLayers.USECASES,
			Code:    exceptions.RFC400_CODE,
			From:    "LoginUseCase",
			Message: "user is inactive: " + user.ID,
		})

		return LoginOutputDto{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(exceptions.Forbidden, util.GetErrorMessage("LoginUseCase", "UserNotActive")),
		}
	}

	if !user.Login.DecryptPassword(input.Password) {
		util.NewLogger(util.Logger{
			Context: ctx,
			TypeLog: util.LoggerTypes.INFO,
			Layer:   util.LoggerLayers.USECASES,
			Code:    exceptions.RFC400_CODE,
			From:    "LoginUseCase",
			Message: "invalid credentials: " + input.Email,
		})

		return LoginOutputDto{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("LoginUseCase", "InvalidCredentials")),
		}
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
		util.NewLogger(util.Logger{
			Context: ctx,
			TypeLog: util.LoggerTypes.ERROR,
			Layer:   util.LoggerLayers.USECASES,
			Code:    exceptions.RFC400_CODE,
			From:    "LoginUseCase",
			Message: "error signing JWT: " + err.Error(),
		})

		return LoginOutputDto{}, []exceptions.ProblemDetails{
			exceptions.NewProblemDetails(exceptions.InternalServerError, util.GetErrorMessage("LoginUseCase", "JWTError")),
		}
	}

	util.NewLogger(util.Logger{
		Context: ctx,
		TypeLog: util.LoggerTypes.INFO,
		Layer:   util.LoggerLayers.USECASES,
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
