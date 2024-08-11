package main

import (
    "github.com/hajimehoshi/ebiten/v2"
	"log"
	"image/color"
)

const (
	screenWidthGlobal = 900
	screenHeightGlobal = 600
	blockSizeX = 100
	blockSizeY = 40
	blockMargin = 10
	topThirdHeight = screenHeightGlobal / 3
)

var balls = []*Ball{
	{radius: 10, center: NewPoint(screenWidthGlobal/2,screenHeightGlobal/2), speed_x: 0, speed_y: -5, color: color.RGBA{255,0,0,255}},
	//{radius: 25, center: NewPoint(500,500), speed_x: 0, speed_y: 0, color: color.RGBA{0,0,255,255}},
	//{radius: 15, center: NewPoint(200,200), speed_x: 0, speed_y: 0, color: color.RGBA{0,255,0,255}},
}

var tiles = []*Tile{}
var tile = Tile{
	color: color.RGBA{0, 240, 0, 255}, // Color vermell
	sizeX: blockSizeX,                       // Amplada de 50 píxels
	sizeY: blockSizeY,                       // Alçada de 20 píxels
	X:     (screenWidthGlobal/2)- blockSizeX/2,                      // Coordenada X
	Y:     500,                      // Coordenada Y
	alive: true,                       // El bloc està actiu (alive)
}

var userTile = UserTile{
	tile: tile,
}

func tilesInitializations(){
    // Calcula el nombre de blocs que caben a l'amplada i a l'altura amb l'espai entre blocs
    blocksPerRow := (screenWidthGlobal + blockMargin) / (blockSizeX + blockMargin)
    blocksPerColumn := (topThirdHeight + blockMargin) / (blockSizeY + blockMargin)

    // Inicialitza els blocs i afegeix-los al slice
    for i := 0; i < blocksPerColumn; i++ {
        for j := 0; j < blocksPerRow; j++ {
            tile := &Tile{
                X:     float32(j*(blockSizeX+blockMargin) + blockMargin/2),
                Y:     float32(i*(blockSizeY+blockMargin) + blockMargin/2),
                sizeX: blockSizeX,
                sizeY: blockSizeY,
                color: color.White, // Pots canviar el color si vols
				alive: true,
            }
            tiles = append(tiles, tile)
        }
    }
}

func main() {
    tilesInitializations()
	game := &Game{
		balls: balls,
		tiles: tiles,
		userTile: userTile,
	}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidthGlobal, screenHeightGlobal)
    ebiten.SetWindowTitle("Brick breaker")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
