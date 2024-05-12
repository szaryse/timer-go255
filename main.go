package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  float64 = 360.0
	screenHeight float64 = 360.0
)

func main() {
	app := NewApp()

	ebiten.SetWindowFloating(true)
	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Timer Go255 v0.1")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
