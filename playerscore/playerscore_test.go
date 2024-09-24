package playerscore_test

import (
	"github.com/zlw8844/playerscore/playerscore"
	"testing"
)

func TestPlayerScore(t *testing.T) {
	t.Run("TestPlayerStore", func(t *testing.T) {
		ps := playerscore.NewPlayerScore()
		ps.AddScore("player1")
		got, _ := ps.GetScore("player1")
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
