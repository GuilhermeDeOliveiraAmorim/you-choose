package entities

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/oklog/ulid/v2"
)

type Combination struct {
	ID           string `json:"id"`
	ListID       string `json:"list_id"`
	FirstItemID  string `json:"first_item_id"`
	SecondItemID string `json:"second_item_id"`
}

func NewCombination(listId, firstItem, secondItem string) (*Combination, []util.ProblemDetails) {
	return &Combination{
		ID:           ulid.Make().String(),
		ListID:       listId,
		FirstItemID:  firstItem,
		SecondItemID: secondItem,
	}, nil
}

func (c *Combination) Equals(combination Combination) bool {
	return c.FirstItemID == combination.FirstItemID && c.SecondItemID == combination.SecondItemID && c.ListID == combination.ListID
}
