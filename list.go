package tui

import "image"

var _ Widget = &List{}

// List is a widget for displaying and selecting items.
type List struct {
	WidgetBase

	items    []string
	selected int
	pos      int

	onItemActivated    func(*List)
	onSelectionChanged func(*List)
}

// NewList returns a new List with no selection.
func NewList() *List {
	return &List{
		selected: -1,
	}
}

// Draw draws the list.
func (l *List) Draw(p *Painter) {
	for i, item := range l.items {
		style := "list.item"
		if i == l.selected-l.pos {
			style += ".selected"
		}
		p.WithStyle(style, func(p *Painter) {
			p.FillRect(0, i, l.Size().X, 1)
			p.DrawText(0, i, item)
		})
	}
}

// SizeHint returns the recommended size for the list.
func (l *List) SizeHint() image.Point {
	var width int
	for _, item := range l.items {
		if w := stringWidth(item); w > width {
			width = w
		}
	}
	return image.Point{width, len(l.items)}
}

func (l *List) OnEvent(ev Event) {
	if !l.IsFocused() || ev.Type != EventKey {
		return
	}

	switch ev.Key {
	case KeyArrowUp:
		l.moveUp()
	case KeyArrowDown:
		l.moveDown()
	case KeyEnter:
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
		if l.selected >= l.pos+len(l.items) {
			l.pos++
		}
	}
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

func (l *List) AddItems(items ...string) {
	l.items = append(l.items, items...)
}

func (l *List) RemoveItems() {
       l.items = []string{}
       l.pos = 0
       l.selected = -1
}

func (l *List) SetSelected(i int) {
	l.selected = i
}

func (l *List) Selected() int {
	return l.selected
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
