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

// OnKeyEvent handles terminal events.
func (l *List) OnKeyEvent(ev KeyEvent) {
	if !l.IsFocused() {
		return
	}

	switch ev.Key {
	case KeyUp:
		l.moveUp()
	case KeyDown:
		l.moveDown()
	case KeyEnter:
		if l.onItemActivated != nil {
			l.onItemActivated(l)
		}
	}

	switch ev.Rune {
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

// AddItems appends items to the end of the list.
func (l *List) AddItems(items ...string) {
	l.items = append(l.items, items...)
}

// RemoveItems clears all the items from the list.
func (l *List) RemoveItems() {
	l.items = []string{}
	l.pos = 0
	l.selected = -1
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

// RemoveItem removes the item at the given position.
func (l *List) RemoveItem(i int) {
	// Adjust pos and selected before removing.
	if l.pos >= len(l.items) {
		l.pos--
	}
	if l.selected == i {
		l.selected = -1
	} else if l.selected > i {
		l.selected--
	}

	// Copy items following i to position i.
	copy(l.items[i:], l.items[i+1:])

	// Shrink items by one.
	l.items[len(l.items)-1] = ""
	l.items = l.items[:len(l.items)-1]

	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

// Length returns the number of items in the list.
func (l *List) Length() int {
	return len(l.items)
}

// SetSelected sets the currently selected item.
func (l *List) SetSelected(i int) {
	l.selected = i
}

// Selected returns the index of the currently selected item.
func (l *List) Selected() int {
	return l.selected
}

// Select calls SetSelected and the OnSelectionChanged function.
func (l *List) Select(i int) {
	l.SetSelected(i)
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

// SelectedItem returns the currently selected item.
func (l *List) SelectedItem() string {
	return l.items[l.selected]
}

// OnItemActivated gets called when activated (through pressing KeyEnter).
func (l *List) OnItemActivated(fn func(*List)) {
	l.onItemActivated = fn
}

// OnSelectionChanged gets called whenever a new item is selected.
func (l *List) OnSelectionChanged(fn func(*List)) {
	l.onSelectionChanged = fn
}
