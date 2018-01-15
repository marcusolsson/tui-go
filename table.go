package tui

import "image"

var _ Widget = &Table{}

// Table is a widget that lays out widgets in a table.
type Table struct {
	selected           int
	onItemActivated    func(*Table)
	onSelectionChanged func(*Table)

	*Grid
}

// NewTable returns a new Table.
func NewTable(cols, rows int) *Table {
	return &Table{
		Grid: NewGrid(cols, rows),
	}
}

// Draw draws the table.
func (t *Table) Draw(p *Painter) {
	s := t.Size()

	if t.hasBorder {
		border := 1

		// Draw outmost border.
		p.DrawRect(0, 0, s.X, s.Y)

		// Draw column dividers.
		var coloff int
		for i := 0; i < t.cols-1; i++ {
			x := t.colWidths[i] + coloff + border
			p.DrawVerticalLine(x, 0, s.Y-1)
			p.DrawRune(x, 0, '┬')
			p.DrawRune(x, s.Y-1, '┴')
			coloff = x
		}

		// Draw row dividers.
		var rowoff int
		for j := 0; j < t.rows-1; j++ {
			y := t.rowHeights[j] + rowoff + border
			p.DrawHorizontalLine(0, s.X-1, y)
			p.DrawRune(0, y, '├')
			p.DrawRune(s.X-1, y, '┤')
			rowoff = y
		}

		// Polish the intersections.
		rowoff = 0
		for j := 0; j < t.rows-1; j++ {
			y := t.rowHeights[j] + rowoff + border
			coloff = 0
			for i := 0; i < t.cols-1; i++ {
				x := t.colWidths[i] + coloff + border
				p.DrawRune(x, y, '┼')
				coloff = x
			}
			rowoff = y
		}
	}

	// Draw cell content.
	for i := 0; i < t.cols; i++ {
		for j := 0; j < t.rows; j++ {
			style := "table.cell"
			if j == t.selected {
				style += ".selected"
			}

			p.WithStyle(style, func(p *Painter) {
				pos := image.Point{i, j}
				wp := t.mapCellToLocal(pos)

				p.Translate(wp.X, wp.Y)
				defer p.Restore()

				if w, ok := t.cells[pos]; ok {
					size := w.Size()
					size.X = t.colWidths[i]

					p.FillRect(0, 0, size.X, size.Y)

					p.WithMask(image.Rectangle{
						Min: image.ZP,
						Max: size,
					}, func(p *Painter) {
						w.Draw(p)
					})
				}
			})
		}
	}
}

// OnKeyEvent handles an event and propagates it to all children.
func (t *Table) OnKeyEvent(ev KeyEvent) {
	if !t.IsFocused() {
		return
	}

	switch ev.Key {
	case KeyUp:
		t.moveUp()
	case KeyDown:
		t.moveDown()
	case KeyEnter:
		if t.onItemActivated != nil {
			t.onItemActivated(t)
		}
	}

	switch ev.Rune {
	case 'k':
		t.moveUp()
	case 'j':
		t.moveDown()
	}
}

func (t *Table) moveUp() {
	if t.selected > 0 {
		t.selected--
	}
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

func (t *Table) moveDown() {
	if t.selected < t.rows-1 {
		t.selected++
	}
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

// SetSelected changes the currently selected item.
func (t *Table) SetSelected(i int) {
	t.selected = i
}

// Selected returns the index of the currently selected item.
func (t *Table) Selected() int {
	return t.selected
}

// Select calls SetSelected and the OnSelectionChanged function.
func (t *Table) Select(i int) {
	t.SetSelected(i)
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

// RemoveRow removes specific row from the table
func (t *Table) RemoveRow(index int) {
	t.Grid.RemoveRow(index)
	if t.selected == index {
		t.selected = -1
	} else if t.selected > index {
		t.selected--
	}
}

// RemoveRows removes all the rows added to the table.
func (t *Table) RemoveRows() {
	t.Grid.RemoveRows()
	t.selected = -1
}

// OnItemActivated sets the function that is called when an item was activated.
func (t *Table) OnItemActivated(fn func(*Table)) {
	t.onItemActivated = fn
}

// OnSelectionChanged sets the function that is called when an item was selected.
func (t *Table) OnSelectionChanged(fn func(*Table)) {
	t.onSelectionChanged = fn
}
