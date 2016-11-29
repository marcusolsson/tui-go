package tui

import "image"

var _ Widget = &Box{}

type Alignment int

const (
	Horizontal Alignment = iota
	Vertical
)

// Box is a layout for placing widgets.
type Box struct {
	children []Widget

	horizontalSizePolicy SizePolicy
	verticalSizePolicy   SizePolicy

	border bool

	size      image.Point
	alignment Alignment
}

// NewVBox returns a new vertical Box.
func NewVBox(c ...Widget) *Box {
	return &Box{
		children:  c,
		alignment: Vertical,
	}
}

// NewHBox returns a new horizontal Box.
func NewHBox(c ...Widget) *Box {
	return &Box{
		children:  c,
		alignment: Horizontal,
	}
}

// Append adds a new widget to the layout.
func (b *Box) Append(w Widget) {
	b.children = append(b.children, w)
}

// SetSizePolicy sets the size policy for each axis.
func (b *Box) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.horizontalSizePolicy = horizontal
	b.verticalSizePolicy = vertical
}

// SetBorder sets whether the border is visible or not.
func (b *Box) SetBorder(enabled bool) {
	b.border = enabled
}

func (b *Box) Alignment() Alignment {
	return b.alignment
}

// Draw recursively draws the children it contains.
func (b *Box) Draw(p *Painter) {
	sz := b.Size()

	if b.border {
		p.DrawRect(0, 0, sz.X, sz.Y)
		p.Translate(1, 1)
		defer p.Restore()
	}

	var off image.Point
	for _, child := range b.children {
		switch b.Alignment() {
		case Horizontal:
			p.Translate(off.X, 0)
		case Vertical:
			p.Translate(0, off.Y)
		}

		child.Draw(p)
		p.Restore()

		sz := child.Size()
		off = off.Add(sz)
	}
}

// MinSize returns the minimum size for the layout.
func (b *Box) MinSize() image.Point {
	var minSize image.Point

	for _, child := range b.children {
		size := child.MinSize()
		if b.Alignment() == Horizontal {
			minSize.X += size.X
			if size.Y > minSize.Y {
				minSize.Y = size.Y
			}
		} else {
			minSize.Y += size.Y
			if size.X > minSize.X {
				minSize.X = size.X
			}
		}
	}

	if b.border {
		minSize = minSize.Add(image.Point{2, 2})
	}

	return minSize
}

// SizeHint returns the recommended size for the layout.
func (b *Box) SizeHint() image.Point {
	var sizeHint image.Point

	for _, child := range b.children {
		size := child.SizeHint()
		if b.Alignment() == Horizontal {
			sizeHint.X += size.X
			if size.Y > sizeHint.Y {
				sizeHint.Y = size.Y
			}
		} else {
			sizeHint.Y += size.Y
			if size.X > sizeHint.X {
				sizeHint.X = size.X
			}
		}
	}

	if b.border {
		sizeHint = sizeHint.Add(image.Point{2, 2})
	}

	return sizeHint
}

// Size returns the size of the layout.
func (b *Box) Size() image.Point {
	return b.size
}

// Resize updates the size of the layout.
func (b *Box) Resize(size image.Point) {
	switch b.horizontalSizePolicy {
	case Minimum:
		b.size.X = b.SizeHint().X
	case Expanding:
		b.size.X = size.X
	}

	switch b.verticalSizePolicy {
	case Minimum:
		b.size.Y = b.SizeHint().Y
	case Expanding:
		b.size.Y = size.Y
	}

	inner := b.size

	if b.border {
		inner = b.size.Sub(image.Point{2, 2})
	}

	b.layoutChildren(inner)
}

func (b *Box) layoutChildren(size image.Point) {
	space := b.distributeSpace(size)

	for _, child := range b.children {
		child.Resize(space[child])
	}
}

func (b *Box) distributeSpace(available image.Point) map[Widget]image.Point {
	widgets := make(map[Widget]image.Point)

	// Distribute minimum space.
	for _, child := range b.children {
		widgets[child] = child.MinSize()
	}

	var used image.Point
	for _, space := range widgets {
		used = used.Add(space)
	}

	// Expand children.
	for child, space := range widgets {
		hpol, vpol := child.SizePolicy()
		if b.Alignment() == Horizontal {
			if vpol == Expanding {
				space.Y = available.Y
				widgets[child] = space
			}
		} else {
			if hpol == Expanding {
				space.X = available.X
				widgets[child] = space
			}
		}
	}

	// Distribute remaining space (if any).
	var extra int
	if b.Alignment() == Horizontal {
		extra = available.X - used.X
	} else {
		extra = available.Y - used.Y
	}

	// Distribute preferred space
K:
	for extra > 0 {
		if b.Alignment() == Horizontal {
			starting := extra
			for child, space := range widgets {
				hint := child.SizeHint()
				if space.X < hint.X {
					space.X++
					widgets[child] = space

					extra--
					if extra == 0 {
						break K
					}
				}
			}
			if starting == extra {
				break K
			}
		} else {
			starting := extra
			for child, space := range widgets {
				hint := child.SizeHint()
				if space.Y < hint.Y {
					space.Y++
					widgets[child] = space

					extra--
					if extra == 0 {
						break K
					}
				}
			}
			if starting == extra {
				break K
			}
		}
	}

	// Distribute surplus space.
L:
	for extra > 0 {
		if b.Alignment() == Horizontal {
			starting := extra
			for child, space := range widgets {
				hpol, _ := child.SizePolicy()
				if hpol == Expanding && isSmallestWidth(space.X, widgets) {
					space.X++
					widgets[child] = space

					extra--
					if extra == 0 {
						break L
					}
				}
			}
			if starting == extra {
				break L
			}
		} else {
			starting := extra
			for child, space := range widgets {
				_, vpol := child.SizePolicy()
				if vpol == Expanding && isSmallestHeight(space.Y, widgets) {
					space.Y++
					widgets[child] = space

					extra--
					if extra == 0 {
						break L
					}
				}
			}
			if starting == extra {
				break L
			}
		}
	}

	return widgets
}

// SizePolicy returns the default layout behavior.
func (b *Box) SizePolicy() (SizePolicy, SizePolicy) {
	return b.horizontalSizePolicy, b.verticalSizePolicy
}

func (b *Box) OnEvent(ev Event) {
	for _, child := range b.children {
		child.OnEvent(ev)
	}
}

func isSmallestWidth(w int, widgets map[Widget]image.Point) bool {
	for child, space := range widgets {
		hpol, _ := child.SizePolicy()
		if hpol == Expanding && space.X < w {
			return false
		}
	}
	return true
}

func isSmallestHeight(h int, widgets map[Widget]image.Point) bool {
	for child, space := range widgets {
		_, vpol := child.SizePolicy()
		if vpol == Expanding && space.Y < h {
			return false
		}
	}
	return true
}
