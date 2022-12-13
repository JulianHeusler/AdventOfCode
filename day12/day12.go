package day12

import (
	"math"
)

type Position struct {
	x int
	y int
}

func Solve(lines []string) (part1, part2 int) {
	heightMap := parseHeightMap(lines)
	start := findSymbolPositions('S', heightMap)
	goal := findSymbolPositions('E', heightMap)
	lowStartingPoints := findSymbolPositions('a', heightMap)

	part1 = solvePart1(heightMap, start[0], goal[0])
	return part1, solvePart2(heightMap, lowStartingPoints, goal[0], part1)
}

func solvePart1(heightMap [][]rune, start, goal Position) int {
	currentPositions := []Position{start}
	initVisited(heightMap)
	setVisited(start)
	_, pathLength := search(heightMap, currentPositions, goal)
	return pathLength - 1
}

func solvePart2(heightMap [][]rune, startPositions []Position, goal Position, part1 int) int {
	pathLengths := []int{part1}

	for _, start := range startPositions {
		currentPositions := []Position{start}
		initVisited(heightMap)
		setVisited(start)
		found, pathLength := search(heightMap, currentPositions, goal)
		if found {
			pathLengths = append(pathLengths, pathLength-1)
		}
	}

	return getMinimum(pathLengths)
}

func getMinimum(pathLengths []int) int {
	minimum := math.MaxInt
	for _, pathLength := range pathLengths {
		if pathLength < minimum && pathLength > 0 {
			minimum = pathLength
		}
	}
	return minimum
}

var visited [][]bool

func search(heightMap [][]rune, currentPositions []Position, goal Position) (found bool, pathLength int) {
	var possibleNextPositions []Position
	pathLength++
	for _, current := range currentPositions {
		if current == goal {
			return true, pathLength
		}
		// already visited?
		//if !notVisited(current) {
		//	continue // maybe useless
		//}

		possibleNextPositions = append(possibleNextPositions, createNextPositionForEachDirection(heightMap, current)...)
	}

	if len(possibleNextPositions) == 0 {
		return false, -1
	}

	found2, pathLength2 := search(heightMap, possibleNextPositions, goal)
	return found2, pathLength + pathLength2
}

func createNextPositionForEachDirection(heightMap [][]rune, current Position) (nextValidPositions []Position) {
	lineLength := len(heightMap[current.y])
	columnLength := len(heightMap)

	nextPositions := []Position{
		{current.x, current.y - 1}, // up
		{current.x + 1, current.y}, // right
		{current.x, current.y + 1}, // down
		{current.x - 1, current.y}, // left
	}

	for _, position := range nextPositions {
		if isInBounds(lineLength, columnLength, position) {
			if notVisited(position) {
				if isNotToHigh(getHeight(heightMap, current), getHeight(heightMap, position)) {
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

func notVisited(pos Position) bool {
	return !visited[pos.y][pos.x]
}

func setVisited(pos Position) {
	visited[pos.y][pos.x] = true
}

func getHeight(heightMap [][]rune, pos Position) rune {
	r := heightMap[pos.y][pos.x]
	if r == 'S' {
		return 'a'
	}
	if r == 'E' {
		return 'z'
	}
	return r
}

func isNotToHigh(currentHeight, nextHeight rune) bool {
	i := int(currentHeight) + 1
	r := rune(i)
	return nextHeight <= r
}

func initVisited(heightMap [][]rune) {
	visited = make([][]bool, len(heightMap))
	for y, line := range heightMap {
		visited[y] = make([]bool, len(line))
	}
}

// could also use are real map instead of slices
func parseHeightMap(lines []string) (heightMap [][]rune) {
	heightMap = make([][]rune, len(lines))
	for y, line := range lines {
		heightMap[y] = make([]rune, len(line))
		for x, letter := range line {
			heightMap[y][x] = letter
		}
	}
	return heightMap
}

func findSymbolPositions(symbol rune, heightMap [][]rune) (positions []Position) {
	for y := range heightMap {
		for x := range heightMap[y] {
			if heightMap[y][x] == symbol {
				positions = append(positions, Position{x, y})
			}
		}
	}
	return positions
}
