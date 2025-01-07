package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type PaginatedUsers struct {
	Users      []entities.User
	TotalCount int64
}

type UserRepositoryInterface interface {
	CreateUser(user entities.User) error
	ThisUserEmailExists(userEmail string) (bool, error)
	ThisUserNameExists(userName string) (bool, error)
	GetUserByEmail(email string) (entities.User, error)
}
