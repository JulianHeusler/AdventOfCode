package day1_test

import (
	"adventofcode/day1"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveWithTestData(t *testing.T) {
	testInput := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	resultPart1, resultPart2 := day1.Solve(testInput)
	assert.Equal(t, 24000, resultPart1)
	assert.Equal(t, 45000, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day1.Solve(reader.ReadInput(1))
	assert.Equal(t, 69836, resultPart1)
	assert.Equal(t, 207968, resultPart2)
}
