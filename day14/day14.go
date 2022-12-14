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
var xOffset = -1
var yOffset = -2

var sandPos = Position{500, 0}

var cave [][]rune

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) int {
	pathList := parse(lines)
	setMaxBounds(pathList)
	initCave(pathList, false)

	setRuneAt('+', sandPos)

	//printCave()

	return simulateSandrush(sandPos, false)
}

func solvePart2(lines []string) int {
	pathList := parse(lines)
	setMaxBounds(pathList)
	initCave(pathList, true)

	sandPos := Position{500, 0}
	//setRuneAt('+', sandPos)

	// printCave()

	return simulateSandrush(sandPos, true)
}

func printCave() {
	for _, caveLine := range cave {
		var printLine string
		for _, caveRune := range caveLine {
			printLine += string(caveRune)
		}
		fmt.Println(printLine)
	}
}

func simulateSandrush(startPos Position, hasFloor bool) int {
	sandBlocks := 0
	for {
		fmt.Printf("----------iteration:%0d----------\n", sandBlocks)
		err, pos := simSandBlock(startPos, hasFloor)
		if err != nil {
			fmt.Println(err)
			return sandBlocks
		}
		if pos.x == startPos.x && pos.y == startPos.y {
			return sandBlocks + 1
		}
		// printCave()
		sandBlocks++
	}
}

func simSandBlock(pos Position, hasFloor bool) (error, Position) {
	for {
		if !hasFloor && pos.y > maxY {
			return errors.New("endless"), Position{}
		}

		if isAir(pos) {
			pos.y++
			continue
		}
		diagonalDownLeft := Position{pos.x - 1, pos.y}
		diagonalDownRight := Position{pos.x + 1, pos.y}

		if isAir(diagonalDownLeft) {
			pos = diagonalDownLeft
			continue
		} else if isAir(diagonalDownRight) {
			pos = diagonalDownRight
			continue
		}
		break
	}
	setRuneAt('o', Position{pos.x, pos.y - 1})
	return nil, Position{pos.x, pos.y - 1}
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
	columnsCount := 2*(maxY+4) - 1
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

	for _, rockPath := range pathList {
		setRockPaths(rockPath)
	}
}

func setRockPaths(path []Position) {
	prev := path[0]
	for i := 1; i < len(path); i++ {
		curr := path[i]
		line := getLine(prev, curr)
		setRockLine(line)
		prev = curr
	}
}

func getLine(prev, curr Position) (line []Position) {
	dx := getVerticalDistance(prev, curr)
	dy := getHorizontalDistance(prev, curr)
	for i := 0; i < int(math.Abs(float64(dx))); i++ {
		line = append(line, Position{prev.x + sgn(dx)*i, curr.y})
	}
	for i := 0; i < int(math.Abs(float64(dy))); i++ {
		line = append(line, Position{curr.x, prev.y + sgn(dy)*i})
	}
	line = append(line, curr)
	return line
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

func setRockLine(path []Position) {
	for _, pos := range path {
		setRuneAt('#', pos)
	}
}

func setRuneAt(r rune, pos Position) {
	x := pos.x - sandPos.x + maxY + 3
	y := pos.y
	cave[y][x] = r
}

func isAir(pos Position) bool {
	x := pos.x - sandPos.x + maxY + 3
	y := pos.y
	return cave[y][x] == '.' || cave[y][x] == '+'
}

// parsing

func parse(lines []string) (pathList [][]Position) {
	for _, line := range lines {
		rockPath := parseLine(line)
		pathList = append(pathList, rockPath)
	}
	return pathList
}

func parseLine(line string) (path []Position) {
	coords := strings.Split(line, " -> ")
	for _, c := range coords {
		s := strings.Split(c, ",")
		path = append(path, Position{util.GetInt(s[0]), util.GetInt(s[1])})
	}
	return path
}
