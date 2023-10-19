// Package diffsquares is a small library for calculating solving difference of squares problem by applying
// Gauss Summation law
package diffsquares

// SquareOfSum returns square of the sum of the first ten natural numbers based on initial value by applying
// Gauss Summation law
func SquareOfSum(n int) int {
	res := (n * (n + 1)) / 2
	return res * res
}

// SumOfSquares returns sum of the squares of the first ten natural numbers based on initial value by applying
// Square pyramidal number principle
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns difference between square of the sum and sum of the squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
