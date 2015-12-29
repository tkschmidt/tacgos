package bot

import tl "github.com/JoelOtter/termloop"

type Bot struct {
	entity *tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
	vision bool
}

func (bot *Bot) Draw(screen *tl.Screen) {
	if bot.vision {
		bot.entity.Draw(screen)
	}
}

func NewBot() *Bot {
	newBot := Bot{
		entity: tl.NewEntity(1, 1, 1, 1),
		vision: true,
	}
	return &newBot
}

func (bot *Bot) Tick(event tl.Event) {
}

func (bot *Bot) Size() (int, int)     { return bot.entity.Size() }
func (bot *Bot) Position() (int, int) { return bot.entity.Position() }
