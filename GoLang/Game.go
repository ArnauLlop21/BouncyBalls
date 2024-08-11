package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{
	balls []*Ball
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return screenWidthGlobal, screenHeightGlobal
}


// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	for _, ball := range g.balls{
		vector.DrawFilledCircle(screen, float32(ball.center.X), float32(ball.center.Y), ball.radius, ball.color, true)
	}
}

// This method is called every tick (tipically 60 times per second)
func (g *Game) Update() error{
	//Write here the logical update for the game
	for _, ball := range g.balls{
		ball.speed_y -= gravity
		ball.center.Y -= ball.speed_y

		if (ball.center.Y >= screenHeightGlobal){
			ball.center.Y = screenHeightGlobal - ball.radius
			ball.speed_y *= -coefficientOfRestitution

		}else if (ball.center.Y <= 0){
			ball.center.Y = 0 + ball.radius
			ball.speed_y *= -coefficientOfRestitution
		}
	}
	
	return nil
}