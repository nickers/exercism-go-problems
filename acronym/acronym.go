package acronym

import "strings"

const testVersion = 3

func Abbreviate(longName string) string {
	var r string
	var take bool = true
	var upperLongName string = strings.ToUpper(longName)

	for i, c := range upperLongName {
		switch c >= 'A' && c <= 'Z' {
		case true:
			if take {
				take = false
				r += upperLongName[i : i+1]
			}
		default:
			take = true
		}
	}
	return r
}
