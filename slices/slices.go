package slices

import (
	"strings"

	"github.com/nodasoft/go-utils/generics"
)

// ConvertSlice changes the type of slice elements.
// Example: newSlice := slices.ConvertSlice[int32, uint]([]int32{1,2,3})
func ConvertSlice[T, R generics.Numeric](s []T) []R {
	r := make([]R, 0, len(s))
	for _, i := range s {
		r = append(r, R(i))
	}

	return r
}

// FilterNil returns a slice without empty values - 0, "", etc.
// Be careful - modifies the original slice.
func FilterNil[T comparable](sl []T) []T {
	if len(sl) == 0 {
		return sl
	}

	var nilVar T
	var newSize int

	for i := range sl {
		if sl[i] != nilVar {
			sl[newSize] = sl[i]
			newSize++
		}
	}

	return sl[:newSize]
}

// Unique returns a slice without duplicates.
// Be careful - modifies the original slice.
func Unique[T comparable](sl []T) []T {
	if len(sl) == 0 {
		return nil
	}

	filterMap := make(map[T]struct{})
	var newSize int

	for i := range sl {
		if _, ok := filterMap[sl[i]]; !ok {
			sl[newSize] = sl[i]
			filterMap[sl[i]] = struct{}{}
			newSize++
		}
	}

	return sl[:newSize]
}

// Union merges two slices, excluding duplicates.
func Union[T comparable](sl1, sl2 []T) []T {
	uintsMap := make(map[T]struct{})

	for _, v := range sl1 {
		uintsMap[v] = struct{}{}
	}

	for _, v := range sl2 {
		uintsMap[v] = struct{}{}
	}

	result := make([]T, 0, len(uintsMap))

	for v := range uintsMap {
		result = append(result, v)
	}

	return result
}

// Cross returns a slice with values present in both slices.
func Cross[T comparable](sl1, sl2 []T) []T {
	if len(sl1) == 0 || len(sl2) == 0 {
		return nil
	}

	res := make([]T, 0, len(sl1))
	for _, v1 := range sl1 {
		for _, v2 := range sl2 {
			if v1 == v2 {
				res = append(res, v1)
				break
			}
		}
	}

	return res
}

// IsEqual checks if slices are identical regardless of the order of elements.
func IsEqual[T comparable](sl1, sl2 []T) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	tmpMap := make(map[T]uint, len(sl1))
	for _, v := range sl1 {
		tmpMap[v]++
	}
	for _, v := range sl2 {
		tmpMap[v]--
	}

	for _, v := range tmpMap {
		if v != 0 {
			return false
		}
	}

	return true
}

// Has checks if the slice contains the given value.
func Has[T comparable](sl []T, n T) bool {
	for _, v := range sl {
		if v == n {
			return true
		}
	}

	return false
}

// TrimStrings trims every string in a slice.
func TrimStrings(ss []string) []string {
	for i := range ss {
		ss[i] = strings.TrimSpace(ss[i])
	}

	return ss
}

// ToKeyMap returns a map where the slice values are the keys.
func ToKeyMap[T comparable](sl []T) map[T]bool {
	keyMap := make(map[T]bool, len(sl))

	for _, val := range sl {
		keyMap[val] = true
	}

	return keyMap
}

// SliceDiff returns a slice that contains elements present in the first slice but absent in the others.
func SliceDiff[T comparable](slices ...[]T) []T {
	if len(slices) == 1 {
		return slices[0]
	}

	mainSl := slices[0]

	excludeValues := make(map[T]struct{})
	for i := 1; i < len(slices); i++ {
		sl := slices[i]
		for _, v := range sl {
			excludeValues[v] = struct{}{}
		}
	}

	res := make([]T, 0, len(mainSl))
	for _, v := range mainSl {
		if _, ok := excludeValues[v]; !ok {
			res = append(res, v)
		}
	}

	return Unique(res)
}

// SliceIntersect returns a slice with unique values present in all provided slices.
func SliceIntersect[T comparable](slices ...[]T) []T {
	tmpValuesMap := make(map[int]map[T]struct{}, len(slices))

	for i := 0; i < len(slices); i++ {
		sl := slices[i]
		if len(sl) == 0 {
			return make([]T, 0) // empty slice encountered - further processing is pointless
		}

		tmpValuesMap[i] = make(map[T]struct{}, len(sl))
		for _, el := range sl {
			tmpValuesMap[i][el] = struct{}{}
		}
	}

	result := make([]T, 0, (len(slices[0])+len(slices[1]))/2) // len - middle length of two slices

	// iterate over all slices
	for _, sl := range slices {
		// iterate over all elements
		for _, el := range sl {
			existInAll := true
			for _, mm := range tmpValuesMap {
				if _, ok := mm[el]; !ok {
					existInAll = false
					break
				}
			}

			if existInAll {
				result = append(result, el)
			}
		}
	}

	return Unique(result)
}

// Max returns the maximum value from the provided values.
func Max[T generics.Numeric](n []T) T {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > m {
			m = n[i]
		}
	}

	return m
}

// Min returns the minimum value from the provided values.
func Min[T generics.Numeric](n []T) T {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}

	return m
}

// Sum returns the sum of all provided values.
func Sum[T generics.Numeric](n []T) T {
	var sum T
	for _, v := range n {
		sum += v
	}

	return sum
}
