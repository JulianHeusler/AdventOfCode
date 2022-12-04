package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"adventofcode/day1"
)

func main() {
	resultPart1, resultPart2 := day1.Solve(readInput(1))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}

func readInput(day int) (lines []string) {
	input, _ := os.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	parsedInput := strings.ReplaceAll(string(input), "\r", "")
	return strings.Split(parsedInput, "\n")
}
