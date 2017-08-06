package etl

import "strings"

const testVersion = 1

// Transform old points system to new format.
func Transform(old map[int][]string) map[string]int {
	r := make(map[string]int)

	for points, letters := range old {
		for _, letter := range letters {
			r[strings.ToLower(letter)] = points
		}
	}

	return r
}
