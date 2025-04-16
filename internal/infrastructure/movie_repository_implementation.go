package repositories_implementation

import (
	"errors"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"gorm.io/gorm"
)

type MovieRepository struct {
	gorm *gorm.DB
}

func NewMovieRepository(gorm *gorm.DB) *MovieRepository {
	return &MovieRepository{
		gorm: gorm,
	}
}

func (c *MovieRepository) CreateMovie(movie entities.Movie) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&models.Movies{
		ID:            movie.ID,
		Active:        movie.Active,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
		DeactivatedAt: movie.DeactivatedAt,
		Name:          movie.Name,
		Year:          movie.Year,
		Poster:        movie.Poster,
		ExternalID:    movie.ExternalID,
		VotesCount:    movie.VotesCount,
	}).Error; err != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateMovie",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *MovieRepository) GetMovieByID(movieID string) (entities.Movie, error) {
	var movieModel models.Movies

	result := c.gorm.Model(&models.Movies{}).Where("id =? AND active =?", movieID, true).First(&movieModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Movie{}, errors.New("movie not found")
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetMovieByID",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		return entities.Movie{}, result.Error
	}

	return *movieModel.ToEntity(), nil
}

func (c *MovieRepository) ThisMovieExist(movieExternalID string) (bool, error) {
	var movieModel models.Movies

	result := c.gorm.Model(&models.Movies{}).Where("external_id =?", movieExternalID).First(&movieModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "ThisMovieExist",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		return false, result.Error
	}

	return true, nil
}

func (c *MovieRepository) GetMoviesByIDs(moviesIDs []string) ([]entities.Movie, error) {
	var moviesModel []models.Movies

	result := c.gorm.Model(&models.Movies{}).Where("id IN?", moviesIDs).Find(&moviesModel)
	if result.Error != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetMoviesByIDs",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var movies []entities.Movie
	for _, movieModel := range moviesModel {
		movies = append(movies, *movieModel.ToEntity())
	}

	return movies, nil
}

func (c *MovieRepository) UpdadeMovie(movie entities.Movie) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Model(&models.Movies{}).Where("id =?", movie.ID).Updates(models.Movies{
		Active:        movie.Active,
		Name:          movie.Name,
		Year:          movie.Year,
		Poster:        movie.Poster,
		VotesCount:    movie.VotesCount,
		DeactivatedAt: movie.DeactivatedAt,
		UpdatedAt:     movie.UpdatedAt,
		ExternalID:    movie.ExternalID,
	}).Error; err != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: err.Error(),
			From:    "UpdadeMovie",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *MovieRepository) GetMovies() ([]entities.Movie, error) {
	var moviesModel []models.Movies

	result := c.gorm.Model(&models.Movies{}).Where("active =?", true).Find(&moviesModel)
	if result.Error != nil {
		util.NewLogger(util.Logger{
			Code:    util.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetMovies",
			Layer:   util.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: util.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var movies []entities.Movie
	for _, movieModel := range moviesModel {
		movies = append(movies, *movieModel.ToEntity())
	}

	return movies, nil
}
