package tui

import(
	"image"
	"testing"
)

func TestIsTextRemaining(t *testing.T) {
	for _, tt := range []struct {
		text   string
		offset int
		width  int
		want   bool
	}{
		{"Lorem ipsum", 0, 11, false},
		{"Lorem ipsum", 1, 11, false},
		{"Lorem ipsum", 0, 10, true},
		{"Lorem ipsum", 5, 5, true},
	} {
		t.Run("", func(t *testing.T) {
			e := NewEntry()
			e.SetText(tt.text)
			e.SetFocused(true)
			e.Resize(image.Pt(tt.width, 1))

			e.Offset = tt.offset

			if e.isTextRemaining() != tt.want {
				t.Fatalf("want = %v; got = %v", tt.want, e.isTextRemaining())
			}
		})
	}
}

