package series

const testVersion = 2

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	r := make([]string, len(s)-n+1)
	for i := range r {
		r[i] = s[i : i+n]
	}
	return r
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
