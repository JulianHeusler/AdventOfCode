package day17_test

import (
	"adventofcode/day17"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>",
	}

	resultPart1, resultPart2 := day17.Solve(lines)
	assert.Equal(t, 3068, resultPart1)
	assert.Equal(t, 1514285714288, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day17.Solve(reader.ReadInput(17))
	assert.Equal(t, 3193, resultPart1)
	assert.Equal(t, 1577650429835, resultPart2)
}
