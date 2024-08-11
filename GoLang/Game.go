package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{
	balls []*Ball
	tiles []*Tile
	userTile UserTile
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

	for _, rect := range g.tiles{
		if rect.alive{
			vector.DrawFilledRect(screen,rect.X, rect.Y, rect.sizeX, rect.sizeY, rect.color, true)
		}
	}

	vector.DrawFilledRect(screen, userTile.tile.X, userTile.tile.Y, userTile.tile.sizeX, userTile.tile.sizeY, userTile.tile.color, true)
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

		for _, block := range g.tiles{
			if block.alive{
				if (isColliding(float64(ball.center.X), float64(ball.center.Y), float64(ball.radius), block)){
					ball.speed_y *= -coefficientOfRestitution
					block.alive = false
				}
			}

		}
	}
	
	return nil
}

