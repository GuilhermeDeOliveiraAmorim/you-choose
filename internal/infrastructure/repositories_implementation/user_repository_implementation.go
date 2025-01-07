package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) *UserRepository {
	return &UserRepository{
		gorm: gorm,
	}
}

func (u *UserRepository) CreateUser(user entities.User) error {
	tx := u.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Users{
		ID:            user.ID,
		Active:        user.Active,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DeactivatedAt: user.DeactivatedAt,
		Name:          user.Name,
		Email:         user.Login.Email,
		Password:      user.Login.Password,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *UserRepository) GetUserByEmail(userEmail string) (entities.User, error) {
	var userModel Users

	result := u.gorm.Model(&Users{}).Where("email = ?", userEmail).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, errors.New(result.Error.Error())
	}

	return *userModel.ToEntity(), nil
}

func (u *UserRepository) ThisUserEmailExists(userEmail string) (bool, error) {
	var userModel Users

	result := u.gorm.Model(&Users{}).Where("email = ?", userEmail).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("not found")
		}
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}

func (u *UserRepository) ThisUserNameExists(userName string) (bool, error) {
	var userModel Users

	result := u.gorm.Model(&Users{}).Where("name = ?", userName).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("not found")
		}
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}
