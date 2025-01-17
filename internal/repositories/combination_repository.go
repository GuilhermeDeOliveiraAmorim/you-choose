package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type CombinationRepository interface {
	GetCombinationsByListID(listID string) ([]entities.Combination, error)
	GetCombinationsAlreadyVoted(listID string) ([]entities.Combination, error)
}
