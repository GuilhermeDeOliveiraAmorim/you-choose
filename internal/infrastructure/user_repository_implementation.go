package repositories_implementation

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
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
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateUser",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
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
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetUserByEmail",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
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
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisUserEmailExists",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
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
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisUserNameExists",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
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
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetUser",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
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

func (u *UserRepository) HashEmailWithHMAC(email string) (string, error) {
	key := []byte(config.SECRETS_VAR.JWT_SECRET)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(email))

	return hex.EncodeToString(h.Sum(nil)), nil
}
