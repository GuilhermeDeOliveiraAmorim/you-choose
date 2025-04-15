package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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

	if err := tx.Create(&models.Users{
		ID:            user.ID,
		Active:        user.Active,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DeactivatedAt: user.DeactivatedAt,
		Name:          user.Name,
		Email:         user.Login.Email,
		Password:      user.Login.Password,
		IsAdmin:       user.IsAdmin,
	}).Error; err != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateUser",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *UserRepository) GetUserByEmail(userEmail string) (entities.User, error) {
	var userModel models.Users

	result := u.gorm.Model(&models.Users{}).Where("email = ?", userEmail).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetUserByEmail",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		return entities.User{}, result.Error
	}

	return *userModel.ToEntity(), nil
}

func (u *UserRepository) ThisUserEmailExists(userEmail string) (bool, error) {
	var userModel models.Users

	result := u.gorm.Model(&models.Users{}).Where("email = ?", userEmail).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("email not found")
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisUserEmailExists",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		return false, result.Error
	}

	return true, nil
}

func (u *UserRepository) ThisUserNameExists(userName string) (bool, error) {
	var userModel models.Users

	result := u.gorm.Model(&models.Users{}).Where("name = ?", userName).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, errors.New("username not found")
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisUserNameExists",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		return false, result.Error
	}

	return true, nil
}

func (u *UserRepository) GetUser(userID string) (entities.User, error) {
	var userModel models.Users

	result := u.gorm.Model(&models.Users{}).Where("id = ?", userID).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetUser",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
		})
		return entities.User{}, result.Error
	}

	user := entities.User{
		SharedEntity: entities.SharedEntity{
			ID:            userModel.ID,
			Active:        userModel.Active,
			CreatedAt:     userModel.CreatedAt,
			UpdatedAt:     userModel.UpdatedAt,
			DeactivatedAt: userModel.DeactivatedAt,
		},
		Name:    userModel.Name,
		IsAdmin: userModel.IsAdmin,
	}

	return user, nil
}
