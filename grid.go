package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// Grid is a widget that lays out widgets in a grid.
type Grid struct {
	size image.Point

	rows, cols int

	rowHeights []int
	colWidths  []int

	hasBorder bool

	cells map[image.Point]Widget

	sizePolicyX SizePolicy
	sizePolicyY SizePolicy

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
			if w, ok := g.cells[pos]; ok {
				wp := g.mapCellToLocal(image.Point{i, j})
				p.Translate(wp.X, wp.Y)
				w.Draw(p)
				p.Restore()
			}
		}
	}
}

// Size returns the size of the grid.
func (g *Grid) Size() image.Point {
	return g.size
}

// SizeHint returns the recommended size for the grid.
func (g *Grid) SizeHint() image.Point {
	if g.cols == 0 || g.rows == 0 {
		return image.Point{}
	}

	var width int
	for i := 0; i < g.cols; i++ {
		width += g.columnWidth(i)
	}

	var height int
	for j := 0; j < g.rows; j++ {
		height += g.rowHeight(j)
	}

	if g.hasBorder {
		width += g.cols + 1
		height += g.rows + 1
	}

	return image.Point{width, height}
}

// SizePolicy returns the default layout behavior.
func (g *Grid) SizePolicy() (SizePolicy, SizePolicy) {
	return g.sizePolicyX, g.sizePolicyY
}

// Resize updates the size of the grid.
func (g *Grid) Resize(size image.Point) {
	hpol, vpol := g.SizePolicy()

	switch hpol {
	case Minimum:
		g.size.X = g.SizeHint().X
	case Expanding:
		g.size.X = size.X
	}

	switch vpol {
	case Minimum:
		g.size.Y = g.SizeHint().Y
	case Expanding:
		g.size.Y = size.Y
	}

	inner := g.size

	if g.hasBorder {
		inner.X = g.size.X - (g.cols + 1)
		inner.Y = g.size.Y - (g.rows + 1)
	}

	g.rowHeights = g.distributeRowHeight(inner)
	g.colWidths = g.distributeColumnWidth(inner)

	// Resize children.
	for pos, w := range g.cells {
		w.Resize(image.Point{g.colWidths[pos.X], g.rowHeights[pos.Y]})
	}
}

func (g *Grid) distributeColumnWidth(available image.Point) []int {
	columns := make([]int, g.cols)

	// Distribute minimum space.
	for i := 0; i < g.cols; i++ {
		columns[i] = g.columnWidth(i)
	}

	var used int
	for _, w := range columns {
		used += w
	}

	// Distribute remaining space (if any).
	extra := available.X - used

L:
	for extra > 0 {
		starting := extra
		for i, w := range columns {
			if s, ok := g.columnStretch[i]; ok && s > 0 {
				if extra > s {
					columns[i] = w + s
					extra -= s
				} else {
					columns[i] = w + extra
					extra -= extra
				}
				if extra == 0 {
					break L
				}
			}
		}
		if starting == extra {
			break L
		}
	}

	return columns
}

func (g *Grid) distributeRowHeight(available image.Point) []int {
	rows := make([]int, g.rows)

	// Distribute minimum space.
	for i := 0; i < g.rows; i++ {
		rows[i] = g.rowHeight(i)
	}

	var used int
	for _, h := range rows {
		used += h
	}

	// Distribute remaining space (if any).
	extra := available.Y - used

L:
	for extra > 0 {
		starting := extra
		for i, h := range rows {
			if s, ok := g.rowStretch[i]; ok && s > 0 {
				if extra > s {
					rows[i] = h + s
					extra -= s
				} else {
					rows[i] = h + extra
					extra -= extra
				}
				if extra == 0 {
					break L
				}
			}
		}
		if starting == extra {
			break L
		}
	}

	return rows
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

func (b *Grid) rowHeight(i int) int {
	result := 0
	for pos, w := range b.cells {
		if pos.Y != i {
			continue
		}

		if w.SizeHint().Y > result {
			result = w.SizeHint().Y
		}
	}

	return result
}

func (b *Grid) columnWidth(i int) int {
	result := 0
	for pos, w := range b.cells {
		if pos.X != i {
			continue
		}

		if w.SizeHint().X > result {
			result = w.SizeHint().X
		}
	}
	return result
}

func (g *Grid) OnEvent(ev termbox.Event) {
	for _, w := range g.cells {
		w.OnEvent(ev)
	}
}

func (g *Grid) IsVisible() bool {
	return true
}

func (g *Grid) SetCell(pos image.Point, w Widget) {
	g.cells[pos] = w
}

func (g *Grid) SetBorder(enabled bool) {
	g.hasBorder = enabled
}

func (g *Grid) SetSizePolicy(horizontal, vertical SizePolicy) {
	g.sizePolicyX = horizontal
	g.sizePolicyY = vertical
}

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

func (g *Grid) SetColumnStretch(col, stretch int) {
	g.columnStretch[col] = stretch
}

func (g *Grid) SetRowStretch(row, stretch int) {
	g.rowStretch[row] = stretch
}
