package day13_test

import (
	"adventofcode/day13"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day13.Solve(reader.ReadExampleInput(13))
	assert.Equal(t, 13, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day13.Solve(reader.ReadInput(12, false))
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
