package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pong_game/game"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Engine struct {
	game *game.Game
}

func (e *Engine) Draw(screen *ebiten.Image) {
	e.game.Draw(screen)
}

func (e *Engine) Update() error {
	e.game.Update()
	return nil
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {

	g := game.NewGame()

	engine := &Engine{
		game: g,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Pong")

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}
