package usecases

import (
	"fmt"
	"strings"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
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

func (c *LoginUseCase) Execute(input LoginInputDto) (LoginOutputDto, []util.ProblemDetails) {
	email, hashEmailWithHMACErr := util.HashEmailWithHMAC(input.Email)
	if hashEmailWithHMACErr != nil {
		return LoginOutputDto{}, hashEmailWithHMACErr
	}

	user, getUserByEmailErr := c.UserRepository.GetUserByEmail(email)
	if getUserByEmailErr != nil {
		if strings.Compare(getUserByEmailErr.Error(), "user not found") == 0 {
			return LoginOutputDto{}, []util.ProblemDetails{
				{
					Type:     "Forbidden",
					Title:    util.GetErrorMessage("LoginUseCase", "UserNotFound", "Title"),
					Status:   403,
					Detail:   util.GetErrorMessage("LoginUseCase", "UserNotFound", "Detail"),
					Instance: util.RFC403,
				},
			}
		}

		return LoginOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    util.GetErrorMessage("LoginUseCase", "ErrorGettingUser", "Title"),
				Status:   500,
				Detail:   util.GetErrorMessage("LoginUseCase", "ErrorGettingUser", "Detail"),
				Instance: util.RFC500,
			},
		}
	} else if !user.Active {
		return LoginOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    util.GetErrorMessage("LoginUseCase", "UserNotActive", "Title"),
				Status:   403,
				Detail:   util.GetErrorMessage("LoginUseCase", "UserNotActive", "Detail"),
				Instance: util.RFC403,
			},
		}
	}

	if !user.Login.DecryptPassword(input.Password) {
		return LoginOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Unauthorized",
				Title:    util.GetErrorMessage("LoginUseCase", "InvalidCredentials", "Title"),
				Status:   401,
				Detail:   util.GetErrorMessage("LoginUseCase", "InvalidCredentials", "Detail"),
				Instance: util.RFC401,
			},
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
		return LoginOutputDto{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    util.GetErrorMessage("LoginUseCase", "JWTError", "Title"),
				Status:   500,
				Detail:   err.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return LoginOutputDto{
		Name:           user.Name,
		AccessToken:    tokenString,
		SuccessMessage: "Logged in successfully",
		ContentMessage: "Welcome, " + user.Name + "!",
	}, nil
}
