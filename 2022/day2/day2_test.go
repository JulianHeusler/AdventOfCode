package day2_test

import (
	"adventofcode/day2"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}

	resultPart1, resultPart2 := day2.Solve(lines)
	assert.Equal(t, 15, resultPart1)
	assert.Equal(t, 12, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day2.Solve(reader.ReadInput(2))
	assert.Equal(t, 14163, resultPart1)
	assert.Equal(t, 12091, resultPart2)
}
