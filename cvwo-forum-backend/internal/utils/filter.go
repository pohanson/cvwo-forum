package utils

// Returns elements in vs[] of type T such that pred returns true.
func Filter[T any](vs []T, pred func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range vs {
		if pred(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
