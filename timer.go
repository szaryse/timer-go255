package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	Break   = "Break"
	Session = "Session"
	Start   = "Start"
)

const (
	TPS       int = 60
	oneMinute int = 60 * TPS
)

type Timer struct {
	count int
	state string
}

func NewTimer() *Timer {
	return &Timer{
		count: 5 * TPS, //25 * oneMinute,
		state: Start,
	}
}

func (t *Timer) Render(screen *ebiten.Image) {
	x, y := 12, 0

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.LineSpacing = fragmentMonoRegularBigFace.Size * 1.5
	op.ColorScale.Scale(192, 192, 192, 255)
	text.Draw(screen, t.state, fragmentMonoRegularBigFace, op)

	x = 160
	timeToRender := countToString(t.count)

	op = &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.LineSpacing = fragmentMonoRegularBigFace.Size * 1.5
	op.ColorScale.Scale(0, 0, 255, 1)
	text.Draw(screen, timeToRender, fragmentMonoRegularBigFace, op)
}

func (t *Timer) Update() error {
	t.count -= 1

	if t.count == 0 && t.state == Start {
		t.state = Session
		t.count = 15 * TPS
	}

	if t.count == 0 && t.state == Session {
		t.state = Break
		t.count = 5 * TPS
	}

	if t.count == 0 && t.state == Break {
		t.state = Session
		t.count = 15 * TPS
	}

	return nil
}

func countToString(count int) string {
	minStr, secStr := "", ""

	minutes := count / oneMinute
	if minutes < 10 {
		minStr += "0"
	}
	minStr += fmt.Sprintf("%d", minutes)

	seconds := (count - minutes*oneMinute) / TPS
	if seconds < 10 {
		secStr += "0"
	}
	secStr += fmt.Sprintf("%d", seconds)

	return fmt.Sprintf("%s:%s", minStr, secStr)
}
