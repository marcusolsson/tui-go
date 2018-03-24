package tui

const (
	RuneMinus      = '-'
	RunePlusSmall  = '+'
	RuneVLineSmall = '|'

	RunePlus     = '┼'
	RuneHLine    = '─'
	RuneVLine    = '│'
	RuneTTee     = '┬'
	RuneRTee     = '┤'
	RuneLTee     = '├'
	RuneBTee     = '┴'
	RuneULCorner = '┌'
	RuneURCorner = '┐'
	RuneLLCorner = '└'
	RuneLRCorner = '┘'
)

var runeFallbacks = map[rune]rune{}

func init() {
	DefaultBorder()
}

// default
//	┌──────────────────┐
//	│┌────────┬───────┐│
//	││A       │apple  ││
//	│└────────┴───────┘│
//	└──────────────────┘
func DefaultBorder() {
	runeFallbacks = map[rune]rune{
		RunePlus:     RunePlus,
		RuneHLine:    RuneHLine,
		RuneVLine:    RuneVLine,
		RuneTTee:     RuneTTee,
		RuneRTee:     RuneRTee,
		RuneLTee:     RuneLTee,
		RuneBTee:     RuneBTee,
		RuneULCorner: RuneULCorner,
		RuneURCorner: RuneURCorner,
		RuneLLCorner: RuneLLCorner,
		RuneLRCorner: RuneLRCorner,
	}
}

// small
//	+------------------+
//	|+--------+-------+|
//	||A       |apple  ||
//	|+--------+-------+|
//	+------------------+
func SimpleBorder() {
	runeFallbacks = map[rune]rune{
		RuneHLine: RuneMinus,
		RuneVLine: RuneVLineSmall,

		RunePlus:     RunePlusSmall,
		RuneLLCorner: RunePlusSmall,
		RuneLRCorner: RunePlusSmall,
		RuneTTee:     RunePlusSmall,
		RuneRTee:     RunePlusSmall,
		RuneLTee:     RunePlusSmall,
		RuneBTee:     RunePlusSmall,
		RuneULCorner: RunePlusSmall,
		RuneURCorner: RunePlusSmall,
	}
}

func GetBorder(k rune) rune {
	return runeFallbacks[k]
}
