package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Equal length required.")
	}

	var cnt int
	for i, v := range a {
		if byte(v) != b[i] {
			cnt++
		}
	}
	return cnt, nil
}
