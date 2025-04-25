package entities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list, err := NewList("Top 10 Filmes", "some-cover-url")
	assert.NotNil(t, list)
	assert.Nil(t, err)
	assert.Equal(t, "Top 10 Filmes", list.Name)
	assert.Equal(t, "some-cover-url", list.Cover)
}

func TestAddItems(t *testing.T) {
	list, _ := NewList("Minha Lista", "")
	item1 := Movie{
		SharedEntity: SharedEntity{
			ID: "m1",
		},
		Name: "Filme 1",
	}
	item2 := Movie{
		SharedEntity: SharedEntity{
			ID: "m1",
		},
		Name: "Filme 1",
	}
	list.AddItems([]interface{}{item1, item2})

	assert.Len(t, list.Items, 2)

	list.AddItems([]interface{}{item1})

	assert.Len(t, list.Items, 2)
}

func TestAddCombinations(t *testing.T) {
	list := List{
		SharedEntity: SharedEntity{
			ID: "list1",
		},
		Name: "Lista de Teste",
	}

	// Cria combinações
	comb1 := Combination{ListID: list.ID, FirstItemID: "a", SecondItemID: "b"}
	comb2 := Combination{ListID: list.ID, FirstItemID: "a", SecondItemID: "b"} // Duplicado
	comb3 := Combination{ListID: list.ID, FirstItemID: "b", SecondItemID: "c"}

	fmt.Println(comb1.Equals(comb2)) // true
	fmt.Println(comb1.Equals(comb3)) // false

	// Adiciona comb1 e comb2 (duplicado)
	list.AddCombinations([]Combination{comb1, comb2})
	// Espera que comb1 e comb2 sejam tratadas como única, ou seja, comb2 não deve ser adicionada.

	// Verifica o resultado após a adição das duas primeiras combinações
	if len(list.Combinations) != 1 {
		t.Errorf("Combinations after adding comb1 and comb2: %v, expected 1 combination.", list.Combinations)
	}

	// Adiciona comb3
	list.AddCombinations([]Combination{comb3})

	// Verifica o resultado após adicionar comb3
	if len(list.Combinations) != 2 {
		t.Errorf("Combinations after adding comb3: %v, expected 2 combinations.", list.Combinations)
	}
}

func TestGetCombinations(t *testing.T) {
	list, _ := NewList("Teste", "")
	list.ID = "xyz"

	combinations := list.GetCombinations([]string{"1", "2", "3"})

	assert.Len(t, combinations, 3)
	assert.Equal(t, "1", combinations[0].FirstItemID)
	assert.Equal(t, "2", combinations[0].SecondItemID)
}

func TestFormatRanking(t *testing.T) {
	list, _ := NewList("Test Format", "")
	list.AddType(MOVIE_TYPE)

	movie := Movie{
		SharedEntity: SharedEntity{
			ID: "m1",
		},
		Name: "Test Movie",
	}
	ranked, err := list.FormatRanking([]interface{}{movie})
	assert.Nil(t, err)
	assert.Len(t, ranked.([]Movie), 1)

	list.ListType = "UNSUPPORTED"
	_, err = list.FormatRanking([]interface{}{movie})
	assert.Error(t, err)
}
