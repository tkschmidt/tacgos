package level

import tl "github.com/JoelOtter/termloop"

func GenerateLvl() *tl.BaseLevel {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorBlack,
	})
	level.AddEntity(tl.NewRectangle(10, 10, 10, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(10, 10, 1, 2, tl.ColorBlack))
	return level
}
