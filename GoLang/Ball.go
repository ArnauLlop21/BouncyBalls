package main

import(
	"image/color"
)

type Ball struct {
	radius float32
	center Point
	speed_x int
	speed_y int
	color color.Color
}

