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
	assert.Equal(t, 14848514+8504156, resultPart1)
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
	assert.Equal(t, 48381165, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day7.Solve(reader.ReadInput(5, false))
	assert.Equal(t, "DHBJQJCCW", resultPart1)
	assert.Equal(t, "WJVRLSJJT", resultPart2)
}
