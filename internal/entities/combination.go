package entities

import (
	"github.com/oklog/ulid/v2"
)

type Combination struct {
	ID           string `json:"id"`
	ListID       string `json:"list_id"`
	FirstItemID  string `json:"first_item_id"`
	SecondItemID string `json:"second_item_id"`
}

func NewCombination(listId, firstItem, secondItem string) *Combination {
	return &Combination{
		ID:           ulid.Make().String(),
		ListID:       listId,
		FirstItemID:  firstItem,
		SecondItemID: secondItem,
	}
}

func (c *Combination) Equals(combination Combination) bool {
	return c.FirstItemID == combination.FirstItemID && c.SecondItemID == combination.SecondItemID && c.ListID == combination.ListID
}
