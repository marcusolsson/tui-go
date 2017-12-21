package tui

import (
	"image"
	"testing"
)

var TailBoxTests = []struct {
	Test string
	Setup func() Widget
	Want string
} {
	{
		Test: "draw small labels",
		Setup: func() Widget {
			return NewTailBox(
				NewLabel("hello mom"),
				NewLabel("hello dad"),
			)
		},
		Want: `
          
          
          
hello mom 
hello dad 
`,
	},
	{
		Test: "draw unwrapped labels",
		Setup: func() Widget {
			l1, l2 := NewLabel("hello muddah"), NewLabel("hello faddah")
			return NewTailBox(l1, l2)
		},
		Want: `
          
          
          
hello mudd
hello fadd
`,
	},
	{
		Test: "draw wrapped labels",
		Setup: func() Widget {
			l1, l2 := NewLabel("hello muddah"), NewLabel("hello faddah")
			l1.SetWordWrap(true)
			l2.SetWordWrap(true)
			return NewTailBox(l1, l2)
		},
		Want: `
          
hello     
muddah    
hello     
faddah    
`,
	},


}

func TestTailBox(t *testing.T) {
	for _, tt := range TailBoxTests {
		tt := tt
		t.Run(tt.Test, func(t *testing.T) {
			surface := NewTestSurface(10, 5)
			p := NewPainter(surface, NewTheme())
			p.Repaint(tt.Setup())

			if surface.String() != tt.Want {
				t.Errorf("unexpected contents: got = \n%s\nwant = \n%s", surface.String(), tt.Want)
			}
		})
	}
}


// TailBox is a container Widget that may not show all its 
// While Box attempts to show every contained Widget - sometimes shrinking
// those Widgets to do so- TailBox prioritizes completely displaying its last
// Widget, then the next-to-last widget, etc.
// It is vertically-aligned, i.e. all the contained Widgets have the same width.
type TailBox struct {
	WidgetBase
	sz image.Point
	contents []Widget
}

var _ Widget = &TailBox{}

func NewTailBox(w ...Widget) *TailBox {
	return &TailBox{
		contents: w,
	}
}

func (t *TailBox) Append(w Widget) {
	t.contents = append(t.contents, w)
	t.doLayout(t.Size())
}

func (t *TailBox) SetContents(w ...Widget) {
	t.contents = w
	t.doLayout(t.Size())
}

func (t *TailBox) Draw(p *Painter) {
	p.WithMask(image.Rect(0, 0, t.sz.X, t.sz.Y), func(p *Painter) {
		// Draw background
		p.FillRect(0, 0, t.sz.X, t.sz.Y)

		// Draw from the bottom up.
		space := t.sz.Y
		p.Translate(0, space)
		defer p.Restore()
		for i := len(t.contents) - 1; i >= 0 && space > 0; i-- {
			w := t.contents[i]
			space -= w.Size().Y
			p.Translate(0, -w.Size().Y)
			defer p.Restore()
			w.Draw(p)
		}
	})
}

// Resize recalculates the layout of the box's contents.
func (t *TailBox) Resize(size image.Point) {
	t.WidgetBase.Resize(size)
	defer func() {
		t.sz = size
	}()

	// If it's just a height change, Draw should do the right thing already.
	if size.X != t.sz.X {
		t.doLayout(size)
	}
}

func (t *TailBox) doLayout(size image.Point) {
	for _, w := range t.contents {
		hint := w.SizeHint()
		// Set the width to the container width, and height to the requested height
		w.Resize(image.Pt(size.X, hint.Y))
		// ...and then resize again, now that the Y-hint has been refreshed by the X-value.
		hint = w.SizeHint()
		w.Resize(image.Pt(size.X, hint.Y))
	}
}


