package day25

import (
	"adventofcode/util"
	"strconv"
)

// SNAFU =,-,0,1,2

func Solve(lines []string) (part1 string, part2 int) {
	return solvePart1(lines), 0
}

func solvePart1(lines []string) (result string) {
	for _, current := range lines {
		result = addSNAFUNumbers(result, current)
	}
	return result
}

func addSNAFUNumbers(numberA, numberB string) (result string) {
	numberA, numberB = paddOnSameLength(numberA, numberB)
	overflow := "0"
	var currentSymbol string

	for i := len(numberA) - 1; i >= 0; i-- {
		currentSymbol, overflow = fullAdder(string(numberA[i]), string(numberB[i]), overflow)
		result = currentSymbol + result
	}

	if overflow != "0" {
		result = overflow + result
	}

	return result
}

func paddOnSameLength(a, b string) (string, string) {
	if len(a) > len(b) {
		count := len(a) - len(b)
		for i := 0; i < count; i++ {
			b = "0" + b
		}
	} else {
		count := len(b) - len(a)
		for i := 0; i < count; i++ {
			a = "0" + a
		}
	}
	return a, b
}

func fullAdder(a, b, oldOverflow string) (res, overflow string) {
	sum := toInt(a) + toInt(b) + toInt(oldOverflow)
	if sum < -2 {
		return toSymbol(sum + 5), "-"
	} else if sum > 2 {
		return toSymbol(sum - 5), "1"
	} else {
		return toSymbol(sum), "0"
	}
}

func toInt(letter string) int {
	if letter == "=" {
		return -2
	}
	if letter == "-" {
		return -1
	}
	return util.GetInt(letter)
}

func toSymbol(i int) string {
	if i == -2 {
		return "="
	}
	if i == -1 {
		return "-"
	}
	return strconv.Itoa(i)
}
