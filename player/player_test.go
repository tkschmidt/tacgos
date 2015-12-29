package player

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/tkschmidt/tacgos/level"
	"testing"
)

func TestVision(t *testing.T) {
	game := tl.NewGame()
	level := level.GenerateLvl()
	pl := GeneratePlayer(level, game)
	if pl.direction != 1 {
		t.Error("Expected 1", pl.direction)
	}
}
