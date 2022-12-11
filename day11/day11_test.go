package day11_test

import (
	"adventofcode/day11"
	"adventofcode/reader"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	lines := []string{
		"Monkey 0:",
		"  Starting items: 79, 98",
		"  Operation: new = old * 19",
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
		"",
		"Monkey 1:",
		"  Starting items: 54, 65, 75, 74",
		"  Operation: new = old + 6",
		"  Test: divisible by 19",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 0",
		"",
		"Monkey 2:",
		"  Starting items: 79, 60, 97",
		"  Operation: new = old * old",
		"  Test: divisible by 13",
		"    If true: throw to monkey 1",
		"    If false: throw to monkey 3",
		"",
		"Monkey 3:",
		"  Starting items: 74",
		"  Operation: new = old + 3",
		"  Test: divisible by 17",
		"    If true: throw to monkey 0",
		"    If false: throw to monkey 1",
	}

	resultPart1, resultPart2 := day11.Solve(lines)
	assert.Equal(t, 10605, resultPart1)
	assert.Equal(t, 0, resultPart2)
}

func TestSolveRealInput(t *testing.T) {
	resultPart1, resultPart2 := day11.Solve(reader.ReadInput(11, false))
	assert.Equal(t, 1447046, resultPart1)
	assert.Equal(t, 578710, resultPart2)
}
