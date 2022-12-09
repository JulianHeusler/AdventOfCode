package day9_test

import (
	"adventofcode/day9"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	resultPart1, resultPart2 := day9.Solve(lines)
	assert.Equal(t, 13, resultPart1)
	assert.Equal(t, 1, resultPart2)
}

func TestSolvePart2(t *testing.T) {
	lines := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	_, resultPart2 := day9.Solve(lines)
	assert.Equal(t, 36, resultPart2)
}

func TestTouching(t *testing.T) {
	assert.True(t, day9.IsNotTouching(day9.Position{1, 0}, day9.Position{0, 0}))

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			assert.True(t, day9.IsNotTouching(day9.Position{x, y}, day9.Position{0, 0}))
			assert.True(t, day9.IsNotTouching(day9.Position{0, 0}, day9.Position{x, y}))
		}
	}

	for i := 2; i <= 4; i++ {
		assert.False(t, day9.IsNotTouching(day9.Position{i, 0}, day9.Position{0, 0}))
		assert.False(t, day9.IsNotTouching(day9.Position{0, i}, day9.Position{0, 0}))
	}
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day9.Solve(reader.ReadInput(9, false))
	assert.Equal(t, 6745, resultPart1)
	assert.Equal(t, 2793, resultPart2)
}
