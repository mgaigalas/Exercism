// Package collatzconjecture contains tools for solving 'Collatz Conjecture' problem
package collatzconjecture

import "errors"

// CollatzConjecture takes number and returns number of steps required to solve
// 'Collatz Conjecture' problem
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("number cannot be zero or negative number")
	}

	var stepCount int
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		stepCount++
	}
	return stepCount, nil
}
