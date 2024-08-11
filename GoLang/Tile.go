package main

import(
	"image/color"
)

type Tile struct{
	color color.Color
	sizeX float32
	sizeY float32
	X float32
	Y float32
	alive bool
}

type UserTile struct{
	tile Tile
}
