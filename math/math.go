package math

import (
	"github.com/nodasoft/go-utils/generics"
)

// Max return maximum value from presented
func Max[T generics.Numeric](n ...T) T {
	// TODO нужна ли проверка на пустоту передаваемого среза?
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > m {
			m = n[i]
		}
	}

	return m
}

// Min return minimal value from presented
func Min[T generics.Numeric](n ...T) T {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}

	return m
}

// Sum return sum of all values
func Sum[T generics.Numeric](n ...T) T {
	var sum T
	for _, v := range n {
		sum += v
	}

	return sum
}
