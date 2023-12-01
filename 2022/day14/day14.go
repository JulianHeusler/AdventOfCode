package day14

import (
	"adventofcode/util"
	"errors"
	"fmt"
	"math"
	"strings"
)

type Position struct {
	x int
	y int
}

var minX = math.MaxInt
var maxX = 0
var maxY = 0
var xOffset = 3

var sandStartingPosition = Position{500, 0}
var cave [][]rune

func Solve(lines []string) (part1, part2 int) {
	pathList := parsePathLines(lines)
	setMaxBounds(pathList)
	return solvePart1(pathList), solvePart2(pathList)
}

func solvePart1(pathList [][]Position) int {
	initCave(pathList, false)
	setRuneAt('+', sandStartingPosition)

	sandBlockCount := simulateSandblocks(false)
	printCave()
	return sandBlockCount
}

func solvePart2(pathList [][]Position) int {
	initCave(pathList, true)
	setRuneAt('+', sandStartingPosition)

	sandBlockCount := simulateSandblocks(true)
	printCave()
	return sandBlockCount
}

func printCave() {
	for _, caveLine := range cave {
		var printLine string
		for _, caveBlock := range caveLine {
			printLine += string(caveBlock)
		}
		fmt.Println(printLine)
	}
}

func simulateSandblocks(hasFloor bool) (sandBlocks int) {
	for {
		fmt.Printf("----------iteration:%0d----------\n", sandBlocks)
		sandPosition, err := simSandBlock(sandStartingPosition, hasFloor)
		if err != nil {
			fmt.Println(err)
			return sandBlocks
		}

		setRuneAt('o', sandPosition)
		sandBlocks++

		if sandPosition == sandStartingPosition {
			return sandBlocks
		}

		// printCave()
	}
}

func simSandBlock(position Position, hasFloor bool) (Position, error) {
	for {
		if !hasFloor && position.y > maxY {
			return Position{}, errors.New("reached abyss")
		}
		if isAir(position) {
			position.y++
			continue
		}

		diagonalDownLeft := Position{position.x - 1, position.y}
		diagonalDownRight := Position{position.x + 1, position.y}

		if isAir(diagonalDownLeft) {
			position = diagonalDownLeft
			continue
		} else if isAir(diagonalDownRight) {
			position = diagonalDownRight
			continue
		}
		break
	}
	return Position{position.x, position.y - 1}, nil
}

func setMaxBounds(pathList [][]Position) {
	for _, path := range pathList {
		for _, pos := range path {
			if pos.x < minX {
				minX = pos.x
			}
			if pos.x > maxX {
				maxX = pos.x
			}
			if pos.y > maxY {
				maxY = pos.y
			}
		}
	}
}

// cave

func initCave(pathList [][]Position, hasFloor bool) {
	columnsCount := 2*(maxY+1+xOffset) - 1
	rowCount := maxY + 3

	cave = make([][]rune, rowCount)
	for y := 0; y < rowCount; y++ {
		cave[y] = make([]rune, columnsCount)
		for x := 0; x < len(cave[y]); x++ {
			if hasFloor && y == rowCount-1 {
				cave[y][x] = '#'
			} else {
				cave[y][x] = '.'
			}
		}
	}

	for _, path := range pathList {
		setRocksAtPath(path)
	}
}

func setRocksAtPath(path []Position) {
	prevPos := path[0]
	for i := 1; i < len(path); i++ {
		currPos := path[i]
		setRockForLine(createLineBetweenPoints(prevPos, currPos))
		prevPos = currPos
	}
}

func createLineBetweenPoints(prevPos, currPos Position) (line []Position) {
	dx := getVerticalDistance(prevPos, currPos)
	dy := getHorizontalDistance(prevPos, currPos)
	for i := 0; i < abs(dx); i++ {
		line = append(line, Position{prevPos.x + sgn(dx)*i, currPos.y})
	}
	for i := 0; i < abs(dy); i++ {
		line = append(line, Position{currPos.x, prevPos.y + sgn(dy)*i})
	}
	return append(line, currPos)
}

func abs(number int) int {
	return int(math.Abs(float64(number)))
}

func getVerticalDistance(pos1, pos2 Position) int {
	return pos2.x - pos1.x
}

func getHorizontalDistance(pos1, pos2 Position) int {
	return pos2.y - pos1.y
}

func sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func setRockForLine(path []Position) {
	for _, pos := range path {
		setRuneAt('#', pos)
	}
}

func setRuneAt(r rune, pos Position) {
	x := pos.x - sandStartingPosition.x + maxY + xOffset
	y := pos.y
	cave[y][x] = r
}

func isAir(pos Position) bool {
	x := pos.x - sandStartingPosition.x + maxY + xOffset
	y := pos.y
	return cave[y][x] == '.' || cave[y][x] == '+'
}

// parsing

func parsePathLines(lines []string) (pathList [][]Position) {
	for _, line := range lines {
		pathList = append(pathList, parsePathLine(line))
	}
	return pathList
}

func parsePathLine(line string) (path []Position) {
	positions := strings.Split(line, " -> ")
	for _, dimension := range positions {
		splitted := strings.Split(dimension, ",")
		path = append(path, Position{util.GetInt(splitted[0]), util.GetInt(splitted[1])})
	}
	return path
}
