package day4_test

import (
	"adventofcode/day4"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	resultPart1, resultPart2 := day4.Solve(lines)
	assert.Equal(t, 2, resultPart1)
	assert.Equal(t, 4, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day4.Solve(reader.ReadInput(4))
	assert.Equal(t, 513, resultPart1)
	assert.Equal(t, 878, resultPart2)
}
