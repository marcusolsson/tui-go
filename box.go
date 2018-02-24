package tui

import (
	"image"
)

var _ Widget = &Box{}

// Alignment is used to set the direction in which widgets are laid out.
type Alignment int

// Available alignment options.
const (
	Horizontal Alignment = iota
	Vertical
)

// Box is a layout for placing widgets either horizontally or vertically. If
// horizontally, all widgets will have the same height. If vertically, they
// will all have the same width.
type Box struct {
	WidgetBase

	children []Widget

	border bool
	title  string

	alignment Alignment
}

// NewVBox returns a new vertically aligned Box.
func NewVBox(c ...Widget) *Box {
	return &Box{
		children:  c,
		alignment: Vertical,
	}
}

// NewHBox returns a new horizontally aligned Box.
func NewHBox(c ...Widget) *Box {
	return &Box{
		children:  c,
		alignment: Horizontal,
	}
}

// Append adds the given widget at the end of the Box.
func (b *Box) Append(w Widget) {
	b.children = append(b.children, w)
}

// Prepend adds the given widget at the start of the Box.
func (b *Box) Prepend(w Widget) {
	b.children = append([]Widget{w}, b.children...)
}

// Insert adds the widget into the Box at a given index.
func (b *Box) Insert(i int, w Widget) {
	if len(b.children) < i || i < 0 {
		return
	}

	b.children = append(b.children, nil)
	copy(b.children[i+1:], b.children[i:])
	b.children[i] = w
}

// Remove deletes the widget from the Box at a given index.
func (b *Box) Remove(i int) {
	if len(b.children) <= i || i < 0 {
		return
	}

	b.children = append(b.children[:i], b.children[i+1:]...)
}

// Length returns the number of items in the box.
func (b *Box) Length() int {
	return len(b.children)
}

// SetBorder sets whether the border is visible or not.
func (b *Box) SetBorder(enabled bool) {
	b.border = enabled
}

// SetTitle sets the title of the box.
func (b *Box) SetTitle(title string) {
	b.title = title
}

// Alignment returns the current alignment of the Box.
func (b *Box) Alignment() Alignment {
	return b.alignment
}

// IsFocused return true if one of the children is focused.
func (b *Box) IsFocused() bool {
	for _, w := range b.children {
		if w.IsFocused() {
			return true
		}
	}
	return false
}

// Draw recursively draws the widgets it contains.
func (b *Box) Draw(p *Painter) {
	style := "box"
	if b.IsFocused() {
		style += ".focused"
	}

	p.WithStyle(style, func(p *Painter) {

		sz := b.Size()

		if b.border {
			p.WithStyle(style+".border", func(p *Painter) {
				p.DrawRect(0, 0, sz.X, sz.Y)
			})
			p.WithStyle(style, func(p *Painter) {
				p.WithMask(image.Rect(0, 0, sz.X-1, 1), func(p *Painter) {
					p.DrawText(1, 0, b.title)
				})
			})
			p.FillRect(1, 1, sz.X-2, sz.Y-2)
			p.Translate(1, 1)
			defer p.Restore()
		} else {
			p.FillRect(0, 0, sz.X, sz.Y)
		}

		var off image.Point
		for _, child := range b.children {
			switch b.Alignment() {
			case Horizontal:
				p.Translate(off.X, 0)
			case Vertical:
				p.Translate(0, off.Y)
			}

			p.WithMask(image.Rectangle{
				Min: image.ZP,
				Max: child.Size(),
			}, func(p *Painter) {
				child.Draw(p)
			})

			p.Restore()

			off = off.Add(child.Size())
		}
	})
}

// MinSizeHint returns the minimum size hint for the layout.
func (b *Box) MinSizeHint() image.Point {
	var minSize image.Point

	for _, child := range b.children {
		size := child.MinSizeHint()
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

// SizeHint returns the recommended size hint for the layout.
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

// OnKeyEvent handles an event and propagates it to all children.
func (b *Box) OnKeyEvent(ev KeyEvent) {
	for _, child := range b.children {
		child.OnKeyEvent(ev)
	}
}

// Resize recursively updates the size of the Box and all the widgets it
// contains. This is a potentially expensive operation and should be invoked
// with restraint.
//
// Resize is called by the layout engine and is not intended to be used by end
// users.
func (b *Box) Resize(size image.Point) {
	b.size = size
	inner := b.size
	if b.border {
		inner = b.size.Sub(image.Point{2, 2})
	}
	b.layoutChildren(inner)
}

func (b *Box) layoutChildren(size image.Point) {
	space := doLayout(b.children, dim(b.Alignment(), size), b.Alignment())

	for i, s := range space {
		switch b.Alignment() {
		case Horizontal:
			b.children[i].Resize(image.Point{s, size.Y})
		case Vertical:
			b.children[i].Resize(image.Point{size.X, s})
		}
	}
}

func doLayout(ws []Widget, space int, a Alignment) []int {
	sizes := make([]int, len(ws))

	if len(sizes) == 0 {
		return sizes
	}

	remaining := space

	// Distribute MinSizeHint
	for {
		var changed bool
		for i, sz := range sizes {
			if sz < dim(a, ws[i].MinSizeHint()) {
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
			p := alignedSizePolicy(a, ws[i])
			if p == Minimum && sz < dim(a, ws[i].SizeHint()) {
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

	// Distribute Preferred
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Maximum) && sz < dim(a, ws[i].SizeHint()) {
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

	// Distribute Expanding
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if p == Expanding {
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
		min := maxInt
		for i, s := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Minimum) && s <= min {
				min = s
			}
		}
		var changed bool
		for i, sz := range sizes {
			if sz != min {
				continue
			}
			p := alignedSizePolicy(a, ws[i])
			if p == Preferred || p == Minimum {
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

Resize:

	return sizes
}

func dim(a Alignment, pt image.Point) int {
	if a == Horizontal {
		return pt.X
	}
	return pt.Y
}

func alignedSizePolicy(a Alignment, w Widget) SizePolicy {
	hpol, vpol := w.SizePolicy()
	if a == Horizontal {
		return hpol
	}
	return vpol
}
