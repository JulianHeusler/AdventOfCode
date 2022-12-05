package main

import (
	"log"

	"adventofcode/day5"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day5.Solve(reader.ReadInput(5, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
