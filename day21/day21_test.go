package day21_test

import (
	"adventofcode/day21"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	resultPart1, resultPart2 := day21.Solve(reader.ReadExampleInput(21))
	assert.Equal(t, 3, resultPart1)
	assert.Equal(t, int64(1623178306), resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day21.Solve(reader.ReadInput(21))
	assert.Equal(t, 8764, resultPart1)
	assert.Equal(t, int64(535648840980), resultPart2)
}
