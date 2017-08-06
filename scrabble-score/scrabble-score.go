package scrabble

import "strings"

const testVersion = 5

var points map[rune]int

func Score(word string) int {
	r := 0
	for _, l := range strings.ToUpper(word) {
		r += points[l]
	}
	return r
}

// Fill points table.
func init() {
	points = make(map[rune]int)
	src := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}
	for p, letters := range src {
		for _, l := range letters {
			points[l] = p
		}
	}
}
