package diffsquares

import (
	"math"
)

// SquareOfSum calculates the square of the sum of the first N natural
// numbers by using the following formula:
// (n(n+1)/2)^2
func SquareOfSum(n int) int {
	return int(math.Pow(float64((n * (n + 1) / 2)), 2))
}

// SumOfSquares calculates the sum of the squares of the first N natural
// numbers by using the following formula:
// n^3/3 + n^2/2 + n/6
func SumOfSquares(n int) int {
	return int(math.Pow(float64(n), 3)/3 + math.Pow(float64(n), 2)/2 + float64(n)/6)
}

// Difference calculates the difference between the square of the sum of the
// first N natural numbers and the sum of the squares of the first N natural
// numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
