package main

import (
	"log"

	"adventofcode/day8"
	"adventofcode/reader"
)

func main() {
	resultPart1, resultPart2 := day8.Solve2(reader.ReadInput(8, true))
	log.Printf("Part 1: %v\n", resultPart1)
	log.Printf("Part 2: %v\n", resultPart2)
}
