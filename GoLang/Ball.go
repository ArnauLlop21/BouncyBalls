package main

import(
	"image/color"
	"math"
)

type Ball struct {
	radius float32
	center Point
	speed_x float32
	speed_y float32
	color color.Color
}

const (
	gravity = 0
	coefficientOfRestitution = 1
)

// Funció per detectar col·lisió entre la pilota i un rectangle
func isColliding(ballX, ballY, radius float64, tile *Tile) bool {
    // Calcular les distàncies més properes entre el centre de la pilota i els límits del rectangle
    nearestX := math.Max(float64(tile.X), math.Min(ballX, float64(tile.X + tile.sizeX)))
    nearestY := math.Max(float64(tile.Y), math.Min(ballY, float64(tile.Y + tile.sizeY)))

    // Calcular la distància entre el centre de la pilota i el punt més proper al rectangle
    deltaX := ballX - nearestX
    deltaY := ballY - nearestY

    // Si la distància entre el punt més proper i el centre de la pilota és menor que el radi, hi ha col·lisió
    return (deltaX*deltaX + deltaY*deltaY) < (radius * radius)
}

