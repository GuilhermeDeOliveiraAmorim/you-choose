package gorm_implementation

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type VoteRepository struct {
	gorm *gorm.DB
}

func NewVoteRepository(gorm *gorm.DB) *VoteRepository {
	return &VoteRepository{
		gorm: gorm,
	}
}

func (c *VoteRepository) CreateVote(movie entities.Vote) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Votes{
		ID:            movie.ID,
		Active:        movie.Active,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
		DeactivatedAt: movie.DeactivatedAt,
		ListID:        movie.ListID,
		FirstMovieID:  movie.FirstMovieID,
		SecondMovieID: movie.SecondMovieID,
		WinnerID:      movie.WinnerID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
