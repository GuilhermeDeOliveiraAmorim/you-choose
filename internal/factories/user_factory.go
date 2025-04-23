package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	repositories_implementation "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type UserFactory struct {
	CreateUser *usecases.CreateUserUseCase
	Login      *usecases.LoginUseCase
}

func NewUserFactory(input database.StorageInput) *UserFactory {
	userRepository := repositories_implementation.NewUserRepository(input.DB)

	createUser := usecases.NewCreateUserUseCase(userRepository)
	login := usecases.NewLoginUseCase(userRepository)

	return &UserFactory{
		CreateUser: createUser,
		Login:      login,
	}
}
