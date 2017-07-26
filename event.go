package tui

import (
	"fmt"
	"image"
	"strings"
)

// ModMask is a mask of modifier keys.
type ModMask int16

// Modifiers that can be sent with a KeyEvent or a MouseEvent.
const (
	ModShift ModMask = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
	ModNone ModMask = 0
)

// KeyEvent represents a key press.
type KeyEvent struct {
	Key       Key
	Rune      rune
	Modifiers ModMask
}

// Name returns a user-friendly description of the key press.
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

// Key represents both normal and special keys. For normal letters, KeyRune is
// used together with the Rune field in the KeyEvent.
type Key int16

// These are named keys that can be handled.
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

// These are the supported control keys.
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

// These are the defined ASCII values for key codes.
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

// These are aliases for other keys.
const (
	KeyBackspace  = KeyBS
	KeyTab        = KeyTAB
	KeyEsc        = KeyESC
	KeyEscape     = KeyESC
	KeyEnter      = KeyCR
	KeyBackspace2 = KeyDEL
)

var keyNames = map[Key]string{
	KeyEnter:          "Enter",
	KeyBackspace:      "Backspace",
	KeyTab:            "Tab",
	KeyBacktab:        "Backtab",
	KeyEsc:            "Esc",
	KeyBackspace2:     "Backspace2",
	KeyInsert:         "Insert",
	KeyDelete:         "Delete",
	KeyHelp:           "Help",
	KeyUp:             "Up",
	KeyDown:           "Down",
	KeyLeft:           "Left",
	KeyRight:          "Right",
	KeyHome:           "Home",
	KeyEnd:            "End",
	KeyUpLeft:         "UpLeft",
	KeyUpRight:        "UpRight",
	KeyDownLeft:       "DownLeft",
	KeyDownRight:      "DownRight",
	KeyCenter:         "Center",
	KeyPgDn:           "PgDn",
	KeyPgUp:           "PgUp",
	KeyClear:          "Clear",
	KeyExit:           "Exit",
	KeyCancel:         "Cancel",
	KeyPause:          "Pause",
	KeyPrint:          "Print",
	KeyF1:             "F1",
	KeyF2:             "F2",
	KeyF3:             "F3",
	KeyF4:             "F4",
	KeyF5:             "F5",
	KeyF6:             "F6",
	KeyF7:             "F7",
	KeyF8:             "F8",
	KeyF9:             "F9",
	KeyF10:            "F10",
	KeyF11:            "F11",
	KeyF12:            "F12",
	KeyF13:            "F13",
	KeyF14:            "F14",
	KeyF15:            "F15",
	KeyF16:            "F16",
	KeyF17:            "F17",
	KeyF18:            "F18",
	KeyF19:            "F19",
	KeyF20:            "F20",
	KeyF21:            "F21",
	KeyF22:            "F22",
	KeyF23:            "F23",
	KeyF24:            "F24",
	KeyF25:            "F25",
	KeyF26:            "F26",
	KeyF27:            "F27",
	KeyF28:            "F28",
	KeyF29:            "F29",
	KeyF30:            "F30",
	KeyF31:            "F31",
	KeyF32:            "F32",
	KeyF33:            "F33",
	KeyF34:            "F34",
	KeyF35:            "F35",
	KeyF36:            "F36",
	KeyF37:            "F37",
	KeyF38:            "F38",
	KeyF39:            "F39",
	KeyF40:            "F40",
	KeyF41:            "F41",
	KeyF42:            "F42",
	KeyF43:            "F43",
	KeyF44:            "F44",
	KeyF45:            "F45",
	KeyF46:            "F46",
	KeyF47:            "F47",
	KeyF48:            "F48",
	KeyF49:            "F49",
	KeyF50:            "F50",
	KeyF51:            "F51",
	KeyF52:            "F52",
	KeyF53:            "F53",
	KeyF54:            "F54",
	KeyF55:            "F55",
	KeyF56:            "F56",
	KeyF57:            "F57",
	KeyF58:            "F58",
	KeyF59:            "F59",
	KeyF60:            "F60",
	KeyF61:            "F61",
	KeyF62:            "F62",
	KeyF63:            "F63",
	KeyF64:            "F64",
	KeyCtrlUnderscore: "Ctrl-_",
	KeyCtrlRightSq:    "Ctrl-]",
	KeyCtrlBackslash:  "Ctrl-\\",
	KeyCtrlCarat:      "Ctrl-^",
	KeyCtrlSpace:      "Ctrl-Space",
	KeyCtrlA:          "Ctrl-A",
	KeyCtrlB:          "Ctrl-B",
	KeyCtrlC:          "Ctrl-C",
	KeyCtrlD:          "Ctrl-D",
	KeyCtrlE:          "Ctrl-E",
	KeyCtrlF:          "Ctrl-F",
	KeyCtrlG:          "Ctrl-G",
	KeyCtrlJ:          "Ctrl-J",
	KeyCtrlK:          "Ctrl-K",
	KeyCtrlL:          "Ctrl-L",
	KeyCtrlN:          "Ctrl-N",
	KeyCtrlO:          "Ctrl-O",
	KeyCtrlP:          "Ctrl-P",
	KeyCtrlQ:          "Ctrl-Q",
	KeyCtrlR:          "Ctrl-R",
	KeyCtrlS:          "Ctrl-S",
	KeyCtrlT:          "Ctrl-T",
	KeyCtrlU:          "Ctrl-U",
	KeyCtrlV:          "Ctrl-V",
	KeyCtrlW:          "Ctrl-W",
	KeyCtrlX:          "Ctrl-X",
	KeyCtrlY:          "Ctrl-Y",
	KeyCtrlZ:          "Ctrl-Z",
}

// MouseEvent represents the event where a mouse button was pressed or
// released.
type MouseEvent struct {
	Pos image.Point
}

type paintEvent struct{}

// callbackEvent holds a user-defined function which has been submitted
// to be called on the render thread.
type callbackEvent struct {
	cbFn func()
}

type event interface{}
