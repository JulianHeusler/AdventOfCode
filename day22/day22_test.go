package day22_test

import (
	"adventofcode/day22"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day22.Solve(reader.ReadExampleInput(22))
	assert.Equal(t, 6032, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day22.Solve(reader.ReadInput(22))
	assert.Equal(t, 65368, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
