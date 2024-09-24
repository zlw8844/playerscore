package playerscore

import (
	"github.com/zlw8844/playerscore/player"
)

type PlayerScore struct {
	players map[string]player.Player
}

func NewPlayerScore() *PlayerScore {
	return &PlayerScore{}
}

func (ps *PlayerScore) AddScore(name string) error {
	return nil
}

func (ps *PlayerScore) GetScore(name string) (int, error) {
	return 0, nil
}
