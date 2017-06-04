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
	Mod  ModMask
}

type MouseEvent struct {
	Pos image.Point
}

type PaintEvent struct{}

type Event interface{}

type HEvent struct {
	Key       Key
	Rune      rune
	Modifiers ModMask
}

type ModMask int16

const (
	ModShift ModMask = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
	ModNone ModMask = 0
)

// When an Event is fired in tcell, the ev.Ch pressed
// is modified by Ctrl. So that 'n' (110) -> 14
// when the Mod Ctrl, (2) is pressed.
const (
	KeyCtrlSpace rune = iota
	KeyCtrlA
	KeyCtrlB
	KeyCtrlC
	KeyCtrlD
	KeyCtrlE
	KeyCtrlF
	KeyCtrlG
	KeyCtrlH
	KeyCtrlI
	KeyCtrlJ
	KeyCtrlK
	KeyCtrlL
	KeyCtrlM
	KeyCtrlN
	KeyCtrlO
	KeyCtrlP
	KeyCtrlQ
	KeyCtrlR
	KeyCtrlS
	KeyCtrlT
	KeyCtrlU
	KeyCtrlV
	KeyCtrlW
	KeyCtrlX
	KeyCtrlY
	KeyCtrlZ
	KeyCtrlLeftSq // Escape
	KeyCtrlBackslash
	KeyCtrlRightSq
	KeyCtrlCarat
	KeyCtrlUnderscore
)
