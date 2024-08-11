package main

import(
	"image/color"
)

type Ball struct {
	radius float32
	center Point
	speed_x float32
	speed_y float32
	color color.Color
}

const (
	gravity = 0.5
	coefficientOfRestitution = 1
)
