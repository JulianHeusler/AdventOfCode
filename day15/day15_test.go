package day15_test

import (
	"adventofcode/day15"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadExampleInput(15))
	assert.Equal(t, 26, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadInput(15))
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
