package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// VerticalBox is a layout for placing widgets vertically.
type VerticalBox struct {
	children []Widget

	horizontalSizePolicy SizePolicy
	verticalSizePolicy   SizePolicy

	border  bool
	padding int
	margin  int

	hidden bool

	size image.Point
}

// NewVerticalBox returns a new VerticalBox.
func NewVerticalBox(c ...Widget) *VerticalBox {
	return &VerticalBox{
		children: c,
	}
}

func (b *VerticalBox) Append(w Widget) {
	b.children = append(b.children, w)
}

func (b *VerticalBox) Show() {
	b.hidden = false
}

func (b *VerticalBox) Hide() {
	b.hidden = true
}

func (b *VerticalBox) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.horizontalSizePolicy = horizontal
	b.verticalSizePolicy = vertical
}

func (b *VerticalBox) SetBorder(enabled bool) {
	b.border = enabled
}

func (b *VerticalBox) SetPadding(pad int) {
	b.padding = pad
}

func (b *VerticalBox) SetMargin(margin int) {
	b.margin = margin
}

// Draw recursively draws the children it contains.
func (b *VerticalBox) Draw(p *Painter) {
	if b.hidden {
		return
	}

	sz := b.Size()

	if b.margin > 0 {
		p.Translate(b.margin, b.margin)
		defer p.Restore()
	}

	if b.border {
		p.DrawRect(0, 0, sz.X-b.margin*2, sz.Y-b.margin*2)
		p.Translate(1, 1)
		defer p.Restore()
	}

	var off int
	for _, child := range b.children {
		if !child.IsVisible() {
			continue
		}

		p.Translate(0, off)
		child.Draw(p)
		p.Restore()

		sz := child.Size()
		off += sz.Y + b.padding
	}
}

// SizeHint returns the recommended size for the layout.
func (b *VerticalBox) SizeHint() image.Point {
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
	if b.padding > 0 {
		height += b.padding * (len(b.children) - 1)
	}
	if b.margin > 0 {
		margin := b.margin * 2
		width += margin
		height += margin
	}

	return image.Point{width, height}
}

// Size returns the size of the layout.
func (b *VerticalBox) Size() image.Point {
	return b.size
}

// Resize updates the size of the layout. This will recursively resize the
// all the children in the layout.
func (b *VerticalBox) Resize(size image.Point) {
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

	b.layoutChildren(b.size)
}

func (b *VerticalBox) layoutChildren(size image.Point) {
	space := b.distributeSpace(size)

	for _, child := range b.children {
		child.Resize(space[child])
	}
}

func (b *VerticalBox) distributeSpace(available image.Point) map[Widget]image.Point {
	widgets := make(map[Widget]image.Point)

	// Distribute minimum space.
	for _, child := range b.children {
		widgets[child] = child.SizeHint()
	}

	var used image.Point
	for _, space := range widgets {
		used.Y += space.Y
	}

	// Distribute remaining space.
	extra := available.Y - used.Y
L:
	for extra > 0 {
		starting := extra
		for child, space := range widgets {
			hpol, vpol := child.SizePolicy()

			if hpol == Expanding {
				space.X = available.X
				widgets[child] = space
			}

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
func (b *VerticalBox) SizePolicy() (SizePolicy, SizePolicy) {
	return b.horizontalSizePolicy, b.verticalSizePolicy
}

func (b *VerticalBox) OnEvent(ev termbox.Event) {
	for _, child := range b.children {
		child.OnEvent(ev)
	}
}

func (b *VerticalBox) IsVisible() bool {
	return !b.hidden
}

// HorizontalBox is a layout for placing widgets horizontally.
type HorizontalBox struct {
	children []Widget

	horizontalSizePolicy SizePolicy
	verticalSizePolicy   SizePolicy

	border  bool
	padding int
	margin  int

	hidden bool

	size image.Point
}

// NewHorizontalBox returns a new HorizontalBox.
func NewHorizontalBox(c ...Widget) *HorizontalBox {
	return &HorizontalBox{
		children: c,
	}
}

// SetSizePolicy updates the size policy for the layout. This will not visible
// until next resize.
func (b *HorizontalBox) SetSizePolicy(horizontal, vertical SizePolicy) {
	b.horizontalSizePolicy = horizontal
	b.verticalSizePolicy = vertical
}

func (b *HorizontalBox) SetBorder(border bool) {
	b.border = border
}

func (b *HorizontalBox) SetPadding(pad int) {
	b.padding = pad
}

// Draw recursively draws the children it contains.
func (b *HorizontalBox) Draw(p *Painter) {
	sz := b.Size()

	if b.margin > 0 {
		p.Translate(b.margin, b.margin)
		defer p.Restore()
	}

	if b.border {
		p.DrawRect(0, 0, sz.X-b.margin*2, sz.Y-b.margin*2)
		p.Translate(1, 1)
		defer p.Restore()
	}

	var off int
	for _, child := range b.children {
		if !child.IsVisible() {
			continue
		}

		p.Translate(off, 0)
		child.Draw(p)
		p.Restore()

		sz := child.Size()
		off += sz.X + b.padding
	}
}

// SizeHint returns the recommended size for the layout.
func (b *HorizontalBox) SizeHint() image.Point {
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
	if b.padding > 0 {
		pad := b.padding * (len(b.children) - 1)
		width += pad
	}
	if b.margin > 0 {
		margin := b.margin * 2
		width += margin
		height += margin
	}

	return image.Point{width, height}
}

// Size returns the size of the layout.
func (b *HorizontalBox) Size() image.Point {
	return b.size
}

// Resize updates the size of the layout. This will recursively resize the
// all the children in the layout.
func (b *HorizontalBox) Resize(size image.Point) {
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

	b.layoutChildren(size)
}

func (b *HorizontalBox) layoutChildren(size image.Point) {
	space := b.distributeSpace(size)

	for _, child := range b.children {
		child.Resize(space[child])
	}
}

func (b *HorizontalBox) distributeSpace(available image.Point) map[Widget]image.Point {
	widgets := make(map[Widget]image.Point)

	// Distribute minimum space.
	for _, child := range b.children {
		widgets[child] = child.SizeHint()
	}

	var used image.Point
	for _, space := range widgets {
		used.X += space.X
	}

	// Distribute remaining space.
	extra := available.X - used.X
L:
	for extra > 0 {
		starting := extra
		for child, space := range widgets {
			hpol, vpol := child.SizePolicy()

			if vpol == Expanding {
				space.Y = available.Y
				widgets[child] = space
			}

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
func (b *HorizontalBox) SizePolicy() (SizePolicy, SizePolicy) {
	return b.horizontalSizePolicy, b.verticalSizePolicy
}

func (b *HorizontalBox) OnEvent(ev termbox.Event) {
	for _, child := range b.children {
		child.OnEvent(ev)
	}
}

func (b *HorizontalBox) IsVisible() bool {
	return !b.hidden
}

func (b *HorizontalBox) Show() {
	b.hidden = false
}

func (b *HorizontalBox) Hide() {
	b.hidden = true
}
