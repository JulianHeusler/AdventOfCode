package day12_test

import (
	"adventofcode/day12"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}

	resultPart1, resultPart2 := day12.Solve(lines)
	assert.Equal(t, 31, resultPart1)
	assert.Equal(t, 29, resultPart2)
}

func TestSolve2(t *testing.T) {
	lines := []string{
		"Sabcdefg",
		"onmlkjih",
		"pqrstuvw",
		"azzzzEyx",
	}

	resultPart1, resultPart2 := day12.Solve(lines)
	assert.Equal(t, 26, resultPart1)
	assert.Equal(t, 25, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day12.Solve(reader.ReadInput(12))
	assert.Equal(t, 361, resultPart1)
	assert.Equal(t, 354, resultPart2)
}
