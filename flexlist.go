package tui

import "image"

var _ Widget = &Flexlist{}

// Flexlist is a widget for displaying and selecting items.
type Flexlist struct {
	WidgetBase

	items    []string
	selected int
	pos      int

	onItemActivated    func(*Flexlist)
	onSelectionChanged func(*Flexlist)
}

// NewFlexlist returns a new Flexlist with no selection.
func NewFlexlist() *Flexlist {
	return &Flexlist{
		selected: -1,
	}
}

// Draw draws the Flexlist.
func (l *Flexlist) Draw(p *Painter) {
	y := 0
	x := 0
	maxW := 0
	for i, item := range l.items {
		style := "list.item"
		if i == l.selected-l.pos {
			style += ".selected"
		}
		p.WithStyle(style, func(p *Painter) {
			p.FillRect(x+1, y, len(item), 1)
			p.DrawText(x+1, y, item)
		})

		if len(item)+1 > maxW {
			maxW = len(item) + 1
		}
		y = y + 1
		if y == l.Size().Y {
			y = 0
			x = x + maxW
			maxW = 0
		}
	}
}

// SizeHint returns the recommended size for the Flexlist.
func (l *Flexlist) SizeHint() image.Point {
	var width int
	for _, item := range l.items {
		if w := stringWidth(item); w > width {
			width = w
		}
	}
	return image.Point{width, len(l.items)}
}

// OnKeyEvent handles terminal events.
func (l *Flexlist) OnKeyEvent(ev KeyEvent) {
	if !l.IsFocused() {
		return
	}

	switch ev.Key {
	case KeyUp:
		l.moveUp()
	case KeyDown:
		l.moveDown()
	case KeyLeft:
		l.moveLeft()
	case KeyRight:
		l.moveRight()
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

func (l *Flexlist) moveUp() {
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

func (l *Flexlist) moveLeft() {
	if l.selected >= l.Size().Y {
		l.selected = l.selected - l.Size().Y

		if l.selected < l.pos {
			l.pos = l.pos - l.Size().Y
		}
	}
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

func (l *Flexlist) moveRight() {
	if l.selected < len(l.items)-l.Size().Y {
		l.selected = l.selected + l.Size().Y
		if l.selected >= l.pos+len(l.items) {
			l.pos = l.pos + l.Size().Y
		}
	}
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

func (l *Flexlist) moveDown() {
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

// AddItems appends items to the end of the Flexlist.
func (l *Flexlist) AddItems(items ...string) {
	l.items = append(l.items, items...)
}

// RemoveItems clears all the items from the Flexlist.
func (l *Flexlist) RemoveItems() {
	l.items = []string{}
	l.pos = 0
	l.selected = -1
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

// RemoveItem removes the item at the given position.
func (l *Flexlist) RemoveItem(i int) {
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

// Length returns the number of items in the Flexlist.
func (l *Flexlist) Length() int {
	return len(l.items)
}

// SetSelected sets the currently selected item.
func (l *Flexlist) SetSelected(i int) {
	l.selected = i
}

// Selected returns the index of the currently selected item.
func (l *Flexlist) Selected() int {
	return l.selected
}

// Select calls SetSelected and the OnSelectionChanged function.
func (l *Flexlist) Select(i int) {
	l.SetSelected(i)
	if l.onSelectionChanged != nil {
		l.onSelectionChanged(l)
	}
}

// SelectedItem returns the currently selected item.
func (l *Flexlist) SelectedItem() string {
	return l.items[l.selected]
}

// OnItemActivated gets called when activated (through pressing KeyEnter).
func (l *Flexlist) OnItemActivated(fn func(*Flexlist)) {
	l.onItemActivated = fn
}

// OnSelectionChanged gets called whenever a new item is selected.
func (l *Flexlist) OnSelectionChanged(fn func(*Flexlist)) {
	l.onSelectionChanged = fn
}
