package cursor

import tl "github.com/JoelOtter/termloop"

type Cursor struct {
	entity *tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
	active bool
}

func GenerateCursor(level *tl.BaseLevel) *Cursor {
	newCursor := Cursor{level: level,
		entity: tl.NewEntity(4, 4, 1, 1),
		active: true}
	newCursor.entity.SetCell(0, 0, &tl.Cell{Fg: 53, Bg: tl.ColorCyan})
	return &newCursor

}

func (cursor *Cursor) Draw(screen *tl.Screen) {
	if cursor.active {
		screenWidth, screenHeight := screen.Size()
		x, y := cursor.entity.Position()
		cursor.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
		cursor.entity.Draw(screen)
	}
}

func (cursor *Cursor) Tick(event tl.Event) {
	if cursor.active {
		if event.Type == tl.EventKey { // Is it a keyboard event?
			cursor.prevX, cursor.prevY = cursor.entity.Position()
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				cursor.entity.SetPosition(cursor.prevX+1, cursor.prevY)
				break
			case tl.KeyArrowLeft:
				cursor.entity.SetPosition(cursor.prevX-1, cursor.prevY)
				break
			case tl.KeyArrowUp:
				cursor.entity.SetPosition(cursor.prevX, cursor.prevY-1)
				break
			case tl.KeyArrowDown:
				cursor.entity.SetPosition(cursor.prevX, cursor.prevY+1)
				break
			case tl.KeyEnter:
				cursor.entity.SetCell(0, 0, &tl.Cell{Fg: 53, Bg: tl.ColorCyan})
				break
			}

		}
	}
}

func (cursor *Cursor) Size() (int, int)     { return cursor.entity.Size() }
func (cursor *Cursor) Position() (int, int) { return cursor.entity.Position() }
