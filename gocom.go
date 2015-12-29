package main

import (
	tl "github.com/JoelOtter/termloop"
	//	"github.com/tkschmidt/gocom/bot"
	"github.com/tkschmidt/gocom/cursor"
	"github.com/tkschmidt/gocom/level"
	"github.com/tkschmidt/gocom/player"
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
