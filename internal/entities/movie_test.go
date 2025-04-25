package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	tests := []struct {
		name       string
		year       int64
		externalID string
	}{
		{"Movie 1", 2021, "ext-12345"},
		{"Movie 2", 1995, "ext-67890"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movie, _ := NewMovie(tt.name, tt.year, tt.externalID)

			assert.Equal(t, tt.name, movie.Name)
			assert.Equal(t, tt.year, movie.Year)
			assert.Equal(t, tt.externalID, movie.ExternalID)

			assert.Empty(t, movie.Poster)
		})
	}
}

func TestAddPoster(t *testing.T) {
	movie, _ := NewMovie("Movie 1", 2021, "ext-12345")

	assert.Empty(t, movie.Poster)

	movie.AddPoster("poster_url")

	assert.Equal(t, "poster_url", movie.Poster)
}

func TestUpdatePoster(t *testing.T) {
	movie, _ := NewMovie("Movie 1", 2021, "ext-12345")

	movie.AddPoster("poster_url")

	updatedTime := time.Now()
	movie.UpdatePoster("new_poster_url")

	assert.Equal(t, "new_poster_url", movie.Poster)

	assert.NotNil(t, movie.UpdatedAt)
	assert.WithinDuration(t, updatedTime, *movie.UpdatedAt, time.Second)
}

func TestMovieEquals(t *testing.T) {
	movie1, _ := NewMovie("Movie 1", 2021, "ext-12345")
	movie2, _ := NewMovie("Movie 1", 2021, "ext-12345")
	movie3, _ := NewMovie("Movie 2", 2022, "ext-54321")

	assert.True(t, movie1.Equals(*movie2))

	assert.False(t, movie1.Equals(*movie3))
}
