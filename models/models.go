package models

import (
	"github.com/nodasoft/go-utils/slices"
)

// HasID is an interface that requires an implementation of the GetID method.
// The GetID method should return the unique identifier of the entity as an unsigned integer.
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

// EntityDiff returns a slice that contains elements present in the first slice but absent in the others. Works for entities with ID.
func EntityDiff[T HasID](slices ...[]T) []T {
	if len(slices) == 1 {
		return slices[0]
	}

	mainSlice := slices[0]

	exclude := make(map[uint]struct{})
	for i := 1; i < len(slices); i++ {
		for _, entity := range slices[i] {
			exclude[entity.GetID()] = struct{}{}
		}
	}

	result := make([]T, 0, len(mainSlice))
	for _, entity := range mainSlice {
		if _, ok := exclude[entity.GetID()]; !ok {
			result = append(result, entity)
		}
	}

	return result
}
