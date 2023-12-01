package day12

import (
	"log"
	"math"
)

type Position struct {
	x int
	y int
}

var heightMap [][]rune
var visited [][]bool

func Solve(lines []string) (part1, part2 int) {
	parseHeightMap(lines)
	start := findFirstSymbolPosition('S')
	goal := findFirstSymbolPosition('E')
	lowStartingPoints := findSymbolPositions('a')

	pathLengthStart := findPathLengths([]Position{start}, goal)[0]
	return pathLengthStart, getMinimum(append(findPathLengths(lowStartingPoints, goal), pathLengthStart))
}

func findPathLengths(startPositions []Position, goal Position) (pathLengths []int) {
	for _, start := range startPositions {
		currentPositions := []Position{start}
		initVisited(start)
		found, pathLength := breadthFirstSearch(currentPositions, goal)
		if found {
			pathLengths = append(pathLengths, pathLength-1) // minus first check
		}
	}
	return pathLengths
}

func getMinimum(pathLengths []int) int {
	minimum := math.MaxInt
	for _, pathLength := range pathLengths {
		if pathLength < minimum && pathLength >= 0 {
			minimum = pathLength
		}
	}
	return minimum
}

func breadthFirstSearch(currentPositions []Position, goal Position) (bool, int) {
	var possibleNextPositions []Position
	for _, current := range currentPositions {
		if current == goal {
			return true, 1
		}

		possibleNextPositions = append(possibleNextPositions, getNextValidPositions(current)...)
	}

	if len(possibleNextPositions) == 0 {
		log.Println("did not find goal")
		return false, -1
	}

	found, pathLength := breadthFirstSearch(possibleNextPositions, goal)
	return found, 1 + pathLength
}

func getNextValidPositions(current Position) (nextValidPositions []Position) {
	lineLength := len(heightMap[current.y])
	columnLength := len(heightMap)

	candidates := []Position{
		{current.x, current.y - 1}, // up
		{current.x + 1, current.y}, // right
		{current.x, current.y + 1}, // down
		{current.x - 1, current.y}, // left
	}

	for _, position := range candidates {
		if isInBounds(lineLength, columnLength, position) {
			if notVisited(position) {
				if isNotToHigh(getHeight(current), getHeight(position)) {
					nextValidPositions = append(nextValidPositions, position)
					setVisited(position)
				}
			}
		}
	}
	return nextValidPositions
}

func isInBounds(lineLength, columnLength int, position Position) bool {
	return 0 <= position.x && position.x < lineLength &&
		0 <= position.y && position.y < columnLength
}

func initVisited(startPosition Position) {
	visited = make([][]bool, len(heightMap))
	for y, line := range heightMap {
		visited[y] = make([]bool, len(line))
	}
	setVisited(startPosition)
}

func setVisited(pos Position) {
	visited[pos.y][pos.x] = true
}

func notVisited(pos Position) bool {
	return !visited[pos.y][pos.x]
}

func getHeight(pos Position) rune {
	height := heightMap[pos.y][pos.x]
	if height == 'S' {
		return 'a'
	}
	if height == 'E' {
		return 'z'
	}
	return height
}

func isNotToHigh(currentHeight, nextHeight rune) bool {
	return nextHeight <= rune(int(currentHeight)+1)
}

func parseHeightMap(lines []string) {
	heightMap = make([][]rune, len(lines))
	for y, line := range lines {
		heightMap[y] = make([]rune, len(line))
		for x, letter := range line {
			heightMap[y][x] = letter
		}
	}
}

func findFirstSymbolPosition(symbol rune) Position {
	return findSymbolPositions(symbol)[0]
}

func findSymbolPositions(symbol rune) (positions []Position) {
	for y := range heightMap {
		for x := range heightMap[y] {
			if heightMap[y][x] == symbol {
				positions = append(positions, Position{x, y})
			}
		}
	}
	return positions
}
