package player

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/tkschmidt/tacgos/cursor"
	"github.com/tkschmidt/tacgos/level"
	"github.com/tkschmidt/tacgos/player"
	"testing"
)

func TestVision(t *testing.T) {
	game := tl.NewGame()
	level := level.GenerateLvl()
	pl := player.GeneratePlayer(level, game)
	if 1 != 1 {
		t.Error("Expected True")
	}
}
