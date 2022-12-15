package main

import (
	"log"

	"adventofcode/day15"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day15.Solve(reader.ReadInput(15))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
