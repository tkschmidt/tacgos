package player

import tl "github.com/JoelOtter/termloop"

const (
	RightArrow = '→'
	LeftArrow  = '←'
	TopArrow   = '↑'
	DownArrow  = '↓'
)

type Player struct {
	entity      *tl.Entity
	prevX       int
	prevY       int
	level       *tl.BaseLevel
	active      bool
	direction   int
	visionRange int
}

func GeneratePlayer(level *tl.BaseLevel) *Player {
	newPlayer := Player{level: level,
		entity:      tl.NewEntity(10, 5, 5, 5),
		active:      false,
		direction:   1,
		visionRange: 2,
	}

	return &newPlayer

}

func (player *Player) CalculateVision() {
	nums := []int{0, 1, 2, 3, 4}
	for _, num := range nums {
		for _, num2 := range nums {
			player.entity.SetCell(num, num2, &tl.Cell{Ch: ' '})
		}
	}
	switch player.direction {
	case 1:
		player.entity.SetCell(1, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(2, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(1, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(2, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 0, &tl.Cell{Ch: '.'})
	case 2:
		player.entity.SetCell(3, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 2, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 2, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 4, &tl.Cell{Ch: '.'})
	case 3:
		player.entity.SetCell(1, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(2, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 4, &tl.Cell{Ch: '.'})
		player.entity.SetCell(1, 4, &tl.Cell{Ch: '.'})
		player.entity.SetCell(2, 4, &tl.Cell{Ch: '.'})
		player.entity.SetCell(3, 4, &tl.Cell{Ch: '.'})
		player.entity.SetCell(4, 4, &tl.Cell{Ch: '.'})
	case 4:
		player.entity.SetCell(1, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(1, 2, &tl.Cell{Ch: '.'})
		player.entity.SetCell(1, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 0, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 1, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 2, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 3, &tl.Cell{Ch: '.'})
		player.entity.SetCell(0, 4, &tl.Cell{Ch: '.'})
	}
}
func (player *Player) Draw(screen *tl.Screen) {
	player.CalculateVision()
	switch player.direction {
	case 1:
		player.entity.SetCell(2, 2, &tl.Cell{Ch: TopArrow})
	case 2:
		player.entity.SetCell(2, 2, &tl.Cell{Ch: RightArrow})
	case 3:
		player.entity.SetCell(2, 2, &tl.Cell{Ch: DownArrow})
	case 4:
		player.entity.SetCell(2, 2, &tl.Cell{Ch: LeftArrow})
	}
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		if player.active {
			player.prevX, player.prevY = player.entity.Position()
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				if player.direction == 2 {
					player.entity.SetPosition(player.prevX+1, player.prevY)
					player.direction = 2
				} else {
					player.direction = 2
				}
				break
			case tl.KeyArrowLeft:
				if player.direction == 4 {
					player.entity.SetPosition(player.prevX-1, player.prevY)
					player.direction = 4
				} else {
					player.direction = 4
				}
				break
			case tl.KeyArrowUp:
				if player.direction == 1 {
					player.entity.SetPosition(player.prevX, player.prevY-1)
				} else {
					player.direction = 1
				}
				break
			case tl.KeyArrowDown:
				if player.direction == 3 {
					player.entity.SetPosition(player.prevX, player.prevY+1)
				} else {
					player.direction = 3
				}
				break
			}
		} else {
			if event.Key == tl.KeyEnter {
				player.active = true
			}
		}
	}
}

func (player *Player) Size() (int, int)     { return player.entity.Size() }
func (player *Player) Position() (int, int) { return player.entity.Position() }
