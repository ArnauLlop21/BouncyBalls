package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return screenWidthGlobal, screenHeightGlobal
}


// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    // Creem colors
    lila := color.RGBA{128, 0, 128, 255}
    blau := color.RGBA{0, 0, 255, 255}
    negre := color.RGBA{0, 0, 0, 255}

    // Omplir la meitat esquerra amb lila
    lilaRect := screen.SubImage(image.Rect(0, 0, screenWidthGlobal/2, screenHeightGlobal)).(*ebiten.Image)
    lilaRect.Fill(lila)

    // Omplir la meitat superior dreta amb blau
    blauRect := screen.SubImage(image.Rect(screenWidthGlobal/2, 0, screenWidthGlobal, screenHeightGlobal/2)).(*ebiten.Image)
    blauRect.Fill(blau)

    // Omplir la meitat inferior dreta amb negre
    negreRect := screen.SubImage(image.Rect(screenWidthGlobal/2, screenHeightGlobal/2, screenWidthGlobal, screenHeightGlobal)).(*ebiten.Image)
    negreRect.Fill(negre)
}

// This method is called every tick (tipically 60 times per second)
func (g *Game) Update() error{
	//Write here the logical update for the game
	return nil
}