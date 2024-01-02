package minmax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHighest(t *testing.T) {
	var (
		expectedPos int   = 2
		expectEl    int32 = 5
	)
	el, pos := getHighest([]int32{1, 2, 5, 3, 0})

	assert.Equal(t, expectEl, el)
	assert.Equal(t, expectedPos, pos)
}

func TestGetLowest(t *testing.T) {
	var (
		expectedPos int   = 0
		expectedEl  int32 = 1
	)
	el, pos := getLowest([]int32{1, 2, 3, 4, 5})
	assert.Equal(t, expectedEl, el)
	assert.Equal(t, expectedPos, pos)
}

func TestGetMaxUsingEvaluator(t *testing.T) {
	var (
		expectedPos int   = 4
		expectedEl  int32 = 5
	)
	el, pos := GetEl([]int32{1, 2, 3, 4, 5}, maxEvaluator)
	assert.Equal(t, expectedEl, el)
	assert.Equal(t, expectedPos, pos)
}

func TestGetMinUsingEvaluator(t *testing.T) {
	var (
		expectedPos int   = 4
		expectEl    int32 = 0
	)
	el, pos := GetEl([]int32{1, 2, 5, 3, 0}, minEvaluator)

	assert.Equal(t, expectEl, el)
	assert.Equal(t, expectedPos, pos)
}
