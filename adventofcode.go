package main

import (
	"log"

	"adventofcode/day12"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day12.Solve(reader.ReadInput(12, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
