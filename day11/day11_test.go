package day11_test

import (
	"adventofcode/day11"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"t",
	}

	resultPart1, resultPart2 := day11.Solve(lines)
	assert.Equal(t, 95437, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day11.Solve(reader.ReadInput(11, false))
	assert.Equal(t, 1447046, resultPart1)
	assert.Equal(t, 578710, resultPart2)
}
