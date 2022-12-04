package day3_test

import (
	"adventofcode/day3"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	resultPart1, resultPart2 := day3.Solve(lines)
	assert.Equal(t, 157, resultPart1)
	assert.Equal(t, 70, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day3.Solve(reader.ReadInput(3, false))
	assert.Equal(t, 7821, resultPart1)
	assert.Equal(t, 2752, resultPart2)
}
