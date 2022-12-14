package main

import (
	"log"

	"adventofcode/day14"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day14.Solve(reader.ReadInput(14))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
