package main

import (
	"log"

	"adventofcode/day2"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day2.Solve(reader.ReadInput(2, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
