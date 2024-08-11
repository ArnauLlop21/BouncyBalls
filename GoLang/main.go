package main

import (
    "github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	 screenWidthGlobal = 900
	 screenHeightGlobal = 600
)
	
// This method has nothing to do with the ebiten library.
// It is for internal use and it initializes the balls.
func initializeBalls(){

}



func main() {
    game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(screenWidthGlobal, screenHeightGlobal)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
