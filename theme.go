package tui

// Color represents a color.
type Color int32

// Common colors.
const (
	ColorDefault Color = iota
	ColorBlack
	ColorWhite
	ColorRed
	ColorGreen
	ColorBlue
	ColorCyan
	ColorMagenta
	ColorYellow
)

// Decoration represents a bold/underline/etc. state
type Decoration int

// Decoration modes: Inherit from parent widget, explicitly on, or explicitly off.
const (
	DecorationInherit Decoration = iota
	DecorationOn
	DecorationOff
)

// Style determines how a cell should be painted.
// The zero value uses default from
type Style struct {
	Fg Color
	Bg Color

	Reverse   Decoration
	Bold      Decoration
	Underline Decoration
}

// mergeIn returns the receiver Style, with any changes in delta applied.
func (s Style) mergeIn(delta Style) Style {
	result := s
	if delta.Fg != ColorDefault {
		result.Fg = delta.Fg
	}
	if delta.Bg != ColorDefault {
		result.Bg = delta.Bg
	}
	if delta.Reverse != DecorationInherit {
		result.Reverse = delta.Reverse
	}
	if delta.Bold != DecorationInherit {
		result.Bold = delta.Bold
	}
	if delta.Underline != DecorationInherit {
		result.Underline = delta.Underline
	}
	return result
}

// Theme defines the styles for a set of identifiers.
type Theme struct {
	styles map[string]Style
}

// DefaultTheme is a theme with reasonable defaults.
var DefaultTheme = &Theme{
	styles: map[string]Style{
		"list.item.selected":  {Reverse: DecorationOn},
		"table.cell.selected": {Reverse: DecorationOn},
		"button.focused":      {Reverse: DecorationOn},
	},
}

// NewTheme return an empty theme.
func NewTheme() *Theme {
	return &Theme{
		styles: make(map[string]Style),
	}
}

// SetStyle sets a style for a given identifier.
func (p *Theme) SetStyle(n string, i Style) {
	p.styles[n] = i
}

// Style returns the style associated with an identifier.
// If there is no Style associated with the name, it returns a default Style.
func (p *Theme) Style(name string) Style {
	return p.styles[name]
}

// HasStyle returns whether an identifier is associated with an identifier.
func (p *Theme) HasStyle(name string) bool {
	_, ok := p.styles[name]
	return ok
}
