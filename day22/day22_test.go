package day22_test

import (
	"adventofcode/day22"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day22.Solve(reader.ReadExampleInput(22), 4)
	assert.Equal(t, 6032, resultPart1)
	assert.Equal(t, 5031, resultPart2)
}

func TestSolveTransmuted(t *testing.T) {
	lines := []string{
		"    ...#.#..",
		"    .#......",
		"    #.....#.",
		"    ........",
		"    ...#",
		"    #...",
		"    ....",
		"    ..#.",
		"..#....#",
		"........",
		".....#..",
		"........",
		"#...",
		"..#.",
		"....",
		"....",
		"",
		"10R5L5R10L4R5L5",
	}
	resultPart1, resultPart2 := day22.Solve(lines, 4)
	assert.Equal(t, 10012, resultPart1)
	assert.Equal(t, 10006, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day22.Solve(reader.ReadInput(22), 50)
	assert.Equal(t, 65368, resultPart1)
	assert.Equal(t, 156166, resultPart2)
}
