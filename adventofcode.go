package main

import (
	"log"

	"adventofcode/day10"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day10.Solve(reader.ReadInput(10, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
