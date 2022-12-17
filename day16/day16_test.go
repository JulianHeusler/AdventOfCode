package day16_test

import (
	"adventofcode/day16"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day16.Solve(reader.ReadExampleInput(16))
	assert.Equal(t, 1651, resultPart1)
	assert.Equal(t, 1707, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day16.Solve(reader.ReadInput(16))
	assert.Equal(t, 2029, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
