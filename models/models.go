package models

import (
	"github.com/promax/go-utils/slices"
)

type HasID interface {
	GetID() uint
}

// CollectIDs returns a slice of IDs from a slice of entities that have ID.
func CollectIDs[T HasID](sl []T) []uint {
	return UniqueValues(sl, T.GetID)
}

// CollectIDsFromMap returns a slice of IDs from a map of entities that have ID.
func CollectIDsFromMap[K comparable, T HasID](m map[K]T) []uint {
	return UniqueValuesFromMap(m, T.GetID)
}

// UniqueValues a method for assembling unique values of any model field into a slice with the desired result type.
func UniqueValues[S any, R comparable](slice []S, getter func(S) R) []R {
	values := make([]R, 0, len(slice))
	for _, v := range slice {
		values = append(values, getter(v))
	}

	return slices.Unique(values)
}

// UniqueValuesFromMap a method for assembling unique values of any model field into a slice with the desired result type.
func UniqueValuesFromMap[K comparable, V any, R comparable](m map[K]V, getter func(V) R) []R {
	values := make([]R, 0, len(m))
	for _, v := range m {
		values = append(values, getter(v))
	}

	return slices.Unique(values)
}
