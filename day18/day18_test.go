package day18_test

import (
	"adventofcode/day18"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day18.Solve(reader.ReadExampleInput(18))
	assert.Equal(t, 64, resultPart1)
	assert.Equal(t, 58, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day18.Solve(reader.ReadInput(18))
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
