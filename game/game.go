package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pong_game/models"
)

type Game struct {
	LeftPaddle  models.Paddle
	RightPaddle models.Paddle
	Ball        models.Ball

	ShowExitPrompt bool

	LeftScore  int
	RightScore int
}

func NewGame() *Game {
	return &Game{
		LeftPaddle: models.Paddle{
			X:      20,
			Y:      250,
			Width:  10,
			Height: 100,
			Speed:  5,
		},
		RightPaddle: models.Paddle{
			X:      770,
			Y:      250,
			Width:  10,
			Height: 100,
			Speed:  5,
		},
		Ball: models.Ball{
			X:    400,
			Y:    300,
			DX:   4,
			DY:   4,
			Size: 10,
		},
	}
}

func (g *Game) Update() {
	// ---- 1. Check ESC for exit prompt ----
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.ShowExitPrompt = true
	}

	// ---- 2. If exit prompt is active, only handle Yes/No keys ----
	if g.ShowExitPrompt {
		if ebiten.IsKeyPressed(ebiten.KeyY) {
			// Quit the game
			// Ebiten does not have a direct quit call, so call runtime exit
			os.Exit(0) // simple way to exit
		}
		if ebiten.IsKeyPressed(ebiten.KeyN) {
			// Cancel exit
			g.ShowExitPrompt = false
		}

		// Skip game update while in prompt
		return
	}
	// Paddle movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.LeftPaddle.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.LeftPaddle.MoveDown()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.RightPaddle.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.RightPaddle.MoveDown()
	}

	g.Ball.Update()

	// top/bottom collision
	if g.Ball.Y <= 0 || g.Ball.Y >= 600-g.Ball.Size {
		g.Ball.DY *= -1
	}

	// Paddle Collision
	if g.Ball.X <= g.LeftPaddle.X+g.LeftPaddle.Width &&
		g.Ball.Y+g.Ball.Size >= g.LeftPaddle.Y &&
		g.Ball.Y <= g.LeftPaddle.Y+g.LeftPaddle.Height {
		g.Ball.DX *= -1
	}

	if g.Ball.X+g.Ball.Size >= g.RightPaddle.X &&
		g.Ball.Y+g.Ball.Size >= g.RightPaddle.Y &&
		g.Ball.Y <= g.RightPaddle.Y+g.RightPaddle.Height {
		g.Ball.DX *= -1
	}

	// Left player missed
	if g.Ball.X < 0 {
		g.RightScore++ // Right player gets point
		g.Ball.X = 400 - g.Ball.Size/2
		g.Ball.Y = 300 - g.Ball.Size/2
		g.Ball.DX = 4 // serve toward the player who lost
	}

	// Right player missed
	if g.Ball.X > 800 {
		g.LeftScore++ // Left player gets point
		g.Ball.X = 400 - g.Ball.Size/2
		g.Ball.Y = 300 - g.Ball.Size/2
		g.Ball.DX = -4 // serve toward the player who lost
	}

	//	// This is no longer required due to the presence of the scoreboard
	// if g.Ball.X < 0 || g.Ball.X > 800 { // assuming screen width = 800
	// 	// reset ball to center
	// 	g.Ball.X = 400 - g.Ball.Size/2
	// 	g.Ball.Y = 300 - g.Ball.Size/2

	// 	// optionally reverse X velocity so it goes toward the player who lost
	// 	g.Ball.DX = -g.Ball.DX

	// 	// optional: randomize Y velocity slightly to avoid straight lines
	// 	// g.Ball.VY = rand.Float64()*4 - 2 // random -2..2
	// }

}

func (g *Game) Draw(screen *ebiten.Image) {

	// Draw left and right scores
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.LeftScore), 200, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.RightScore), 600, 20)

	ebitenutil.DrawRect(
		screen,
		g.LeftPaddle.X,
		g.LeftPaddle.Y,
		g.LeftPaddle.Width,
		g.LeftPaddle.Height,
		color.White,
	)

	ebitenutil.DrawRect(
		screen,
		g.RightPaddle.X,
		g.RightPaddle.Y,
		g.RightPaddle.Width,
		g.RightPaddle.Height,
		color.White,
	)

	ebitenutil.DrawRect(
		screen,
		g.Ball.X,
		g.Ball.Y,
		g.Ball.Size,
		g.Ball.Size,
		color.White,
	)

	// Show exit prompt overlay
	if g.ShowExitPrompt {
		ebitenutil.DebugPrintAt(screen, "Quit game? (Y/N)", 300, 280)
	}
}
