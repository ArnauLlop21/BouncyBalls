package main

import (
    "github.com/hajimehoshi/ebiten/v2"
	"log"
	"image/color"
)

const (
	 screenWidthGlobal = 900
	 screenHeightGlobal = 600
)

var balls = []*Ball{
	{radius: 30, center: NewPoint(50,50), speed_x: 0, speed_y: 0, color: color.RGBA{255,0,0,255}},
	{radius: 25, center: NewPoint(500,500), speed_x: 0, speed_y: 0, color: color.RGBA{0,0,255,255}},
	{radius: 15, center: NewPoint(200,200), speed_x: 0, speed_y: 0, color: color.RGBA{0,255,0,255}},
}



func main() {
    game := &Game{
		balls: balls,
	}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidthGlobal, screenHeightGlobal)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
