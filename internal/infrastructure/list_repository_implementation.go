package repositories_implementation

import (
	"errors"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
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

	if err := tx.Create(&models.Lists{
		ID:            list.ID,
		Active:        list.Active,
		CreatedAt:     list.CreatedAt,
		UpdatedAt:     list.UpdatedAt,
		DeactivatedAt: list.DeactivatedAt,
		Name:          list.Name,
		Cover:         list.Cover,
		ListType:      list.ListType,
	}).Error; err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateList 1",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	for _, item := range list.Items {
		switch item := item.(type) {
		case entities.Movie:
			if err := tx.Exec("INSERT INTO list_movies (list_id, movie_id, created_at) VALUES (?, ?, ?)", list.ID, item.ID, time.Now()).Error; err != nil {
				logging.NewLogger(logging.Logger{
					Code:    exceptions.RFC500_CODE,
					Message: err.Error(),
					From:    "CreateList 2",
					Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
					TypeLog: logging.LoggerTypes.ERROR,
				})
				tx.Rollback()
				return err
			}
		case entities.Brand:
			if err := tx.Exec("INSERT INTO list_brands (list_id, brand_id, created_at) VALUES (?, ?,?)", list.ID, item.ID, time.Now()).Error; err != nil {
				logging.NewLogger(logging.Logger{
					Code:    exceptions.RFC500_CODE,
					Message: err.Error(),
					From:    "CreateList 3",
					Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
					TypeLog: logging.LoggerTypes.ERROR,
				})
				tx.Rollback()
				return err
			}
		}
	}

	for _, combination := range list.Combinations {
		if err := tx.Exec("INSERT INTO combinations (id, list_id, first_item_id, second_item_id) VALUES (?, ?, ?, ?)", combination.ID, list.ID, combination.FirstItemID, combination.SecondItemID).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "CreateList 4",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *ListRepository) ThisListExistByName(listName string) (bool, error) {
	var count int64

	result := c.gorm.Model(&models.Lists{}).Where("name =? AND active =?", listName, true).Count(&count)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisListExistByName",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return false, result.Error
	}

	return count > 0, nil
}

func (c *ListRepository) ThisListExistByID(listID string) (bool, error) {
	var count int64

	result := c.gorm.Model(&models.Lists{}).Where("id =? AND active =?", listID, true).Count(&count)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisListExistByID",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return false, result.Error
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

	for _, item := range list.Items {
		switch item := item.(type) {
		case entities.Movie:
			if err := tx.Exec("INSERT INTO list_movies (list_id, movie_id, created_at) VALUES (?, ?, ?)", list.ID, item.ID, time.Now()).Error; err != nil {
				logging.NewLogger(logging.Logger{
					Code:    exceptions.RFC500_CODE,
					Message: err.Error(),
					From:    "AddMovies 1",
					Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
					TypeLog: logging.LoggerTypes.ERROR,
				})
				tx.Rollback()
				return err
			}
		}
	}

	for _, combination := range list.Combinations {
		if err := tx.Exec("INSERT INTO combinations (id, list_id, first_item_id, second_item_id) VALUES (?, ?, ?, ?)", combination.ID, list.ID, combination.FirstItemID, combination.SecondItemID).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "AddMovies 2",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *ListRepository) GetListByID(listID string) (entities.List, error) {
	var listModel models.Lists

	resultListModel := c.gorm.Model(&models.Lists{}).Where("id = ? AND active = ?", listID, true).First(&listModel)
	if resultListModel.Error != nil {
		if errors.Is(resultListModel.Error, gorm.ErrRecordNotFound) {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC200_CODE,
				Message: resultListModel.Error.Error(),
				From:    "GetListByID 0",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.WARNING,
			})

			return entities.List{}, errors.New("list not found")
		}
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: resultListModel.Error.Error(),
			From:    "GetListByID 1",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return entities.List{}, resultListModel.Error
	}

	items, err := c.FetchItemsByListType(listID, listModel.ListType)
	if err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "GetListByID 2",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return entities.List{}, err
	}

	var combinationsModel []models.Combinations
	result := c.gorm.Table("combinations").
		Select("combinations.*").
		Where("list_id = ?", listID).
		Find(&combinationsModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.List{}, errors.New("combinations not found")
		}
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetListByID 3",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return entities.List{}, result.Error
	}

	var combinations []entities.Combination
	for _, combination := range combinationsModel {
		combinations = append(combinations, *combination.ToEntity())
	}

	return *listModel.ToEntity(items, combinations, true), nil
}

func (c *ListRepository) GetLists() ([]entities.List, error) {
	var listsModel []models.Lists

	result := c.gorm.Model(&models.Lists{}).Where("active =?", true).Find(&listsModel)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetLists",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var lists []entities.List

	for _, list := range listsModel {
		lists = append(lists, *list.ToEntity([]interface{}{}, []entities.Combination{}, false))
	}

	return lists, nil
}

func (c *ListRepository) FetchItemsByListType(listID, listType string) ([]interface{}, error) {
	var items []interface{}

	switch listType {
	case entities.MOVIE_TYPE:
		var moviesModel []models.Movies
		resultMoviesModel := c.gorm.Table("movies").
			Select("movies.*").
			Joins("JOIN list_movies ON list_movies.movie_id = movies.id").
			Where("list_movies.list_id = ?", listID).
			Find(&moviesModel)
		if resultMoviesModel.Error != nil {
			if errors.Is(resultMoviesModel.Error, gorm.ErrRecordNotFound) {
				return nil, errors.New("movies not found")
			}
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: resultMoviesModel.Error.Error(),
				From:    "FetchItemsByListType 1",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			return nil, resultMoviesModel.Error
		}

		for _, movie := range moviesModel {
			items = append(items, *movie.ToEntity())
		}

	case entities.BRAND_TYPE:
		var brandsModel []models.Brands
		resultBrandsModel := c.gorm.Table("brands").
			Select("brands.*").
			Joins("JOIN list_brands ON list_brands.brand_id = brands.id").
			Where("list_brands.list_id = ?", listID).
			Find(&brandsModel)
		if resultBrandsModel.Error != nil {
			if errors.Is(resultBrandsModel.Error, gorm.ErrRecordNotFound) {
				return nil, errors.New("brands not found")
			}
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: resultBrandsModel.Error.Error(),
				From:    "FetchItemsByListType 2",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			return nil, resultBrandsModel.Error
		}

		for _, brand := range brandsModel {
			items = append(items, *brand.ToEntity())
		}

	default:
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: errors.New("invalid list type").Error(),
			From:    "FetchItemsByListType 2",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, errors.New("Invalid list type")
	}

	return items, nil
}

func (c *ListRepository) AddBrands(list entities.List) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	for _, item := range list.Items {
		switch item := item.(type) {
		case entities.Brand:
			if err := tx.Exec("INSERT INTO list_brands (list_id, brand_id, created_at) VALUES (?, ?, ?)", list.ID, item.ID, time.Now()).Error; err != nil {
				logging.NewLogger(logging.Logger{
					Code:    exceptions.RFC500_CODE,
					Message: err.Error(),
					From:    "AddBrands 1",
					Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
					TypeLog: logging.LoggerTypes.ERROR,
				})
				tx.Rollback()
				return err
			}
		}
	}

	for _, combination := range list.Combinations {
		if err := tx.Exec("INSERT INTO combinations (id, list_id, first_item_id, second_item_id) VALUES (?, ?, ?, ?)", combination.ID, list.ID, combination.FirstItemID, combination.SecondItemID).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "AddBrands 2",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
