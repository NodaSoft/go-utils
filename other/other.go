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

// If returns the value of then if condition is true, otherwise returns the value of otherwise.
func If[T any](condition bool, then, otherwise T) T {
	if condition {
		return then
	}
	return otherwise
}

// IfFunc returns the result of calling then if condition is true, otherwise returns the result of calling otherwise (lazy evaluation).
func IfFunc[T any](condition bool, then, otherwise func() T) T {
	if condition {
		return then()
	}
	return otherwise()
}

// IfFuncE returns the result of calling then if condition is true, otherwise returns the result of calling otherwise (lazy evaluation with error).
func IfFuncE[T any](condition bool, then, otherwise func() (T, error)) (T, error) {
	if condition {
		return then()
	}
	return otherwise()
}
