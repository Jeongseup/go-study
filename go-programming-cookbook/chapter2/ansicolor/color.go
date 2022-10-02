package ansicolor

import "fmt"

type Color int

const (
	ColorNone = iota
	Red
	Green
	Yello
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

// ColorText 구조체는 색상의 문자열 및 색상 값을 저장한다
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {
	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}
