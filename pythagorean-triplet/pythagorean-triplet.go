package pythagorean

import "math"

// import "fmt"

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (res []Triplet) {
	maxC := int(math.Ceil(float64(max) / math.Sqrt(2)))
	// fmt.Println(">>", maxC, float64(max)/math.Sqrt(2))
	for a := min; a <= maxC; a++ {
		for b := a; b <= maxC+1; b++ {
			cs := int(math.Sqrt(float64(a*a + b*b)))
			if cs <= max && cs*cs == a*a+b*b {
				res = append(res, Triplet{a, b, cs})
			}
		}
	}
	return
}

func Sum(p int) (res []Triplet) {
	for _, t := range Range(1, p*2/3) {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return
}
