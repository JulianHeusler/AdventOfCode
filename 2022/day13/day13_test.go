package day13_test

import (
	"adventofcode/day13"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.True(t, day13.IsDigit('1'))
	assert.True(t, day13.IsDigit('2'))
	assert.True(t, day13.IsDigit('9'))
	assert.True(t, day13.IsDigit('0'))

	assert.False(t, day13.IsDigit('['))
	assert.False(t, day13.IsDigit(']'))
	assert.False(t, day13.IsDigit(','))
	assert.False(t, day13.IsDigit(' '))
}

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day13.Solve(reader.ReadExampleInput(13))
	assert.Equal(t, 13, resultPart1)
	assert.Equal(t, 140, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day13.Solve(reader.ReadInput(13))
	assert.Equal(t, 5760, resultPart1)
	assert.Equal(t, 26670, resultPart2)
}
