package tui

import "image"

var _ Widget = &Entry{}

// Entry is a one-line text editor.
type Entry struct {
	text string

	size image.Point

	focused bool

	onChange func(*Entry)
	onSubmit func(*Entry)

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

// NewEntry returns a new Entry.
func NewEntry() *Entry {
	return &Entry{}
}

// Draw draws the entry.
func (e *Entry) Draw(p *Painter) {
	s := e.Size()

	style := "entry"
	if e.focused {
		style += ".focused"
	}

	p.WithStyledBrush(style, func(p *Painter) {
		p.FillRect(0, 0, s.X, 1)
		p.DrawText(0, 0, e.text)

		if e.focused {
			p.DrawCursor(stringWidth(e.text), 0)
		}
	})
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
	return e.sizePolicyX, e.sizePolicyY
}

// Resize updates the size of the entry.
func (e *Entry) Resize(size image.Point) {
	hpol, vpol := e.SizePolicy()

	switch hpol {
	case Minimum:
		e.size.X = e.SizeHint().X
	case Expanding:
		e.size.X = size.X
	}

	switch vpol {
	case Minimum:
		e.size.Y = e.SizeHint().Y
	case Expanding:
		e.size.Y = size.Y
	}
}

func (e *Entry) OnEvent(ev Event) {
	if !e.focused {
		return
	}

	if ev.Type != EventKey {
		return
	}

	if ev.Key != 0 {
		switch ev.Key {
		case KeyEnter:
			if e.onSubmit != nil {
				e.onSubmit(e)
			}
			return
		case KeySpace:
			e.text = e.text + string(' ')
			if e.onChange != nil {
				e.onChange(e)
			}
			return
		case KeyBackspace2:
			if len(e.text) > 0 {
				e.text = trimRightLen(e.text, 1)
				if e.onChange != nil {
					e.onChange(e)
				}
			}
			return
		}
	} else {
		e.text = e.text + string(ev.Ch)
		if e.onChange != nil {
			e.onChange(e)
		}
	}
}

func (e *Entry) OnChanged(fn func(entry *Entry)) {
	e.onChange = fn
}

func (e *Entry) OnSubmit(fn func(entry *Entry)) {
	e.onSubmit = fn
}

func (e *Entry) SetText(text string) {
	e.text = text
}

func (e *Entry) Text() string {
	return e.text
}

func (e *Entry) SetSizePolicy(horizontal, vertical SizePolicy) {
	e.sizePolicyX = horizontal
	e.sizePolicyY = vertical
}

func (e *Entry) SetFocused(f bool) {
	e.focused = f
}
