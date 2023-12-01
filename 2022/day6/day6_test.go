package day6_test

import (
	"adventofcode/day6"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	}

	resultPart1, resultPart2 := day6.Solve(lines)
	assert.Equal(t, 7, resultPart1)
	assert.Equal(t, 19, resultPart2)

	lines = []string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
	}

	resultPart1, resultPart2 = day6.Solve(lines)
	assert.Equal(t, 5, resultPart1)
	assert.Equal(t, 23, resultPart2)

	lines = []string{
		"nppdvjthqldpwncqszvftbrmjlhg",
	}

	resultPart1, resultPart2 = day6.Solve(lines)
	assert.Equal(t, 6, resultPart1)
	assert.Equal(t, 23, resultPart2)

	lines = []string{
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	}

	resultPart1, resultPart2 = day6.Solve(lines)
	assert.Equal(t, 10, resultPart1)
	assert.Equal(t, 29, resultPart2)

	lines = []string{
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}

	resultPart1, resultPart2 = day6.Solve(lines)
	assert.Equal(t, 11, resultPart1)
	assert.Equal(t, 26, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day6.Solve(reader.ReadInput(6))
	assert.Equal(t, 1909, resultPart1)
	assert.Equal(t, 3380, resultPart2)
}
