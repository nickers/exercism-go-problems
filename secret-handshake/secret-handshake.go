package secret

const testVersion = 2

const reverseMask = 1 << 4

var possibleMoves = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(n uint) []string {
	var r []string

	for i, v := range possibleMoves {
		if n&(1<<uint(i)) != 0 {
			r = append(r, v)
		}
	}
	if n&reverseMask != 0 {
		reverseInPlace(r)
	}
	return r
}

func reverseInPlace(moves []string) {
	for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
		moves[i], moves[j] = moves[j], moves[i]
	}
}
