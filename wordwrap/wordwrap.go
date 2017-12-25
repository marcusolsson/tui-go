package wordwrap

import (
	"bytes"
	"unicode"

	"github.com/mattn/go-runewidth"
)

// WrapString wraps the input string by inserting newline characters. It does
// not remove whitespace, but preserves the original text.
func WrapString(s string, width int) string {
	if len(s) <= 1 {
		return s
	}

	// Output buffer. Does not include the most recent word.
	var buf bytes.Buffer
	// Trailing word.
	var word bytes.Buffer
	var wordLen int

	spaceLeft := width
	var prev rune

	for _, curr := range s {
		if curr == rune('\n') {
			// Received a newline.
			if word.Len() > spaceLeft {
				spaceLeft = width
				buf.WriteRune(curr)
			} else {
				spaceLeft = width
			}
			word.WriteTo(&buf)
			wordLen = 0
		} else if unicode.IsSpace(prev) && !unicode.IsSpace(curr) {
			// At the start of a new word.
			// Does the last word fit on this line, or the next?
			if wordLen > spaceLeft {
				spaceLeft = width - wordLen
				buf.WriteRune('\n')
			} else {
				spaceLeft -= wordLen
			}
			// fmt.Printf("42: writing %q with %d spaces remaining out of %d\n", word.String(), spaceLeft, width)
			word.WriteTo(&buf)
			wordLen = 0
		}
		word.WriteRune(curr)
		wordLen += runewidth.RuneWidth(curr)

		prev = curr
	}

	// Close out the final word.
	if wordLen > spaceLeft {
		buf.WriteRune('\n')
	}
	word.WriteTo(&buf)

	return buf.String()
}
