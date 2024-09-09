package models

import (
	"github.com/promax/go-utils/slices"
)

type HasID interface {
	GetID() uint
}

// CollectIDs возвращает слайс ид из слайса сущностей, имеющих ид.
func CollectIDs[T HasID](sl []T) []uint {
	return UniqueValues(sl, T.GetID)
}

// CollectIDsFromMap возвращает слайс ид из слайса сущностей, имеющих ид.
func CollectIDsFromMap[K comparable, T HasID](m map[K]T) []uint {
	return UniqueValuesFromMap(m, T.GetID)
}

// UniqueValues a method for assembling unique values of any model field into a slice with the desired result type.
func UniqueValues[S interface{}, R comparable](slice []S, getter func(S) R) []R {
	ids := make([]R, 0, len(slice))
	for _, v := range slice {
		ids = append(ids, getter(v))
	}

	return slices.Unique(ids)
}

// UniqueValuesFromMap a method for assembling unique values of any model field into a slice with the desired result type.
func UniqueValuesFromMap[K comparable, V any, R comparable](m map[K]V, getter func(V) R) []R {
	ids := make([]R, 0, len(m))
	for _, v := range m {
		ids = append(ids, getter(v))
	}

	return slices.Unique(ids)
}
