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
	KeyEsc
	KeyBackspace
	KeyBackspace2
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
)

type Event struct {
	Type EventType
	Key  Key
	Ch   rune
}
