package tui

import (
	"image"
)

var _ Widget = &Grid{}

// Grid is a widget that lays out widgets in a grid.
type Grid struct {
	WidgetBase

	rows, cols int

	rowHeights []int
	colWidths  []int

	hasBorder bool

	cells map[image.Point]Widget

	columnStretch map[int]int
	rowStretch    map[int]int
}

// NewGrid returns a new Grid.
func NewGrid(cols, rows int) *Grid {
	return &Grid{
		cols:          cols,
		rows:          rows,
		cells:         make(map[image.Point]Widget),
		columnStretch: make(map[int]int),
		rowStretch:    make(map[int]int),
	}
}

// Draw draws the grid.
func (g *Grid) Draw(p *Painter) {
	s := g.Size()

	if g.hasBorder {
		border := 1

		// Draw outmost border.
		p.DrawRect(0, 0, s.X, s.Y)

		// Draw column dividers.
		var coloff int
		for i := 0; i < g.cols-1; i++ {
			x := g.colWidths[i] + coloff + border
			p.DrawVerticalLine(x, 0, s.Y-1)
			p.DrawRune(x, 0, '┬')
			p.DrawRune(x, s.Y-1, '┴')
			coloff = x
		}

		// Draw row dividers.
		var rowoff int
		for j := 0; j < g.rows-1; j++ {
			y := g.rowHeights[j] + rowoff + border
			p.DrawHorizontalLine(0, s.X-1, y)
			p.DrawRune(0, y, '├')
			p.DrawRune(s.X-1, y, '┤')
			rowoff = y
		}

		// Polish the intersections.
		rowoff = 0
		for j := 0; j < g.rows-1; j++ {
			y := g.rowHeights[j] + rowoff + border
			coloff = 0
			for i := 0; i < g.cols-1; i++ {
				x := g.colWidths[i] + coloff + border
				p.DrawRune(x, y, '┼')
				coloff = x
			}
			rowoff = y
		}
	}

	// Draw cell content.
	for i := 0; i < g.cols; i++ {
		for j := 0; j < g.rows; j++ {
			pos := image.Point{i, j}
			wp := g.mapCellToLocal(pos)

			if w, ok := g.cells[pos]; ok {
				p.Translate(wp.X, wp.Y)
				p.WithMask(image.Rectangle{
					Min: image.ZP,
					Max: w.Size(),
				}, func(p *Painter) {
					w.Draw(p)
				})
				p.Restore()
			}
		}
	}
}

// MinSizeHint returns the minimum size hint for the grid.
func (g *Grid) MinSizeHint() image.Point {
	if g.cols == 0 || g.rows == 0 {
		return image.Point{}
	}

	var width int
	for i := 0; i < g.cols; i++ {
		width += g.minColumnWidth(i)
	}

	var height int
	for j := 0; j < g.rows; j++ {
		height += g.minRowHeight(j)
	}

	if g.hasBorder {
		width += g.cols + 1
		height += g.rows + 1
	}

	return image.Point{width, height}
}

// SizeHint returns the recommended size hint for the grid.
func (g *Grid) SizeHint() image.Point {
	if g.cols == 0 || g.rows == 0 {
		return image.Point{}
	}

	var maxWidth int
	for i := 0; i < g.cols; i++ {
		if w := g.columnWidth(i); w > maxWidth {
			maxWidth = w
		}
	}

	var maxHeight int
	for j := 0; j < g.rows; j++ {
		if h := g.rowHeight(j); h > maxHeight {
			maxHeight = h
		}
	}

	width := maxWidth * g.cols
	height := maxHeight * g.rows

	if g.hasBorder {
		width += g.cols + 1
		height += g.rows + 1
	}

	return image.Point{width, height}
}

// Resize recursively updates the size of the Grid and all the widgets it
// contains. This is a potentially expensive operation and should be invoked
// with restraint.
//
// Resize is called by the layout engine and is not intended to be used by end
// users.
func (g *Grid) Resize(size image.Point) {
	g.size = size
	inner := g.size
	if g.hasBorder {
		inner.X = g.size.X - (g.cols + 1)
		inner.Y = g.size.Y - (g.rows + 1)
	}
	g.layoutChildren(inner)
}

func (g *Grid) layoutChildren(size image.Point) {
	g.colWidths = g.doLayout(dim(Horizontal, size), Horizontal)
	g.rowHeights = g.doLayout(dim(Vertical, size), Vertical)

	for pos, w := range g.cells {
		w.Resize(image.Point{g.colWidths[pos.X], g.rowHeights[pos.Y]})
	}
}

func (g *Grid) doLayout(space int, a Alignment) []int {
	var sizes []int
	var stretch map[int]int

	if a == Horizontal {
		sizes = make([]int, g.cols)
		stretch = g.columnStretch
	} else if a == Vertical {
		sizes = make([]int, g.rows)
		stretch = g.rowStretch
	}

	var nonZeroStretchFactors int
	for _, s := range stretch {
		if s > 0 {
			nonZeroStretchFactors++
		}
	}

	remaining := space

	// Distribute MinSizeHint
	for {
		var changed bool
		for i, sz := range sizes {
			if stretch[i] > 0 {
				continue
			}

			ws := g.rowcol(i, a)
			var sizeHint int
			for _, w := range ws {
				s := dim(a, w.MinSizeHint())
				if sizeHint < s {
					sizeHint = s
				}
			}
			if sz < sizeHint {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Minimum
	for {
		var changed bool
		for i, sz := range sizes {
			if stretch[i] > 0 {
				continue
			}

			ws := g.rowcol(i, a)
			var sizeHint int
			for _, w := range ws {
				s := dim(a, w.SizeHint())
				p := alignedSizePolicy(a, w)
				if p == Minimum && sizeHint < s {
					sizeHint = s
				}
			}
			if sz < sizeHint {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute remaining space
	for {
		var changed bool
		if nonZeroStretchFactors == 0 {
			min := maxInt
			for _, sz := range sizes {
				if sz < min {
					min = sz
				}
			}
			for i, sz := range sizes {
				if sz == min {
					sizes[i] = sz + 1
					remaining--
					if remaining <= 0 {
						goto Resize
					}
					changed = true
				}
			}
		} else {
			for i, sz := range sizes {
				s := stretch[i]
				if s > 0 {
					if remaining-s < 0 {
						s = remaining
					}
					sizes[i] = sz + s
					remaining -= s
					if remaining <= 0 {
						goto Resize
					}
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

Resize:

	return sizes
}

func (g *Grid) rowcol(i int, a Alignment) []Widget {
	cells := make([]Widget, 0)
	for p, w := range g.cells {
		if dim(a, p) == i {
			cells = append(cells, w)
		}
	}
	return cells
}

func (g *Grid) mapCellToLocal(p image.Point) image.Point {
	var lx, ly int

	for x := 0; x < p.X; x++ {
		lx += g.colWidths[x]
	}
	for y := 0; y < p.Y; y++ {
		ly += g.rowHeights[y]
	}

	if g.hasBorder {
		lx += p.X + 1
		ly += p.Y + 1
	}

	return image.Point{lx, ly}
}

func (g *Grid) rowHeight(i int) int {
	result := 0
	for pos, w := range g.cells {
		if pos.Y != i {
			continue
		}

		if w.SizeHint().Y > result {
			result = w.SizeHint().Y
		}
	}

	return result
}

func (g *Grid) columnWidth(i int) int {
	result := 0
	for pos, w := range g.cells {
		if pos.X != i {
			continue
		}

		if w.SizeHint().X > result {
			result = w.SizeHint().X
		}
	}
	return result
}

func (g *Grid) minRowHeight(i int) int {
	result := 0
	for pos, w := range g.cells {
		if pos.Y != i {
			continue
		}

		if w.MinSizeHint().Y > result {
			result = w.MinSizeHint().Y
		}
	}

	return result
}

func (g *Grid) minColumnWidth(i int) int {
	result := 0
	for pos, w := range g.cells {
		if pos.X != i {
			continue
		}

		if w.MinSizeHint().X > result {
			result = w.MinSizeHint().X
		}
	}
	return result
}

// OnKeyEvent handles key events.
func (g *Grid) OnKeyEvent(ev KeyEvent) {
	for _, w := range g.cells {
		w.OnKeyEvent(ev)
	}
}

// SetCell sets or replaces the contents of a cell.
func (g *Grid) SetCell(pos image.Point, w Widget) {
	g.cells[pos] = w
}

// SetBorder sets whether the border is visible or not.
func (g *Grid) SetBorder(enabled bool) {
	g.hasBorder = enabled
}

// AppendRow adds a new row at the end.
func (g *Grid) AppendRow(row ...Widget) {
	g.rows++

	if len(row) > g.cols {
		g.cols = len(row)
	}

	for i, cell := range row {
		pos := image.Point{i, g.rows - 1}
		g.SetCell(pos, cell)
	}
}

// RemoveRow removes the row ( at index ) from the grid
func (g *Grid) RemoveRow(index int) {
	if index < g.rows {
		g.rows--
		for i := index; i <= g.rows; i++ {
			for j := 0; j < g.cols; j++ {
				if i == g.rows {
					delete(g.cells, image.Point{j, g.rows})
				} else {
					g.cells[image.Point{j, i}] = g.cells[image.Point{j, i + 1}]
				}
			}
		}
	}
}

// RemoveRows will remove all the rows in grid
func (g *Grid) RemoveRows() {
	g.rows = 0
	g.cells = make(map[image.Point]Widget)
}

// SetColumnStretch sets the stretch factor for a given column. If stretch > 0,
// the column will expand to fill up available space. If multiple columns have
// a stretch factor > 0, stretch determines how much space the column get in
// respect to the others. E.g. by setting SetColumnStretch(0, 1) and
// SetColumnStretch(1, 2), the second column will fill up twice as much space
// as the first one.
func (g *Grid) SetColumnStretch(col, stretch int) {
	g.columnStretch[col] = stretch
}

// SetRowStretch sets the stretch factor for a given row. For more on stretch
// factors, see SetColumnStretch.
func (g *Grid) SetRowStretch(row, stretch int) {
	g.rowStretch[row] = stretch
}
