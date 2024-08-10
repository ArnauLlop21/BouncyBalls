// CODE EXTRACTED FROM
// https://go.dev/play/p/5KL4HipSJ-

package main

import (
	"math"
)


type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point{return Point{x,y}}

func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}