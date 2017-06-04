package tui

import (
	"fmt"
	"image"
	"strings"
)

type ModMask int16

const (
	ModShift ModMask = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
	ModNone ModMask = 0
)

type KeyEvent struct {
	Key       Key
	Rune      rune
	Modifiers ModMask
}

func (ev *KeyEvent) Name() string {
	s := ""
	m := []string{}
	if ev.Modifiers&ModShift != 0 {
		m = append(m, "Shift")
	}
	if ev.Modifiers&ModAlt != 0 {
		m = append(m, "Alt")
	}
	if ev.Modifiers&ModMeta != 0 {
		m = append(m, "Meta")
	}
	if ev.Modifiers&ModCtrl != 0 {
		m = append(m, "Ctrl")
	}

	ok := false
	if s, ok = keyNames[ev.Key]; !ok {
		if ev.Key == KeyRune {
			s = string(ev.Rune)
		} else {
			s = "Unknown"
		}
	}
	if len(m) != 0 {
		if ev.Modifiers&ModCtrl != 0 && strings.HasPrefix(s, "Ctrl-") {
			s = s[5:]
		}
		return fmt.Sprintf("%s+%s", strings.Join(m, "+"), s)
	}
	return s
}

type Key int16

const (
	KeyRune Key = iota + 256
	KeyUp
	KeyDown
	KeyRight
	KeyLeft
	KeyUpLeft
	KeyUpRight
	KeyDownLeft
	KeyDownRight
	KeyCenter
	KeyPgUp
	KeyPgDn
	KeyHome
	KeyEnd
	KeyInsert
	KeyDelete
	KeyHelp
	KeyExit
	KeyClear
	KeyCancel
	KeyPrint
	KeyPause
	KeyBacktab
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyF26
	KeyF27
	KeyF28
	KeyF29
	KeyF30
	KeyF31
	KeyF32
	KeyF33
	KeyF34
	KeyF35
	KeyF36
	KeyF37
	KeyF38
	KeyF39
	KeyF40
	KeyF41
	KeyF42
	KeyF43
	KeyF44
	KeyF45
	KeyF46
	KeyF47
	KeyF48
	KeyF49
	KeyF50
	KeyF51
	KeyF52
	KeyF53
	KeyF54
	KeyF55
	KeyF56
	KeyF57
	KeyF58
	KeyF59
	KeyF60
	KeyF61
	KeyF62
	KeyF63
	KeyF64
)

const (
	KeyCtrlSpace Key = iota
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

const (
	KeyNUL Key = iota
	KeySOH
	KeySTX
	KeyETX
	KeyEOT
	KeyENQ
	KeyACK
	KeyBEL
	KeyBS
	KeyTAB
	KeyLF
	KeyVT
	KeyFF
	KeyCR
	KeySO
	KeySI
	KeyDLE
	KeyDC1
	KeyDC2
	KeyDC3
	KeyDC4
	KeyNAK
	KeySYN
	KeyETB
	KeyCAN
	KeyEM
	KeySUB
	KeyESC
	KeyFS
	KeyGS
	KeyRS
	KeyUS
	KeyDEL Key = 0x7F
)

const (
	KeyBackspace  = KeyBS
	KeyTab        = KeyTAB
	KeyEsc        = KeyESC
	KeyEscape     = KeyESC
	KeyEnter      = KeyCR
	KeyBackspace2 = KeyDEL
)

var keyNames = map[Key]string{
	KeyEnter:      "Enter",
	KeyBackspace:  "Backspace",
	KeyTab:        "Tab",
	KeyBacktab:    "Backtab",
	KeyEsc:        "Esc",
	KeyBackspace2: "Backspace2",
	KeyDelete:     "Delete",
	KeyInsert:     "Insert",
	KeyUp:         "Up",
	KeyDown:       "Down",
	KeyLeft:       "Left",
	KeyRight:      "Right",
	KeyCtrlSpace:  "Ctrl-Space",
	KeyCtrlA:      "Ctrl-A",
	KeyCtrlB:      "Ctrl-B",
	KeyCtrlC:      "Ctrl-C",
	KeyCtrlD:      "Ctrl-D",
	KeyCtrlE:      "Ctrl-E",
	KeyCtrlF:      "Ctrl-F",
	KeyCtrlG:      "Ctrl-G",
	KeyCtrlJ:      "Ctrl-J",
	KeyCtrlK:      "Ctrl-K",
	KeyCtrlL:      "Ctrl-L",
	KeyCtrlN:      "Ctrl-N",
	KeyCtrlO:      "Ctrl-O",
	KeyCtrlP:      "Ctrl-P",
	KeyCtrlQ:      "Ctrl-Q",
	KeyCtrlR:      "Ctrl-R",
	KeyCtrlS:      "Ctrl-S",
	KeyCtrlT:      "Ctrl-T",
	KeyCtrlU:      "Ctrl-U",
	KeyCtrlV:      "Ctrl-V",
	KeyCtrlW:      "Ctrl-W",
	KeyCtrlX:      "Ctrl-X",
	KeyCtrlY:      "Ctrl-Y",
	KeyCtrlZ:      "Ctrl-Z",
}

type MouseEvent struct {
	Pos image.Point
}

type PaintEvent struct{}

type Event interface{}
