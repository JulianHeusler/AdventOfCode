package day21_test

import (
	"adventofcode/day21"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day21.Solve(reader.ReadExampleInput(21))
	assert.Equal(t, 152, resultPart1)
	assert.Equal(t, 301, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day21.Solve(reader.ReadInput(21))
	assert.Equal(t, 299983725663456, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
