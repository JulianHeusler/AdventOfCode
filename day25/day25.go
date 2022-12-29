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
	for _, line := range lines {
		result = addSNAFU(result, line)
	}
	return result
}

func addSNAFU(a, b string) (result string) {
	a, b = padd(a, b)
	var overflow string
	var r string

	for i := len(a) - 1; i >= 0; i-- {
		r, overflow = vollAddierer(string(a[i]), string(b[i]), overflow)
		result = r + result
	}

	if result[0] == '0' {
		return result[1:]
	}
	return result
}

func padd(a, b string) (string, string) {
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
	return "0" + a, "0" + b
}

func vollAddierer(a, b, o string) (res, overflow string) {
	sum := toInt(a) + toInt(b) + toInt(o)
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
