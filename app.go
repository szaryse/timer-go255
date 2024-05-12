package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed FragmentMono-Regular.ttf
var fragmentMonoRegular []byte

var (
	fragmentMonoRegularSource     *text.GoTextFaceSource
	fragmentMonoRegularNormalFace *text.GoTextFace
	fragmentMonoRegularBigFace    *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fragmentMonoRegular))
	if err != nil {
		log.Fatal(err)
	}
	fragmentMonoRegularSource = s

	fragmentMonoRegularNormalFace = &text.GoTextFace{
		Source: fragmentMonoRegularSource,
		Size:   16,
	}
	fragmentMonoRegularBigFace = &text.GoTextFace{
		Source: fragmentMonoRegularSource,
		Size:   32,
	}
}

type App struct {
	timer *Timer
}

func NewApp() *App {
	return &App{
		timer: NewTimer(),
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

func (a *App) Update() error {
	a.timer.Update()
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.timer.Render(screen)
}
