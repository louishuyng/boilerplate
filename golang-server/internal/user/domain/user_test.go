package domain

import "testing"

func TestGetId(t *testing.T) {
	t.Run("Get ID returns 0", func(t *testing.T) {
		domain := NewUserDomain()

		got := domain.GetID()
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
