package main

import (
	tl "github.com/JoelOtter/termloop"
	//	"github.com/tkschmidt/tacgos/bot"
	"github.com/tkschmidt/tacgos/cursor"
	"github.com/tkschmidt/tacgos/level"
	"github.com/tkschmidt/tacgos/player"
)

func main() {
	game := tl.NewGame()
	level := level.GenerateLvl()
	pl := player.GeneratePlayer(level)
	cursor := cursor.GenerateCursor(level)
	level.AddEntity(pl)
	level.AddEntity(cursor)
	game.Screen().SetLevel(level)
	game.SetDebugOn(true)
	game.Start()
}
