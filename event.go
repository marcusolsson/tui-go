package tui

type EventType int

const (
	EventUnknown EventType = iota
	EventKey
	EventResize
	EventMouse
	EventError
	EventInterrupt
	EventRaw
	EventNone
)

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

type ModMask int16

const (
	ModShift ModMask = 1 << iota
	ModCtrl
	ModAlt
	ModMeta
	ModNone ModMask = 0
)

type Event struct {
	Type      EventType
	Key       Key
	Ch        rune
	Modifiers ModMask
}
