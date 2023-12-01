package day14_test

import (
	"adventofcode/day14"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}

	resultPart1, resultPart2 := day14.Solve(lines)
	assert.Equal(t, 24, resultPart1)
	assert.Equal(t, 93, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day14.Solve(reader.ReadInput(14))
	assert.Equal(t, 696, resultPart1)
	assert.Equal(t, 23610, resultPart2)
}
