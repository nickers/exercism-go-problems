package raindrops

import "fmt"

const testVersion = 3

const (
	Pling = 1 << iota
	Plang
	Plong
)

func Convert(x int) string {
	var c int
	if x%3 == 0 {
		c |= Pling
	}
	if x%5 == 0 {
		c += Plang
	}
	if x%7 == 0 {
		c |= Plong
	}
	switch c {
	case 0:
		return fmt.Sprintf("%d", x)
	case Pling:
		return "Pling"
	case Plang:
		return "Plang"
	case Pling | Plang:
		return "PlingPlang"
	case Plong:
		return "Plong"
	case Pling | Plong:
		return "PlingPlong"
	case Plang | Plong:
		return "PlangPlong"
	case Pling | Plang | Plong:
		return "PlingPlangPlong"
	}
	return ""
}
