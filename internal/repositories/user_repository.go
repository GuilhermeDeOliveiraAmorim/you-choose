package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type UserRepository interface {
	CreateUser(user entities.User) error
	GetUser(userID string) (entities.User, error)
	ThisUserEmailExists(userEmail string) (bool, error)
	ThisUserNameExists(userName string) (bool, error)
	GetUserByEmail(email string) (entities.User, error)
}
