package sliceutils

func Filter[T any](data []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, data := range data {
		if filter(data) {
			result = append(result, data)
		}
	}
	return result
}
