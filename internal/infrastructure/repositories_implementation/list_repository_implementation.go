package repositories_implementation

import (
	"errors"
	"time"

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
		if err := tx.Exec("INSERT INTO list_movies (list_id, movie_id, created_at) VALUES (?, ?, ?)", list.ID, movie.ID, time.Now()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, combination := range list.Combinations {
		if err := tx.Exec("INSERT INTO combinations (id, list_id, first_movie_id, second_movie_id) VALUES (?, ?, ?, ?)", combination.ID, list.ID, combination.FirstMovieID, combination.SecondMovieID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *ListRepository) GetListByUserID(listID string) (entities.List, error) {
	var listModel Lists

	resultListModel := c.gorm.Model(&Lists{}).Where("id = ? AND active = ?", listID, true).First(&listModel)
	if resultListModel.Error != nil {
		if errors.Is(resultListModel.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("list not found")
		}
		return entities.List{}, errors.New(resultListModel.Error.Error())
	}

	var moviesModel []Movies

	resultMoviesModel := c.gorm.Table("movies").
		Select("movies.*").
		Joins("JOIN list_movies ON list_movies.movie_id = movies.id").
		Where("list_movies.list_id = ?", listID).
		Find(&moviesModel)
	if resultMoviesModel.Error != nil {
		if errors.Is(resultMoviesModel.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("list not found")
		}
		return entities.List{}, errors.New(resultMoviesModel.Error.Error())
	}

	var movies []entities.Movie

	for _, movie := range moviesModel {
		movies = append(movies, *movie.ToEntity())
	}

	var combinationsModel []Combinations

	result := c.gorm.Table("combinations").
		Select("combinations.*").
		Where("list_id = ?", listID).
		Find(&combinationsModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("list not found")
		}
		return entities.List{}, errors.New(result.Error.Error())
	}

	var combinations []entities.Combination

	for _, combination := range combinationsModel {
		combinations = append(combinations, *combination.ToEntity())
	}

	return *listModel.ToEntity(movies, combinations), nil
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

	for _, movie := range list.Movies {
		if err := tx.Exec("INSERT INTO list_movies (list_id, movie_id, created_at) VALUES (?, ?, ?)", list.ID, movie.ID, time.Now()).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, combination := range list.Combinations {
		if err := tx.Exec("INSERT INTO combinations (id, list_id, first_movie_id, second_movie_id) VALUES (?, ?, ?, ?)", combination.ID, list.ID, combination.FirstMovieID, combination.SecondMovieID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
