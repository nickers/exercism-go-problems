package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram determines if given text is a word or phrase without a repeating letter.
func IsIsogram(text string) bool {
	r := make(map[rune]bool)
	for _, c := range strings.ToLower(text) {
		if !unicode.IsLetter(c) {
			continue
		}
		if r[c] {
			return false
		}
		r[c] = true
	}
	return true
}
