package gorm_implementation

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
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

	if err := tx.Create(&Movies{
		ID:            movie.ID,
		Active:        movie.Active,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
		DeactivatedAt: movie.DeactivatedAt,
		Name:          movie.Name,
		Year:          movie.Year,
		Poster:        movie.Poster,
		ExternalID:    movie.ExternalID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
