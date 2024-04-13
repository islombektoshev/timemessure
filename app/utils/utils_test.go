package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueWithStruck(t *testing.T) {
	type comp struct {
		a int
		b int
	}
	list := []comp{
		{a: 1, b: 1},
		{a: 2, b: 1},
		{a: 1, b: 1}, // duplicate!
		{a: 1, b: 2},
		{a: 2, b: 2},
	}
	l2 := Unique(list)
	assert.Len(t, l2, 4)
	assert.Equal(t, []comp{
		{a: 1, b: 1},
		{a: 2, b: 1},
		{a: 1, b: 2},
		{a: 2, b: 2},
	}, l2)
}

func TestFlatMap(t *testing.T) {
	x := FlatMap([][]int{{1, 2, 3}, {4, 5, 6}}, func(t []int) []int {
		return t
	})
	assert.Len(t, x, 6)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, x)
}
