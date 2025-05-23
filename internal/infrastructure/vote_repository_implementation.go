package repositories_implementation

import (
	"errors"
	"sort"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
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

	if err := tx.Create(&models.Votes{
		ID:            vote.ID,
		Active:        vote.Active,
		CreatedAt:     vote.CreatedAt,
		DeactivatedAt: vote.DeactivatedAt,
		UserID:        vote.UserID,
		CombinationID: vote.CombinationID,
		WinnerID:      vote.WinnerID,
	}).Error; err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "CreateVote",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *VoteRepository) GetVotesByUserIDAndListID(userID, listID string) ([]entities.Vote, error) {
	var votesModel []models.Votes

	result := c.gorm.
		Model(&models.Votes{}).
		Joins("JOIN combinations ON votes.combination_id = combinations.id").
		Where("combinations.list_id = ? AND votes.user_id = ? AND votes.active = ?", listID, userID, true).
		Find(&votesModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("votes not found")
		}
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetVotesByUserIDAndListID",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, result.Error
	}

	var votes []entities.Vote
	for _, voteModel := range votesModel {
		votes = append(votes, *voteModel.ToEntity())
	}

	return votes, nil
}

func (c *VoteRepository) VoteAlreadyRegistered(userID, combinationID string) (bool, error) {
	var count int64

	result := c.gorm.Model(&models.Votes{}).Where("user_id =? AND combination_id =?", userID, combinationID).Count(&count)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "VoteAlreadyRegistered",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return false, result.Error
	}

	return count > 0, nil
}

func (c *VoteRepository) GetNumberOfVotesByListID(listID string) (int, error) {
	var count int64

	result := c.gorm.Model(&models.Votes{}).Where("combination_id IN (SELECT id FROM combinations WHERE list_id =? AND active =?)", listID, true).Count(&count)
	if result.Error != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: result.Error.Error(),
			From:    "GetNumberOfVotesByListID",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return 0, result.Error
	}

	return int(count), nil
}

func (c *VoteRepository) RankItemsByVotes(listID, listType string) ([]interface{}, error) {
	var combinations []models.Combinations
	if err := c.gorm.Where("list_id = ?", listID).Find(&combinations).Error; err != nil {
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: err.Error(),
			From:    "RankItemsByVotes",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, err
	}

	voteCounts := make(map[string]int)
	for _, combination := range combinations {
		var votes []models.Votes
		if err := c.gorm.Where("combination_id = ?", combination.ID).Find(&votes).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "RankItemsByVotes",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			return nil, err
		}

		for _, vote := range votes {
			voteCounts[vote.WinnerID]++
		}
	}

	switch listType {
	case entities.MOVIE_TYPE:
		return c.RankMoviesByVotes(voteCounts)
	case entities.BRAND_TYPE:
		return c.RankBrandsByVotes(voteCounts)
	default:
		logging.NewLogger(logging.Logger{
			Code:    exceptions.RFC500_CODE,
			Message: "Invalid list type",
			From:    "RankItemsByVotes",
			Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
			TypeLog: logging.LoggerTypes.ERROR,
		})
		return nil, errors.New("Invalid list type")
	}
}

func (c *VoteRepository) RankMoviesByVotes(voteCounts map[string]int) ([]interface{}, error) {
	var movies []models.Movies

	for movieID, count := range voteCounts {
		var movie models.Movies
		if err := c.gorm.First(&movie, "id = ?", movieID).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "RankMoviesByVotes",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			return nil, err
		}
		movie.VotesCount = count
		movies = append(movies, movie)
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].VotesCount > movies[j].VotesCount
	})

	var result []interface{}
	for _, movie := range movies {
		result = append(result, *movie.ToEntity())
	}

	return result, nil
}

func (c *VoteRepository) RankBrandsByVotes(voteCounts map[string]int) ([]interface{}, error) {
	var brands []models.Brands

	for brandID, count := range voteCounts {
		var brand models.Brands
		if err := c.gorm.First(&brand, "id = ?", brandID).Error; err != nil {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC500_CODE,
				Message: err.Error(),
				From:    "RankBrandsByVotes",
				Layer:   logging.LoggerLayers.INFRASTRUCTURE_REPOSITORIES_IMPLEMENTATION,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			return nil, err
		}
		brand.VotesCount = count
		brands = append(brands, brand)
	}

	sort.Slice(brands, func(i, j int) bool {
		return brands[i].VotesCount > brands[j].VotesCount
	})

	var result []interface{}
	for _, brand := range brands {
		result = append(result, *brand.ToEntity())
	}

	return result, nil
}
