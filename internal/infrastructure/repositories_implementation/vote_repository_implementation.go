package repositories_implementation

import (
	"errors"

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

func (c *VoteRepository) CreateVote(vote entities.Vote) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Votes{
		ID:            vote.ID,
		Active:        vote.Active,
		CreatedAt:     vote.CreatedAt,
		DeactivatedAt: vote.DeactivatedAt,
		UserID:        vote.UserID,
		CombinationID: vote.CombinationID,
		WinnerID:      vote.WinnerID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *VoteRepository) GetVotesByUserIDAndListID(userID, listID string) ([]entities.Vote, error) {
	var votesModel []Votes

	result := c.gorm.
		Model(&Votes{}).
		Joins("JOIN combinations ON votes.combination_id = combinations.id").
		Where("combinations.list_id = ? AND votes.user_id = ? AND votes.active = ?", listID, userID, true).
		Find(&votesModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("votes not found")
		}
		return nil, errors.New(result.Error.Error())
	}

	var votes []entities.Vote
	for _, voteModel := range votesModel {
		votes = append(votes, *voteModel.ToEntity())
	}

	return votes, nil
}

func (c *VoteRepository) VoteAlreadyRegistered(userID, combinationID string) (bool, error) {
	var count int64

	result := c.gorm.Model(&Votes{}).Where("user_id =? AND combination_id =?", userID, combinationID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
