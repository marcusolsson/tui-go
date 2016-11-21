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
}

// NewGrid returns a new Grid.
func NewGrid(cols, rows int) *Grid {
	return &Grid{
		cols:  cols,
		rows:  rows,
		cells: make(map[image.Point]Widget),
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

	g.rowHeights = make([]int, g.rows)
	for i := 0; i < g.rows; i++ {
		g.rowHeights[i] = g.rowHeight(i)
	}
	g.colWidths = make([]int, g.cols)
	for i := 0; i < g.cols; i++ {
		g.colWidths[i] = g.columnWidth(i)
	}

	// Expand the last column.
	if len(g.colWidths) > 0 {
		var colsum int
		for _, w := range g.colWidths[:len(g.colWidths)] {
			colsum += w
		}
		remaining := g.size.X - colsum
		if g.hasBorder {
			remaining -= g.cols + 1
		}
		g.colWidths[len(g.colWidths)-1] = g.colWidths[len(g.colWidths)-1] + remaining
	}

	// Resize children.
	for pos, w := range g.cells {
		w.Resize(image.Point{g.colWidths[pos.X], g.rowHeights[pos.Y]})
	}
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
