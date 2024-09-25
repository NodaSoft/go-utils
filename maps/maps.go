package maps

// Has checks if the map contains the given key.
func Has[K comparable, V any](m map[K]V, n K) bool {
	_, ok := m[n]

	return ok
}

// Merge two maps. Values from map "a" has more priority.
func Merge[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	result := make(map[K]V, len(a)+len(b))

	for k, v := range b {
		result[k] = v
	}
	for k, v := range a {
		result[k] = v
	}

	return result
}

// DiffKeys return map "a" without elements from map "b"
func DiffKeys[K comparable, V any](a map[K]V, b map[K]V) map[K]V {
	result := make(map[K]V, len(a))

	for k, v := range a {
		if _, ok := b[k]; !ok {
			result[k] = v
		}
	}

	return result
}
