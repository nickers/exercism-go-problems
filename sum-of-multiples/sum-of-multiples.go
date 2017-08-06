package summultiples

// import "sort"

const testVersion = 2

func SumMultiples(max int, divisors ...int) (r int) {
	for n := 1; n < max; n++ {
		for _, d := range divisors {
			if n%d == 0 {
				r += n
				break
			}
		}
	}
	return
}
