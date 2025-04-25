package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBrand(t *testing.T) {
	brand, problems := NewBrand("Nike", "logo.png")

	assert.Nil(t, problems)
	assert.Equal(t, "Nike", brand.Name)
	assert.Equal(t, "logo.png", brand.Logo)
	assert.NotZero(t, brand.ID)
	assert.NotZero(t, brand.CreatedAt)
}

func TestBrand_UpdateLogo(t *testing.T) {
	brand, _ := NewBrand("Adidas", "old_logo.png")
	previousTime := brand.UpdatedAt

	brand.UpdateLogo("new_logo.png")

	assert.Equal(t, "new_logo.png", brand.Logo)
	assert.NotNil(t, brand.UpdatedAt)
	if previousTime != nil {
		assert.True(t, brand.UpdatedAt.After(*previousTime))
	}
}

func TestBrand_Equals(t *testing.T) {
	brand1, _ := NewBrand("Puma", "logo1.png")
	brand2, _ := NewBrand("Puma", "logo2.png")
	brand3, _ := NewBrand("Reebok", "logo3.png")

	assert.True(t, brand1.Equals(*brand2))
	assert.False(t, brand1.Equals(*brand3))
}
