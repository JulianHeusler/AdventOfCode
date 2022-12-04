package main

import (
	"log"

	"adventofcode/day3"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day3.Solve(reader.ReadInput(3, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
