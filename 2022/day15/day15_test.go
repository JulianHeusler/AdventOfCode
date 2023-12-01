package day15_test

import (
	"adventofcode/day15"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2(t *testing.T) {
	lines := []string{
		"Sensor at x=0, y=0: closest beacon is at x=0, y=-5",
	}
	resultPart1, resultPart2 := day15.Solve(lines, 10, 20)
	assert.Equal(t, 26, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadExampleInput(15), 10, 20)
	assert.Equal(t, 26, resultPart1)
	assert.Equal(t, 56000011, resultPart2)
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 12, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{6, 6}))
	assert.Equal(t, 12, day15.TaxicabDistance(day15.Position{6, 6}, day15.Position{0, 0}))
	assert.Equal(t, 0, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{0, 0}))
	assert.Equal(t, 1, day15.TaxicabDistance(day15.Position{0, 0}, day15.Position{0, 1}))
	assert.Equal(t, 11, day15.TaxicabDistance(day15.Position{7, 9}, day15.Position{5, 18}))
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day15.Solve(reader.ReadInput(15), 2000000, 4000000)
	assert.Equal(t, 5256611, resultPart1)
	assert.Equal(t, 13337919186981, resultPart2)
}
