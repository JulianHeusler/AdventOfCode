package day15_test

import (
	"adventofcode/day15"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadExampleInput(15))
	assert.Equal(t, 26, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 12, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{6, 6}))
	assert.Equal(t, 12, day15.TaxicabDistance(day15.Position{6, 6}, day15.Position{0, 0}))
	assert.Equal(t, 0, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{0, 0}))
	assert.Equal(t, 1, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{0, 1}))
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadInput(15))
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
