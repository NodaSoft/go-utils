package other

// FirstNonEmpty returns the first element with a non-empty value.
func FirstNonEmpty[T comparable](tt ...T) T {
	var nilVal T
	for _, e := range tt {
		if e != nilVal {
			return e
		}
	}

	return nilVal
}
