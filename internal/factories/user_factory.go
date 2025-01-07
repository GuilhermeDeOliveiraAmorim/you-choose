package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type UserFactory struct {
	CreateUser *usecases.CreateUserUseCase
	Login      *usecases.LoginUseCase
}

func NewUserFactory(db *gorm.DB) *UserFactory {
	userRepository := repositories_implementation.NewUserRepository(db)

	createUser := usecases.NewCreateUserUseCase(userRepository)
	login := usecases.NewLoginUseCase(userRepository)

	return &UserFactory{
		CreateUser: createUser,
		Login:      login,
	}
}
