package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModel struct {
	id uint
}

func (m *testModel) GetID() uint {
	return m.id
}

func TestCollectIDs(t *testing.T) {
	sl := []*testModel{{id: 3}, {id: 5}, {id: 1}}

	ids := CollectIDs(sl)
	exp := []uint{3, 5, 1}
	assert.Equal(t, exp, ids)
}

func TestCollectIDsFromMap(t *testing.T) {
	m := map[string]*testModel{"first": {id: 3}, "second": {id: 5}, "third": {id: 1}}

	ids := CollectIDsFromMap(m)
	exp := []uint{3, 5, 1}
	assert.Equal(t, exp, ids)
}
