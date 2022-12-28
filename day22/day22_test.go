package day22_test

import (
	"adventofcode/day22"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	resultPart1, _ := day22.Solve(reader.ReadExampleInput(22), 4)
	assert.Equal(t, 6032, resultPart1)
}

func TestSolvePart2(t *testing.T) {
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
	_, resultPart2 := day22.Solve(lines, 4)
	assert.Equal(t, 10006, resultPart2)
}

func TestSolvePart2Example(t *testing.T) {
	_, resultPart2 := day22.Solve(reader.ReadExampleInput(22), 4)
	assert.Equal(t, 5031, resultPart2)
}

func TestSolvePart22222(t *testing.T) {
	lines := []string{
		"        1111",
		"        1111",
		"        1111",
		"        1111",
		"333322224444",
		"333322224444",
		"333322224444",
		"333322224444",
		"        66665555",
		"        66665555",
		"        66665555",
		"        66665555",
		"",
		"10R5L5R10L4R5L5",
	}
	_, resultPart2 := day22.Solve(lines, 4)
	assert.Equal(t, 5031, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day22.Solve(reader.ReadInput(22), 50)
	assert.Equal(t, 65368, resultPart1)
	assert.Equal(t, 156166, resultPart2)
}
