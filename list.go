package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// List is a widget for displaying and selecting items.
type List struct {
	size image.Point

	items    []string
	selected int
	numRows  int
	pos      int

	onItemActivated    func(*List)
	onSelectionChanged func(*List)

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy
}

// NewList returns a new List with no selection and a default number of rows.
func NewList() *List {
	return &List{
		selected: -1,
		numRows:  5,
	}
}

// Draw draws the list.
func (l *List) Draw(p *Painter) {
	sz := l.Size()

	start := clip(l.pos, 0, len(l.items)-1)
	end := clip(l.pos+l.numRows, start, len(l.items))

	for i, item := range l.items[start:end] {
		tsel := l.selected - l.pos
		if i == tsel {
			p.SetBrush(termbox.ColorBlack, termbox.ColorWhite)
			p.FillRect(0, tsel, sz.X-2, 1)
		}
		p.DrawText(0, i, item)
		if i == tsel {
			p.SetBrush(termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func clip(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

// Size returns the size of the list.
func (l *List) Size() image.Point {
	return l.size
}

// SizeHint returns the recommended size for the list.
func (l *List) SizeHint() image.Point {
	var width int
	for _, item := range l.items {
		if len(item) > width {
			width = len(item)
		}
	}
	height := l.numRows
	return image.Point{width, height}
}

// SizePolicy returns the default layout behavior.
func (l *List) SizePolicy() (SizePolicy, SizePolicy) {
	return l.sizePolicyX, l.sizePolicyY
}

// Resize updates the size of the list.
func (l *List) Resize(size image.Point) {
	hpol, vpol := l.SizePolicy()

	switch hpol {
	case Minimum:
		l.size.X = l.SizeHint().X
	case Expanding:
		l.size.X = size.X
	}

	switch vpol {
	case Minimum:
		l.size.Y = l.SizeHint().Y
	case Expanding:
		l.size.Y = size.Y
	}
}

func (l *List) OnEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		l.moveUp()
	case termbox.KeyArrowDown:
		l.moveDown()
	case termbox.KeyEnter:
		if l.onItemActivated != nil {
			l.onItemActivated(l)
		}
	}

	switch ev.Ch {
	case 'k':
		l.moveUp()
	case 'j':
		l.moveDown()
	}
}

func (l *List) moveUp() {
	if l.selected > 0 {
		l.selected--

		if l.selected < l.pos {
			l.pos--
		}
	}
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

func (l *List) moveDown() {
	if l.selected < len(l.items)-1 {
		l.selected++
		if l.selected >= l.pos+l.numRows {
			l.pos++
		}
	}
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

func (l *List) IsVisible() bool {
	return true
}

func (l *List) SetSizePolicy(h, v SizePolicy) {
	l.sizePolicyX = h
	l.sizePolicyY = v
}

func (l *List) AddItems(items ...string) {
	l.items = append(l.items, items...)
}

func (l *List) SetSelected(i int) {
	l.selected = i
}

func (l *List) Selected() int {
	return l.selected
}

func (l *List) SetRows(n int) {
	l.numRows = n
}

func (l *List) SelectedItem() string {
	return l.items[l.selected]
}

func (l *List) OnItemActivated(fn func(*List)) {
	l.onItemActivated = fn
}

func (l *List) OnSelectionChanged(fn func(*List)) {
	l.onSelectionChanged = fn
}
