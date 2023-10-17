// Package strain contains examples of using generics in GO
// for comparing items in collections.
package strain

// Keep takes slice of elements and predicate function and returns slice containing elements to keep
func Keep[T any](slice []T, predicate func(value T) bool) []T {
	var res []T
	for _, val := range slice {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

// Discard takes slice of elements and predicate function and returns slice containing elements to discard
func Discard[T any](slice []T, predicate func(value T) bool) []T {
	return Keep(slice, func(value T) bool {
		return !predicate(value)
	})
}
