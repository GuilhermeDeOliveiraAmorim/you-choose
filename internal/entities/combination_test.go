package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCombination(t *testing.T) {
	comb := NewCombination("list123", "itemA", "itemB")

	assert.NotEmpty(t, comb.ID)
	assert.Equal(t, "list123", comb.ListID)
	assert.Equal(t, "itemA", comb.FirstItemID)
	assert.Equal(t, "itemB", comb.SecondItemID)
}

func TestCombination_Equals_True(t *testing.T) {
	c1 := NewCombination("list1", "item1", "item2")
	c2 := NewCombination("list1", "item1", "item2")

	assert.True(t, c1.Equals(*c2))
}

func TestCombination_Equals_False_DifferentListID(t *testing.T) {
	c1 := NewCombination("list1", "item1", "item2")
	c2 := NewCombination("list2", "item1", "item2")

	assert.False(t, c1.Equals(*c2))
}

func TestCombination_Equals_False_DifferentItems(t *testing.T) {
	c1 := NewCombination("list1", "item1", "item2")
	c2 := NewCombination("list1", "item2", "item1")

	assert.False(t, c1.Equals(*c2))
}
