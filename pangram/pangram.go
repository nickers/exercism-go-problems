package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	var cnt ['z' - 'a' + 1]int
	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			cnt[c-'a']++
		}
	}
	for _, c := range cnt {
		if c == 0 {
			return false
		}
	}
	return true
}
