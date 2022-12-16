package main

import (
	"log"

	"adventofcode/day13"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day13.Solve(reader.ReadInput(13))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
