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
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight)){
		
		if(userTile.tile.X + userTile.tile.sizeX >= screenWidthGlobal){
			userTile.tile.X = screenWidthGlobal - userTile.tile.sizeX
		}else{
			userTile.tile.X += 5
		}

	}else if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft)){
		if(userTile.tile.X <= 0){
			userTile.tile.X = 0
		}else{
			userTile.tile.X -= 5
		}
	}

	for _, ball := range g.balls{
		ball.speed_y -= gravity
		ball.center.Y -= ball.speed_y
		ball.center.X += ball.speed_x

		if (ball.center.Y + ball.radius >= screenHeightGlobal){
			ball.center.Y = screenHeightGlobal - ball.radius
			ball.speed_y *= -coefficientOfRestitution

		}else if (ball.center.Y - ball.radius <= 0){
			ball.center.Y = 0 + ball.radius
			ball.speed_y *= -coefficientOfRestitution
		}

		if(ball.center.X + ball.radius >= screenWidthGlobal){
			ball.center.X = screenWidthGlobal - ball.radius
			ball.speed_x *= -coefficientOfRestitution
		}else if(ball.center.X - ball.radius <= 0){
			ball.center.X = 0 + ball.radius
			ball.speed_x *= -coefficientOfRestitution
		}

		for _, block := range g.tiles{
			if block.alive{
				if (isColliding(float64(ball.center.X), float64(ball.center.Y), float64(ball.radius), block)){
					collisionReason := determineCollisionType(*ball, *block)
					switch collisionReason {
					case "left":
					case "right":
						ball.speed_x *= -coefficientOfRestitution
					case "top":
					case "bottom":
						ball.speed_y *= -coefficientOfRestitution
					}
					
					block.alive = false
				}
			}

		}

		if(isColliding(float64(ball.center.X), float64(ball.center.Y), float64(ball.radius), &userTile.tile)){
			
			if(ball.center.X < (userTile.tile.X + userTile.tile.sizeX/2)){
				ball.speed_x -= 1
			}else if (ball.center.X > (userTile.tile.X + userTile.tile.sizeX/2)) {
				ball.speed_x += 1
			}
			
			ball.speed_y *= -coefficientOfRestitution
		}
	}
	
	return nil
}

