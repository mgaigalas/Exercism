// Package reverse is a library containing tools for reversing strings
package reverse

// Reverse returns reversed input string
func Reverse(input string) string {
	var res string
	for _, val := range input {
		res = string(val) + res
	}
	return res
}
