package day5_test

import (
	"adventofcode/day5"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	resultPart1, resultPart2 := day5.Solve(lines)
	assert.Equal(t, "CMZ", resultPart1)
	assert.Equal(t, "MCD", resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day5.Solve(reader.ReadInput(5, false))
	assert.Equal(t, "DHBJQJCCW", resultPart1)
	assert.Equal(t, "WJVRLSJJT", resultPart2)
}

func TestStack(t *testing.T) {
	day5.InitStacks(2)
	assert.Equal(t, "", day5.Peek(1))
	assert.Equal(t, "", day5.Peek(2))
	day5.Push(1, "test")
	assert.Equal(t, "test", day5.Peek(1))
	assert.Equal(t, "", day5.Peek(2))
	assert.Equal(t, day5.Pop(1), "test")
	assert.Equal(t, "", day5.Peek(1))
	assert.Equal(t, "", day5.Peek(2))
}
