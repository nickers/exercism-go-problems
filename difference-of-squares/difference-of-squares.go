package diffsquares

const testVersion = 1

// SquareOfSums computed using simple formula
func SquareOfSums(n int) int {
	return ((n + 1) * n / 2) * ((n + 1) * n / 2)
}

// SumOfSquares computed using formula from:
// https://trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm#summation
func SumOfSquares(n int) int {
	return (2*(n*n*n) + 3*(n*n) + n) / 6
}

// Difference between "square of sums" and "sum of squares" of first N natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
