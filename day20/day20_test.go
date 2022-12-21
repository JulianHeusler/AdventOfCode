package day20_test

import (
	"adventofcode/day20"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"1",
		"2",
		"-3",
		"3",
		"-2",
		"0",
		"4",
	}

	resultPart1, resultPart2 := day20.Solve(lines)
	assert.Equal(t, 3, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day20.Solve(reader.ReadInput(20))
	assert.Equal(t, 8764, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
