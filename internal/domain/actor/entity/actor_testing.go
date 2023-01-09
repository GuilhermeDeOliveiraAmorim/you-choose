package domain

import "testing"

func TestCreateActor(t *testing.T) {
	got := -1
	if got != 1 {
		t.Error("Abs(-1) = %d; want 1", got)
	}
}
