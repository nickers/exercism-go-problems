package accumulate

const testVersion = 1

func Accumulate(in []string, f func(string) string) []string {
	r := make([]string, len(in))
	for idx, s := range in {
		r[idx] = f(s)
	}
	return r
}
