package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type UserFactory struct {
	CreateUser *usecases.CreateUserUseCase
	Login      *usecases.LoginUseCase
}

func NewUserFactory(input util.ImputFactory) *UserFactory {
	userRepository := repositories_implementation.NewUserRepository(input.DB)

	createUser := usecases.NewCreateUserUseCase(userRepository)
	login := usecases.NewLoginUseCase(userRepository)

	return &UserFactory{
		CreateUser: createUser,
		Login:      login,
	}
}
