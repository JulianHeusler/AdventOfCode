package day12

import "fmt"

type Position struct {
	x int
	y int
}

func Solve(lines []string) (part1, part2 int) {
	heightMap := parseHeightMap(lines)
	start := findSymbol("S", heightMap)
	goal := findSymbol("E", heightMap)

	fmt.Println(start)
	fmt.Println(goal)
	return 0, 0
}

// could also use are real map instead of slices
func parseHeightMap(lines []string) (heightMap [][]string) {
	heightMap = make([][]string, len(lines))

	for y, line := range lines {
		heightMap[y] = make([]string, len(line))
		for x, letter := range line {
			heightMap[y][x] = string(letter)
		}
	}
	return heightMap
}

func findSymbol(symbol string, heightMap [][]string) Position {
	for y := range heightMap {
		for x := range heightMap[y] {
			if heightMap[y][x] == symbol {
				return Position{x, y}
			}
		}
	}

	return Position{-1, -1}
}

func compare(a, b string) bool {
	return a > b
}

func isOneHigherOrEqual(a, b string) bool {
	return a >= b
}
