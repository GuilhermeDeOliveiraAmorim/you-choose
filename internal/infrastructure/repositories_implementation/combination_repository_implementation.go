package repositories_implementation

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type CombinationRepository struct {
	gorm *gorm.DB
}

func NewCombinationRepository(gorm *gorm.DB) *CombinationRepository {
	return &CombinationRepository{
		gorm: gorm,
	}
}

func (c *CombinationRepository) GetCombinationsByListID(listID string) ([]entities.Combination, error) {
	var combinationsModel []Combinations

	result := c.gorm.Table("combinations").
		Select("combinations.*").
		Where("list_id = ?", listID).
		Find(&combinationsModel)
	if result.Error != nil {
		return nil, result.Error
	}

	var combinations []entities.Combination

	for _, combination := range combinationsModel {
		combinations = append(combinations, *combination.ToEntity())
	}

	return combinations, nil
}

func (c *CombinationRepository) GetCombinationsAlreadyVoted(listID string) ([]entities.Combination, error) {
	var combinationsModel []Combinations

    result := c.gorm.Table("combinations").
        Select("combinations.*").
        Joins("JOIN votes ON combinations.id = votes.combination_id").
        Where("list_id =?", listID).
        Find(&combinationsModel)
    if result.Error!= nil {
        return nil, result.Error
    }

    var combinations []entities.Combination

    for _, combination := range combinationsModel {
        combinations = append(combinations, *combination.ToEntity())
    }

    return combinations, nil
}