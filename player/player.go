package player

import tl "github.com/JoelOtter/termloop"

type Player struct {
	entity *tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
	active bool
}

func GeneratePlayer(level *tl.BaseLevel) *Player {
	newPlayer := Player{level: level,
		entity: tl.NewEntity(3, 1, 4, 4),
		active: false}
	newPlayer.entity.SetCell(0, 0, &tl.Cell{Ch: '↑'})
	newPlayer.entity.SetCell(0, 1, &tl.Cell{Ch: '.'})
	newPlayer.entity.SetCell(1, 1, &tl.Cell{Ch: '.'})
	newPlayer.entity.SetCell(3, 0, &tl.Cell{Ch: '.'})

	return &newPlayer

}

func (player *Player) Draw(screen *tl.Screen) {
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if player.active {
		if event.Type == tl.EventKey { // Is it a keyboard event?
			player.prevX, player.prevY = player.entity.Position()
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				player.entity.SetPosition(player.prevX+1, player.prevY)
				break
			case tl.KeyArrowLeft:
				player.entity.SetPosition(player.prevX-1, player.prevY)
				break
			case tl.KeyArrowUp:
				player.entity.SetPosition(player.prevX, player.prevY-1)
				break
			case tl.KeyArrowDown:
				player.entity.SetPosition(player.prevX, player.prevY+1)
				break
			}
		}
	}
}

func (player *Player) Size() (int, int)     { return player.entity.Size() }
func (player *Player) Position() (int, int) { return player.entity.Position() }
