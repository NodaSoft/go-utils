package maps

import (
	"iter"
)

// OrderedMap represents a map whose elements are iterated over in the order they are inserted.
// Deleting elements is slow.
// To use this object, it is absolutely necessary to create the object using the NewOrderedMap constructor.
type OrderedMap[K comparable, V any] struct {
	data  map[K]V
	order []K
}

// NewOrderedMap allocates and initializes new object of type OrderedMap and returns a pointer to it.
// It is absolutely necessary to create the OrderedMap object using this constructor.
func NewOrderedMap[K comparable, V any](size int) *OrderedMap[K, V] {
	m := OrderedMap[K, V]{
		data:  make(map[K]V, size),
		order: make([]K, 0, size),
	}

	return &m
}

// Len returns the length of map.
func (m *OrderedMap[_, _]) Len() int {
	if m == nil {
		return 0
	}

	return len(m.data)
}

// Has checks if the map contains the given key.
func (m *OrderedMap[K, _]) Has(k K) bool {
	if m == nil {
		return false
	}

	return Has(m.data, k)
}

// GetValue returns a value by the given key.
// If m is nil or k is not found, returns zero value of type V.
func (m *OrderedMap[K, V]) GetValue(k K) V {
	if m == nil {
		var v V
		return v
	}

	return m.data[k]
}

// GetAndCheck like GetValue and Has methods,
// returns a value by the given key and checks if the map contains the key.
func (m *OrderedMap[K, V]) GetAndCheck(k K) (V, bool) {
	if m == nil {
		var v V
		return v, false
	}

	v, ok := m.data[k]

	return v, ok
}

// Set sets the given key-value pair into the map.
// Panics if m is not initialized by NewOrderedMap constructor.
func (m *OrderedMap[K, V]) Set(k K, v V) {
	if !m.Has(k) {
		m.order = append(m.order, k)
	}
	m.data[k] = v
}

// Delete deletes an element by the given key if the key exists.
func (m *OrderedMap[K, _]) Delete(k K) {
	if !m.Has(k) {
		return
	}

	delete(m.data, k)

	for i, found := range m.order {
		if found == k {
			m.order = append(m.order[:i], m.order[i+1:]...)
			break
		}
	}
}

// Iterate iterates over map elements in the order they are inserted.
func (m *OrderedMap[K, V]) Iterate() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m == nil {
			return
		}
		for _, k := range m.order {
			if !yield(k, m.data[k]) {
				return
			}
		}
	}
}
