package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Entry is a one-line text editor.
type Entry struct {
	text string

	size image.Point

	hidden bool

	onChange func(*Entry)
	onSubmit func(*Entry)
}

// NewEntry returns a new Entry.
func NewEntry() *Entry {
	return &Entry{}
}

// Draw draws the entry.
func (e *Entry) Draw(p *Painter) {
	p.DrawText(0, 0, e.text)
}

// Size returns the size of the entry.
func (e *Entry) Size() image.Point {
	return e.size
}

// SizeHint returns the recommended size for the entry.
func (e *Entry) SizeHint() image.Point {
	return image.Point{10, 1}
}

// SizePolicy returns the default layout behavior.
func (e *Entry) SizePolicy() (SizePolicy, SizePolicy) {
	return Minimum, Minimum
}

// Resize updates the size of the entry.
func (e *Entry) Resize(contentSize image.Point) {
	e.size = e.SizeHint()
}

func (e *Entry) OnEvent(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEnter {
			e.onSubmit(e)
			return
		}

		e.text = e.text + string(ev.Ch)
		if e.onChange != nil {
			e.onChange(e)
		}
	}
}

func (e *Entry) IsVisible() bool {
	return !e.hidden
}

func (e *Entry) OnChanged(fn func(entry *Entry)) {
	e.onChange = fn
}

func (e *Entry) OnSubmit(fn func(entry *Entry)) {
	e.onSubmit = fn
}

func (e *Entry) Show() {
	e.hidden = false
}

func (e *Entry) Hide() {
	e.hidden = true
}

func (e *Entry) SetText(text string) {
	e.text = text
}

func (e *Entry) Text() string {
	return e.text
}
