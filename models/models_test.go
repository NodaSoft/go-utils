package models

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModel[T comparable] struct {
	id       uint
	property T
}

func (m *testModel[T]) GetID() uint {
	return m.id
}

func (m *testModel[T]) GetProperty() T {
	return m.property
}

func TestCollectIDs(t *testing.T) {
	sl := []*testModel[string]{{id: 3}, {id: 5}, {id: 1}}

	ids := CollectIDs(sl)
	expected := []uint{3, 5, 1}
	assert.ElementsMatch(t, expected, ids)
}

func TestCollectIDsFromMap(t *testing.T) {
	m := map[string]*testModel[string]{
		"instance1": {id: 3},
		"instance2": {id: 5},
		"instance3": {id: 1},
	}

	ids := CollectIDsFromMap(m)
	expected := []uint{3, 5, 1}
	assert.ElementsMatch(t, expected, ids)
}

func TestUniqueValues(t *testing.T) {
	sl := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 2, property: "value2"},
		{id: 3, property: "value1"}, // duplicate
	}

	properties := UniqueValues(sl, func(m *testModel[string]) string {
		return m.GetProperty()
	})
	expected := []string{"value1", "value2"}
	assert.ElementsMatch(t, expected, properties)
}

func TestUniqueValuesFromMap(t *testing.T) {
	m := map[string]*testModel[string]{
		"instance1": {id: 1, property: "value1"},
		"instance2": {id: 2, property: "value2"},
		"instance3": {id: 3, property: "value1"}, // duplicate
	}

	properties := UniqueValuesFromMap(m, func(m *testModel[string]) string {
		return m.GetProperty()
	})
	expected := []string{"value1", "value2"}
	assert.ElementsMatch(t, expected, properties)
}

func TestEntityDiff(t *testing.T) {
	a := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 2, property: "value2"},
		{id: 3, property: "value3"},
	}

	b := []*testModel[string]{
		{id: 2, property: "value2"},
	}

	expected := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 3, property: "value3"},
	}

	diff := EntityDiff(a, b)
	assert.ElementsMatch(t, expected, diff)
}

func TestEntityDiff_NoDiff(t *testing.T) {
	a := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 2, property: "value2"},
	}

	b := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 2, property: "value2"},
	}

	diff := EntityDiff(a, b)
	assert.Empty(t, diff)
}

func TestEntityDiff_EmptySecondSlice(t *testing.T) {
	a := []*testModel[string]{
		{id: 1, property: "value1"},
		{id: 2, property: "value2"},
	}

	var b []*testModel[string]

	diff := EntityDiff(a, b)
	assert.ElementsMatch(t, a, diff)
}

func TestEntityDiff_EmptyFirstSlice(t *testing.T) {
	var a []*testModel[string]
	b := []*testModel[string]{
		{id: 1, property: "value1"},
	}

	diff := EntityDiff(a, b)
	assert.Empty(t, diff)
}

func BenchmarkCollectIDs(b *testing.B) {
	sl := make([]*testModel[struct{}], 0, 1000)
	for i := 0; i < 1000; i++ {
		sl = append(sl, &testModel[struct{}]{id: uint(i)})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CollectIDs(sl)
	}
}

func BenchmarkCollectIDsFromMap(b *testing.B) {
	m := make(map[string]*testModel[struct{}], 1000)
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = &testModel[struct{}]{id: uint(i)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CollectIDsFromMap(m)
	}
}

func BenchmarkUniqueValues(b *testing.B) {
	sl := make([]*testModel[string], 1000)
	for i := 0; i < 1000; i++ {
		sl[i] = &testModel[string]{id: uint(i), property: fmt.Sprintf("value%d", i)}
		if i%10 == 0 {
			sl[i].property = "duplicateValue" // creating duplicate
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UniqueValues(sl, func(m *testModel[string]) string {
			return m.GetProperty()
		})
	}
}

func BenchmarkUniqueValuesFromMap(b *testing.B) {
	m := make(map[string]*testModel[string], 1000)
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = &testModel[string]{id: uint(i), property: fmt.Sprintf("value%d", i)}
		if i%10 == 0 {
			m[strconv.Itoa(i)].property = "duplicateValue" // creating duplicate
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UniqueValuesFromMap(m, func(m *testModel[string]) string {
			return m.GetProperty()
		})
	}
}

func BenchmarkEntityDiff(b *testing.B) {
	a := make([]*testModel[struct{}], 1000)
	bSlice := make([]*testModel[struct{}], 500)

	for i := 0; i < 1000; i++ {
		a[i] = &testModel[struct{}]{id: uint(i)}
		if i < 500 {
			bSlice[i] = &testModel[struct{}]{id: uint(i)}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EntityDiff(a, bSlice)
	}
}
