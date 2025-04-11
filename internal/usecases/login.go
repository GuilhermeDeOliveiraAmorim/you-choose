package usecases

import (
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
				util.NewProblemDetails(util.Forbidden, util.GetErrorMessage("LoginUseCase", "UserNotFound")),
			}
		}

		return LoginOutputDto{}, []util.ProblemDetails{
			util.NewProblemDetails(util.InternalServerError, util.GetErrorMessage("LoginUseCase", "ErrorGettingUser")),
		}
	} else if !user.Active {
		return LoginOutputDto{}, []util.ProblemDetails{
			util.NewProblemDetails(util.Forbidden, util.GetErrorMessage("LoginUseCase", "UserNotActive")),
		}
	}

	if !user.Login.DecryptPassword(input.Password) {
		return LoginOutputDto{}, []util.ProblemDetails{
			util.NewProblemDetails(util.Unauthorized, util.GetErrorMessage("LoginUseCase", "InvalidCredentials")),
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
			util.NewProblemDetails(util.InternalServerError, util.GetErrorMessage("LoginUseCase", "JWTError")),
		}
	}

	return LoginOutputDto{
		Name:           user.Name,
		AccessToken:    tokenString,
		SuccessMessage: "Logged in successfully",
		ContentMessage: "Welcome, " + user.Name + "!",
	}, nil
}
