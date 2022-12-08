package day8_test

import (
	"adventofcode/day8"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	resultPart1, resultPart2 := day8.Solve2(lines)
	assert.Equal(t, 21, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve2323(t *testing.T) {
	lines := []string{
		"1111",
		"1021",
		"1111",
	}

	resultPart1, resultPart2 := day8.Solve2(lines)
	assert.Equal(t, 11, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve23323(t *testing.T) {
	lines := []string{
		"123",
		"456",
		"789",
	}

	resultPart1, resultPart2 := day8.Solve2(lines)
	assert.Equal(t, 9, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolve999(t *testing.T) {
	lines := []string{
		"999",
		"999",
		"999",
		"999",
	}

	resultPart1, resultPart2 := day8.Solve2(lines)
	assert.Equal(t, 11, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestRotate(t *testing.T) {
	lines := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	expected := []string{
		"32290",
		"71349",
		"35353",
		"05535",
		"32633",
	}
	assert.Equal(t, expected, day8.RotateToLeft(lines))

	lines = []string{
		"123",
		"456",
	}
	expected = []string{
		"36",
		"25",
		"14",
	}
	assert.Equal(t, expected, day8.RotateToLeft(lines))

	lines = []string{
		"123",
	}
	expected = []string{
		"3",
		"2",
		"1",
	}
	assert.Equal(t, expected, day8.RotateToLeft(lines))
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day8.Solve2(reader.ReadInput(8, false))
	assert.Equal(t, 1840, resultPart1)
	assert.Equal(t, 0, resultPart2)
}
