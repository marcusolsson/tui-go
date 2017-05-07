package tui

import "image"

type Key int

const (
	KeyUnknown Key = iota
	KeyEnter
	KeySpace
	KeyTab
	KeyBacktab
	KeyEsc
	KeyBackspace
	KeyBackspace2
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
)

type KeyEvent struct {
	Key  Key
	Rune rune
}

type MouseEvent struct {
	Pos image.Point
}

type PaintEvent struct{}

type Event interface{}
