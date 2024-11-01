package slices

import (
	"strings"

	"github.com/nodasoft/go-utils/generics"
)

// ConvertSlice change type of slice elements
// example: newSlice := slices.ConvertSlice[int32, uint]([]int32{1,2,3})
func ConvertSlice[T, R generics.Numeric](s []T) []R {
	r := make([]R, 0, len(s))
	for _, i := range s {
		r = append(r, R(i))
	}

	return r
}

// FilterNil return slice without empty values - 0, "", etc.
// be careful - modifies original slice
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

// Unique return slice without duplicates
// be careful - modifies original slice
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

// Union two slices exclude duplicates.
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

// Cross returns a slice with values present in both slices
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

// IsEqual checks that slices are identical regardless of the order of elements
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

// Has checks that the slice contains the given value
func Has[T comparable](sl []T, n T) bool {
	for _, v := range sl {
		if v == n {
			return true
		}
	}

	return false
}

// TrimStrings trim every string in slice
func TrimStrings(ss []string) []string {
	for i := range ss {
		ss[i] = strings.TrimSpace(ss[i])
	}

	return ss
}

// ToKeyMap return map with key eq values of slice
func ToKeyMap[T comparable](sl []T) map[T]bool {
	keyMap := make(map[T]bool, len(sl))

	for _, val := range sl {
		keyMap[val] = true
	}

	return keyMap
}

// SliceDiff returns a slice that contains elements present in the first slice but absent from the others
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

// SliceIntersect returns a slice with unique values present in both slices
func SliceIntersect[T comparable](slices ...[]T) []T {
	tmpValuesMap := make(map[int]map[T]struct{}, len(slices))

	for i := 0; i < len(slices); i++ {
		sl := slices[i]
		if len(sl) == 0 {
			return make([]T, 0) // попался пустой слайс - дальнейшая обработка бессмысленна
		}

		tmpValuesMap[i] = make(map[T]struct{}, len(sl))
		for _, el := range sl {
			tmpValuesMap[i][el] = struct{}{}
		}
	}

	result := make([]T, 0, (len(slices[0])+len(slices[1]))/2) // len - middle len of two slices

	// range all slices
	for _, sl := range slices {
		// range all elements
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

// Max return maximum value from presented
func Max[T generics.Numeric](n []T) T {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > m {
			m = n[i]
		}
	}

	return m
}

// Min return minimal value from presented
func Min[T generics.Numeric](n []T) T {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}

	return m
}

// Sum return sum of all values
func Sum[T generics.Numeric](n []T) T {
	var sum T
	for _, v := range n {
		sum += v
	}

	return sum
}
