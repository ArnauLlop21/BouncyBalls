package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var BALLS []Ball

func initializations() fyne.Window {
	a := app.New()
	window := a.NewWindow("My New Window")
	window.Resize(fyne.NewSize(900,600))
	ball := Ball{
		radius:   50,
		circle_x: 10,
		circle_y: 10,
		speed_x:  10,
		speed_y:  10,
		color:    [3]byte{255, 0, 0},
	}
	//ball2 := Ball{
	//	radius:   10,
	//	circle_x: 200,
	//	circle_y: 200,
	//	speed_x:  10,
	//	speed_y:  10,
	//	color:    [3]byte{0, 255, 0},
	//}
	BALLS = append(BALLS, ball)
	//BALLS = append(BALLS, ball2)
	return window
}

func DrawCircles(window fyne.Window) {

    container := container.NewWithoutLayout()
    for i := 0; i < len(BALLS); i++ {
		circle := canvas.NewCircle(color.RGBA{
            R: BALLS[i].color[0],
            G: BALLS[i].color[1],
            B: BALLS[i].color[2],
            A: 0xff,
        })
        circle.StrokeColor = color.Gray{Y: 0x99}
        circle.StrokeWidth = 1

        circle.Resize(fyne.NewSize(BALLS[i].radius, BALLS[i].radius))

        circle.Move(fyne.NewPos(float32(BALLS[i].circle_x), float32(BALLS[i].circle_y)))

        container.Add(circle)
    }

    window.SetContent(container)
}


func main() {
	fmt.Println("Hello Bouncy Balls!")
	window := initializations()
	DrawCircles(window)
	window.ShowAndRun()
}
