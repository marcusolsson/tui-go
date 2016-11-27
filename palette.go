package tui

import termbox "github.com/nsf/termbox-go"

type Color termbox.Attribute

type PaletteItem struct {
	Fg Color
	Bg Color
}

type Palette struct {
	items map[string]PaletteItem
}

var DefaultPalette = &Palette{
	items: map[string]PaletteItem{
		"normal":              {Color(termbox.ColorDefault), Color(termbox.ColorDefault)},
		"list.item.selected":  {Color(termbox.ColorWhite), Color(termbox.ColorBlue)},
		"table.cell.selected": {Color(termbox.ColorWhite), Color(termbox.ColorBlue)},
		"button.focused":      {Color(termbox.ColorWhite), Color(termbox.ColorBlue)},
	},
}

func NewPalette() *Palette {
	return &Palette{
		items: make(map[string]PaletteItem),
	}
}

func (p *Palette) SetItem(n string, i PaletteItem) {
	p.items[n] = i
}

func (p *Palette) Item(name string) PaletteItem {
	if c, ok := p.items[name]; ok {
		return c
	}
	return PaletteItem{Color(termbox.ColorDefault), Color(termbox.ColorDefault)}
}
