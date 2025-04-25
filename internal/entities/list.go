package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
)

const (
	MOVIE_TYPE = "MOVIE"
	BRAND_TYPE = "BRAND"
)

type List struct {
	SharedEntity
	Name         string        `json:"name"`
	Cover        string        `json:"cover"`
	ListType     string        `json:"list_type"`
	Items        []interface{} `json:"items"`
	Combinations []Combination `json:"combinations"`
}

func NewList(name string, cover string) (*List, []exceptions.ProblemDetails) {
	return &List{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Cover:        cover,
	}, nil
}

func (l *List) AddItems(items []interface{}) {
	if len(l.Items) == 0 {
		l.Items = items
		return
	}

	for _, newItem := range items {
		newID := extractID(newItem)
		if newID == "" {
			continue
		}

		exists := false
		for _, existingItem := range l.Items {
			if extractID(existingItem) == newID {
				exists = true
				break
			}
		}

		if !exists {
			l.Items = append(l.Items, newItem)
		}
	}
}

func extractID(item interface{}) string {
	switch v := item.(type) {
	case Movie:
		return v.ID
	case Brand:
		return v.ID
	default:
		return ""
	}
}

func (l *List) ClearItems() {
	l.Items = []interface{}{}
}

func (l *List) AddCombinations(combinations []Combination) {
	if len(l.Combinations) == 0 {
		l.Combinations = combinations
		return
	}

	uniqueCombinations := []Combination{}

	for _, newCombination := range combinations {
		exists := false
		for _, existingCombination := range l.Combinations {
			if existingCombination.Equals(newCombination) {
				fmt.Println("Combination already exists:", existingCombination)
				exists = true
				break
			}
		}

		if !exists {
			uniqueCombinations = append(uniqueCombinations, newCombination)
		}
	}

	l.Combinations = append(l.Combinations, uniqueCombinations...)
}

func (l *List) GetCombinations(itemIDs []string) []Combination {
	var combinations []Combination

	for i := range itemIDs {
		for j := i + 1; j < len(itemIDs); j++ {
			newCombination := NewCombination(l.ID, itemIDs[i], itemIDs[j])
			combinations = append(combinations, *newCombination)
		}
	}

	return combinations
}

func (l *List) GetItemIDs() []string {
	itemIDs := []string{}

	for _, item := range l.Items {
		switch item := item.(type) {
		case Movie:
			itemIDs = append(itemIDs, item.ID)
		}
	}

	return itemIDs
}

func (l *List) AddCover(cover string) {
	l.Cover = cover
}

func (l *List) UpdateCover(cover string) {
	timeNow := time.Now()
	l.UpdatedAt = &timeNow

	l.Cover = cover
}

func (l *List) AddType(ListType string) {
	l.ListType = ListType
}

func (l *List) GetTypes() []string {
	return []string{
		MOVIE_TYPE,
		BRAND_TYPE,
	}
}

func (l *List) FormatRanking(rankItems []interface{}) (interface{}, error) {
	switch l.ListType {
	case MOVIE_TYPE:
		movies := make([]Movie, len(rankItems))
		for i, item := range rankItems {
			movie, ok := item.(Movie)
			if !ok {
				return nil, errors.New("failed to cast item to Movie")
			}
			movies[i] = movie
		}
		return movies, nil
	case BRAND_TYPE:
		brands := make([]Brand, len(rankItems))
		for i, item := range rankItems {
			brand, ok := item.(Brand)
			if !ok {
				return nil, errors.New("failed to cast item to Brand")
			}
			brands[i] = brand
		}
		return brands, nil
	default:
		return nil, errors.New("unsupported list type")
	}
}
