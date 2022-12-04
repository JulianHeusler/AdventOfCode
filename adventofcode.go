package main

import (
	"fmt"
	"os"
	"strings"

	"adventofcode/day1"
)

func main() {
	day1.Solve(readInput(1))
}

func readInput(day int) (lines []string) {
	input, _ := os.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	parsedInput := strings.ReplaceAll(string(input), "\r", "")
	return strings.Split(parsedInput, "\n")
}
