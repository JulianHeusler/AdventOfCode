package main

import (
	"log"

	"adventofcode/day20"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day20.Solve(reader.ReadInput(20))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
