package day25_test

import (
	"adventofcode/day25"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day25.Solve(reader.ReadExampleInput(25))
	assert.Equal(t, "2=-1=0", resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day25.Solve(reader.ReadInput(25))
	assert.Equal(t, "20-1-0=-2=-2220=0011", resultPart1)
	assert.Equal(t, 0, resultPart2)
}
