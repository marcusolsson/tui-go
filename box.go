package tui

import "image"

var _ Widget = &VBox{}
var _ Widget = &HBox{}

// VBox is a layout for placing widgets vertically.
type VBox struct {
	children []Widget

	horizontalSizePolicy SizePolicy
	verticalSizePolicy   SizePolicy

	border bool

	size image.Point
}

// NewVBox returns a new VBox.
func NewVBox(c ...Widget) *VBox {
	return &VBox{
		children: c,
	}
}

func (b *VBox) Append(w Widget) {
	b.children = append(b.children, w)
}

func (b *VBox) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.horizontalSizePolicy = horizontal
	b.verticalSizePolicy = vertical
}

func (b *VBox) SetBorder(enabled bool) {
	b.border = enabled
}

// Draw recursively draws the children it contains.
func (b *VBox) Draw(p *Painter) {
	sz := b.Size()

	if b.border {
		p.DrawRect(0, 0, sz.X, sz.Y)
		p.Translate(1, 1)
		defer p.Restore()
	}

	var off int
	for _, child := range b.children {
		p.Translate(0, off)
		child.Draw(p)
		p.Restore()

		sz := child.Size()
		off += sz.Y
	}
}

// SizeHint returns the recommended size for the layout.
func (b *VBox) SizeHint() image.Point {
	var width, height int

	for _, child := range b.children {
		size := child.SizeHint()
		height += size.Y
		if size.X > width {
			width = size.X
		}
	}

	if b.border {
		width += 2
		height += 2
	}

	return image.Point{width, height}
}

// Size returns the size of the layout.
func (b *VBox) Size() image.Point {
	return b.size
}

// Resize updates the size of the layout. This will recursively resize the
// all the children in the layout.
func (b *VBox) Resize(size image.Point) {
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
		inner.X = b.size.X - 2
		inner.Y = b.size.Y - 2
	}

	b.layoutChildren(inner)
}

func (b *VBox) layoutChildren(size image.Point) {
	space := b.distributeSpace(size)

	for _, child := range b.children {
		child.Resize(space[child])
	}
}

func (b *VBox) distributeSpace(available image.Point) map[Widget]image.Point {
	widgets := make(map[Widget]image.Point)

	// Distribute minimum space.
	for _, child := range b.children {
		widgets[child] = child.SizeHint()
	}

	var used image.Point
	for _, space := range widgets {
		used.Y += space.Y
	}

	// Expand children horizontally.
	for child, space := range widgets {
		hpol, _ := child.SizePolicy()
		if hpol == Expanding {
			space.X = available.X
			widgets[child] = space
		}
	}

	// Distribute remaining vertical space (if any).
	extra := available.Y - used.Y
L:
	for extra > 0 {
		starting := extra
		for child, space := range widgets {
			_, vpol := child.SizePolicy()
			if vpol == Expanding {
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

	return widgets
}

// SizePolicy returns the default layout behavior.
func (b *VBox) SizePolicy() (SizePolicy, SizePolicy) {
	return b.horizontalSizePolicy, b.verticalSizePolicy
}

func (b *VBox) OnEvent(ev Event) {
	for _, child := range b.children {
		child.OnEvent(ev)
	}
}

// HBox is a layout for placing widgets horizontally.
type HBox struct {
	children []Widget

	horizontalSizePolicy SizePolicy
	verticalSizePolicy   SizePolicy

	border bool

	size image.Point
}

// NewHBox returns a new HBox.
func NewHBox(c ...Widget) *HBox {
	return &HBox{
		children: c,
	}
}

// SetSizePolicy updates the size policy for the layout. This will not visible
// until next resize.
func (b *HBox) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.horizontalSizePolicy = horizontal
	b.verticalSizePolicy = vertical
}

func (b *HBox) SetBorder(border bool) {
	b.border = border
}

// Draw recursively draws the children it contains.
func (b *HBox) Draw(p *Painter) {
	sz := b.Size()

	if b.border {
		p.DrawRect(0, 0, sz.X, sz.Y)
		p.Translate(1, 1)
		defer p.Restore()
	}

	var off int
	for _, child := range b.children {
		p.Translate(off, 0)
		child.Draw(p)
		p.Restore()

		sz := child.Size()
		off += sz.X
	}
}

// SizeHint returns the recommended size for the layout.
func (b *HBox) SizeHint() image.Point {
	var width, height int
	for _, child := range b.children {
		size := child.SizeHint()
		width += size.X
		if size.Y > height {
			height = size.Y
		}
	}

	if b.border {
		width += 2
		height += 2
	}

	return image.Point{width, height}
}

// Size returns the size of the layout.
func (b *HBox) Size() image.Point {
	return b.size
}

// Resize updates the size of the layout. This will recursively resize the
// all the children in the layout.
func (b *HBox) Resize(size image.Point) {
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
		inner.X = b.size.X - 2
		inner.Y = b.size.Y - 2
	}

	b.layoutChildren(inner)
}

func (b *HBox) layoutChildren(size image.Point) {
	space := b.distributeSpace(size)

	for _, child := range b.children {
		child.Resize(space[child])
	}
}

func (b *HBox) distributeSpace(available image.Point) map[Widget]image.Point {
	widgets := make(map[Widget]image.Point)

	// Distribute minimum space.
	for _, child := range b.children {
		widgets[child] = child.SizeHint()
	}

	var used image.Point
	for _, space := range widgets {
		used.X += space.X
	}

	// Expand children vertically.
	for child, space := range widgets {
		_, vpol := child.SizePolicy()
		if vpol == Expanding {
			space.Y = available.Y
			widgets[child] = space
		}
	}

	// Distribute remaining horizontal space (if any).
	extra := available.X - used.X
L:
	for extra > 0 {
		starting := extra
		for child, space := range widgets {
			hpol, _ := child.SizePolicy()
			if hpol == Expanding {
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
	}

	return widgets
}

// SizePolicy returns the default layout behavior.
func (b *HBox) SizePolicy() (SizePolicy, SizePolicy) {
	return b.horizontalSizePolicy, b.verticalSizePolicy
}

func (b *HBox) OnEvent(ev Event) {
	for _, child := range b.children {
		child.OnEvent(ev)
	}
}
