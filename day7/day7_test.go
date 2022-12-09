package day7_test

import (
	"adventofcode/day7"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1(t *testing.T) {
	lines := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
	}

	resultPart1, resultPart2 := day7.Solve(lines)
	assert.Equal(t, 0, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve2(t *testing.T) {
	lines := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"$ cd a",
		"$ ls",
		"100000 b.txt",
	}

	resultPart1, resultPart2 := day7.Solve(lines)
	assert.Equal(t, 200000, resultPart1)
	assert.Equal(t, 0, resultPart2)

	lines = []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"$ cd a",
		"$ ls",
		"dir b",
		"$ cd b",
		"$ ls",
		"100000 b.txt",
	}

	resultPart1, resultPart2 = day7.Solve(lines)
	assert.Equal(t, 300000, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve(t *testing.T) {
	lines := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	resultPart1, resultPart2 := day7.Solve(lines)
	assert.Equal(t, 95437, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day7.Solve(reader.ReadInput(7, false))
	assert.Equal(t, 1447046, resultPart1)
	assert.Equal(t, 578710, resultPart2)
}
