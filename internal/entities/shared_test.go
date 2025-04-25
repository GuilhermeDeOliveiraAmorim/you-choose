package entities

import (
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewSharedEntity(t *testing.T) {
	t.Run("should create a shared entity with valid fields", func(t *testing.T) {

		sharedEntity := NewSharedEntity()

		assert.NotEmpty(t, sharedEntity.ID)

		assert.True(t, sharedEntity.Active)

		assert.WithinDuration(t, time.Now(), sharedEntity.CreatedAt, time.Second)

		assert.Nil(t, sharedEntity.UpdatedAt)
		assert.Nil(t, sharedEntity.DeactivatedAt)
	})
}

func TestActivate(t *testing.T) {
	t.Run("should activate an entity", func(t *testing.T) {

		sharedEntity := NewSharedEntity()
		sharedEntity.Deactivate()

		assert.False(t, sharedEntity.Active)
		assert.NotNil(t, sharedEntity.DeactivatedAt)

		sharedEntity.Activate()

		assert.True(t, sharedEntity.Active)
		assert.NotNil(t, sharedEntity.UpdatedAt)
		assert.Nil(t, sharedEntity.DeactivatedAt)
	})
}

func TestDeactivate(t *testing.T) {
	t.Run("should deactivate an entity", func(t *testing.T) {

		sharedEntity := NewSharedEntity()

		assert.True(t, sharedEntity.Active)

		sharedEntity.Deactivate()

		assert.False(t, sharedEntity.Active)
		assert.NotNil(t, sharedEntity.DeactivatedAt)
		assert.NotNil(t, sharedEntity.UpdatedAt)
	})
}

func TestULIDGeneration(t *testing.T) {
	t.Run("should generate a unique ID for each shared entity", func(t *testing.T) {

		sharedEntity1 := NewSharedEntity()
		sharedEntity2 := NewSharedEntity()

		assert.NotEqual(t, sharedEntity1.ID, sharedEntity2.ID)

		_, err := ulid.Parse(sharedEntity1.ID)
		assert.NoError(t, err)
	})
}
