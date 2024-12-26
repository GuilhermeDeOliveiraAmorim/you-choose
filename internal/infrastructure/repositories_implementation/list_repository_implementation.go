package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type ListRepository struct {
	gorm *gorm.DB
}

func NewListRepository(gorm *gorm.DB) *ListRepository {
	return &ListRepository{
		gorm: gorm,
	}
}

func (c *ListRepository) CreateList(list entities.List) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Lists{
		ID:            list.ID,
		Active:        list.Active,
		CreatedAt:     list.CreatedAt,
		UpdatedAt:     list.UpdatedAt,
		DeactivatedAt: list.DeactivatedAt,
		Name:          list.Name,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, movie := range list.Movies {
		if err := tx.Exec("INSERT INTO list_tags (lists_id, tags_id) VALUES (?, ?)", list.ID, movie.ID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *ListRepository) GetListByID(listID string) (entities.List, error) {
	var listModel Lists

	result := c.gorm.Model(&Lists{}).Where("id = ? AND active = ?", listID, true).First(&listModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("list not found")
		}
		return entities.List{}, errors.New(result.Error.Error())
	}

	var moviesModel []Movies

	result.Model(&Movies{}).Joins("JOIN list_tags ON list_tags.tags_id = movies.id").Where("list_tags.lists_id =?", listID).Find(&moviesModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("list not found")
		}
		return entities.List{}, errors.New(result.Error.Error())
	}

	var movies []entities.Movie

	for _, movie := range moviesModel {
		movies = append(movies, *movie.ToEntity())
	}

	return *listModel.ToEntity(movies), nil
}

func (c *ListRepository) ThisListExistByName(listName string) (bool, error) {
	var count int64

	result := c.gorm.Model(&Lists{}).Where("name =? AND active =?", listName, true).Count(&count)
	if result.Error != nil {
		return false, errors.New(result.Error.Error())
	}

	return count > 0, nil
}

func (c *ListRepository) ThisListExistByID(listID string) (bool, error) {
	var count int64

	result := c.gorm.Model(&Lists{}).Where("id =? AND active =?", listID, true).Count(&count)
	if result.Error != nil {
		return false, errors.New(result.Error.Error())
	}

	return count > 0, nil
}

func (c *ListRepository) AddMovies(list entities.List) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	result := tx.Model(&Lists{}).Where("id = ? AND active = ?", list.ID, true).Updates(Lists{
		UpdatedAt: list.UpdatedAt,
	})

	if result.Error != nil {
		tx.Rollback()
		return errors.New(result.Error.Error())
	}

	for _, movie := range list.Movies {
		if err := tx.Exec("INSERT INTO list_tags (lists_id, tags_id) VALUES (?, ?)", list.ID, movie.ID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
