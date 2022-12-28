package day22

import (
	"adventofcode/util"
	"fmt"
	"regexp"
	"sort"
)

type Position struct {
	x int
	y int
	z int
}

type Side struct {
	number      int
	symbols     [][]rune
	original    Position
	orientation int
}

const (
	right = iota
	down
	left
	up
)

const (
	topSide = iota
	frontSide
	bottomSide
	backSide
	rigtSide
	leftSide
)

const (
	zero = iota
	ninety
	oneEighty
	twoSeventy
)

var maxSize int

func Solve(lines []string, size int) (part1 int, part2 int) {
	maxSize = size
	board, instructions := parse(lines)
	// print(board)
	return solvePart1(board, instructions), solvePart22(board, instructions)
}

func solvePart2(lines []string, instructions []string) int {
	sides, ugly := parseCube(lines[:len(lines)-2])

	direction := 0
	position := Position{0, 0, 0}

	for _, instruction := range instructions {
		if instruction == "R" {
			direction = (direction + 1) % 4
		} else if instruction == "L" {
			if direction == right {
				direction = up
			} else {
				direction--
			}
		} else {
			position = move2(position, direction, util.GetInt(instruction), sides)
			fmt.Println(position)
		}
	}

	o := getOffeset(position, ugly)
	position.x += o.x
	position.y += o.y

	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

func getOffeset(curr Position, u []Side) Position {
	for _, v := range u {
		if v.number == curr.z {
			return v.original
		}
	}
	return Position{-1, -1, -1}
}

func solvePart1(board [][]rune, instructions []string) int {
	direction := 0
	position := Position{findEdgeIndexForSide(left, Position{0, 0, 0}, board), 0, 0}

	for _, instruction := range instructions {
		if instruction == "R" {
			direction = (direction + 1) % 4
		} else if instruction == "L" {
			if direction == right {
				direction = up
			} else {
				direction--
			}
		} else {
			position = move(position, direction, util.GetInt(instruction), board)
			// fmt.Println(position)
		}
	}
	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

func move(current Position, direction, amount int, board [][]rune) Position {
	for i := 0; i < amount; i++ {
		switch direction {
		case right:
			rightEdge := findEdgeIndexForSide(right, current, board)
			if rightEdge < current.x+1 {
				leftEdge := findEdgeIndexForSide(left, current, board)
				if board[current.y][leftEdge] != '#' {
					current.x = leftEdge
				}
			} else if board[current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			bottomEdge := findEdgeIndexForSide(down, current, board)
			if bottomEdge < current.y+1 {
				topEdge := findEdgeIndexForSide(up, current, board)
				if board[topEdge][current.x] != '#' {
					current.y = topEdge
				}
			} else if board[current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			leftEdge := findEdgeIndexForSide(left, current, board)
			if leftEdge > current.x-1 {
				rightEdge := findEdgeIndexForSide(right, current, board)
				if board[current.y][rightEdge] != '#' {
					current.x = rightEdge
				}
			} else if board[current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			topEdge := findEdgeIndexForSide(up, current, board)
			if topEdge > current.y-1 {
				bottomEdge := findEdgeIndexForSide(down, current, board)
				if board[bottomEdge][current.x] != '#' {
					current.y = bottomEdge
				}
			} else if board[current.y-1][current.x] == '.' {
				current.y--
			}
		}
	}
	return current
}

func findNextPos(old Position, direction int, cube [][][]rune) Position {
	n, _ := getCubeNumber(old.z+1, direction)
	newPos := Position{old.x, old.y, n}

	if cube[newPos.z][newPos.y][newPos.x] == '#' {
		return old
	}

	return newPos
}

func temp() {

}

func move2(current Position, direction, amount int, cube [][][]rune) Position {
	for i := 0; i < amount; i++ {
		switch direction {
		case right:
			if maxSize <= current.x+1 {
				return findNextPos(current, direction, cube)
			} else if cube[current.z][current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			if maxSize <= current.y+1 {
				return findNextPos(current, direction, cube)
			} else if cube[current.z][current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			if 0 > current.x-1 {
				return findNextPos(current, direction, cube)
			} else if cube[current.z][current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			if 0 > current.y-1 {
				return findNextPos(current, direction, cube)
			} else if cube[current.z][current.y-1][current.x] == '.' {
				current.y--
			}
		}
	}
	return current
}

func findEdgeIndexForSide(side int, current Position, board [][]rune) int {
	switch side {
	case right:
		for x := len(board[current.y]) - 1; x >= 0; x-- {
			if board[current.y][x] != ' ' {
				return x
			}
		}
	case down:
		for y := len(board) - 1; y >= 0; y-- {
			if len(board[y])-1 >= current.x {
				if board[y][current.x] != ' ' {
					return y
				}
			}
		}
	case left:
		for x := 0; x < len(board[current.y])-1; x++ {
			if board[current.y][x] != ' ' {
				return x
			}
		}
	case up:
		for y := 0; y < len(board)-1; y++ {
			if len(board[y])-1 >= current.x {
				if board[y][current.x] != ' ' {
					return y
				}
			}
		}
	}
	return -1
}

func parse(lines []string) (board [][]rune, instructions []string) {
	lastLineIndex := len(lines) - 1
	board = make([][]rune, lastLineIndex-1)

	for y := 0; y < lastLineIndex-1; y++ {
		line := lines[y]
		board[y] = make([]rune, len(line))
		for x, letter := range line {
			board[y][x] = letter
		}
	}

	return board, regexp.MustCompile(`\d+|R|L`).FindAllString(lines[lastLineIndex], -1)
}

// lines - 2
func parseCube(lines []string) (sides [][][]rune, sdhdhd []Side) {
	traversed = make([][]bool, 4)
	for y := 0; y < 4; y++ {
		traversed[y] = make([]bool, 4)
	}

	sides = make([][][]rune, 6)
	for i := 0; i < 6; i++ {
		sides[i] = make([][]rune, maxSize)
		for y := 0; y < maxSize; y++ {
			sides[i][y] = make([]rune, maxSize)
		}
	}

	start := findFirstQuadrant(lines)
	sides[0] = getSide(start, lines)
	s := Side{number: 1, symbols: getSide(start, lines), orientation: 0, original: start}

	mj := tiefenSuche(start, -1, s, lines)
	mj = append([]Side{s}, mj...)

	sort.Slice(mj, func(i, j int) bool {
		return mj[i].number < mj[j].number
	})

	fmt.Println(mj)

	for i, side := range mj {
		sides[i] = translate(side)
	}

	return sides, mj
}

func translate(side Side) (result [][]rune) {
	switch side.orientation {
	case zero:
		return side.symbols
	case ninety:
		return rotate90(side.symbols)
	case oneEighty:
		result = rotate90(side.symbols)
		return rotate90(result)
	case twoSeventy:
		result = rotate90(side.symbols)
		result = rotate90(result)
		return rotate90(result)
	}
	return side.symbols
}

func rotate90(input [][]rune) (res [][]rune) {
	dim := len(input)
	res = make([][]rune, dim)
	for y := 0; y < dim; y++ {
		res[y] = make([]rune, dim)
		for x := 0; x < dim; x++ {
			res[y][x] = input[dim-1-x][y]
		}
	}
	return res
}

func translateHelp(row []rune) (collumn []rune) {
	for _, v := range row {
		collumn = append(collumn, v)
	}
	return collumn
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func tiefenSuche(current Position, movedDirection int, parent Side, lines []string) (sides []Side) {
	traversed[current.y/maxSize][current.x/maxSize] = true
	var currentSide Side
	if movedDirection == -1 {
		currentSide = parent
	} else {
		mathe := (movedDirection - parent.orientation + maxSize) % maxSize
		myNumber, kms := getCubeNumber(parent.number, mathe)
		kys := (parent.orientation + kms + maxSize) % maxSize

		currentSide = Side{number: myNumber, symbols: getSide(current, lines), orientation: kys, original: current}
		sides = append(sides, currentSide)
	}

	directions := []Position{
		{current.x + maxSize, current.y, 0},
		{current.x, current.y + maxSize, 0},
		{current.x - maxSize, current.y, 0},
		{current.x, current.y - maxSize, 0},
	}

	for orientation, direction := range directions {
		if isInBounds(direction, lines) {
			if isNotBlank(direction, lines) {
				if !traversed[direction.y/maxSize][direction.x/maxSize] {
					sides = append(sides, tiefenSuche(direction, orientation, currentSide, lines)...)
				}
			}
		}
	}

	return sides
}

func isNotBlank(pos Position, lines []string) bool {
	return lines[pos.y][pos.x] != ' '
}

func getCubeNumber(parentNumber, movedDirection int) (number, orientation int) {
	switch parentNumber {
	case 1:
		switch movedDirection {
		case right:
			return 5, ninety
		case down:
			return 4, zero
		case left:
			return 2, ninety
		case up:
			return 3, zero
		}
	case 2:
		switch movedDirection {
		case right:
			return 4, zero
		case down:
			return 6, oneEighty
		case left:
			return 3, oneEighty
		case up:
			return 1, twoSeventy
		}
	case 3:
		switch movedDirection {
		case right:
			return 5, zero
		case down:
			return 1, zero
		case left:
			return 2, oneEighty
		case up:
			return 6, twoSeventy
		}
	case 4:
		switch movedDirection {
		case right:
			return 5, oneEighty
		case down:
			return 6, ninety
		case left:
			return 2, zero
		case up:
			return 1, zero
		}
	case 5:
		switch movedDirection {
		case right:
			return 4, oneEighty
		case down:
			return 1, twoSeventy
		case left:
			return 3, zero
		case up:
			return 6, oneEighty
		}
	case 6:
		switch movedDirection {
		case right:
			return 3, twoSeventy
		case down:
			return 2, oneEighty
		case left:
			return 4, twoSeventy
		case up:
			return 5, oneEighty
		}
	}
	fmt.Println("error")
	return -1, -1
}

func isInBounds(pos Position, lines []string) bool {
	return 0 <= pos.y && pos.y < len(lines) && 0 <= pos.x && pos.x < len(lines[pos.y])
}

var traversed [][]bool

func getSide(pos Position, lines []string) (side [][]rune) {
	side = make([][]rune, maxSize)
	for y := pos.y; y < pos.y+maxSize; y++ {
		side[y%maxSize] = make([]rune, maxSize)
		for x := pos.x; x < pos.x+maxSize; x++ {
			side[y%maxSize][x%maxSize] = rune(lines[y][x])
		}
	}
	return side
}

func findFirstQuadrant(lines []string) Position {
	for x := 0; x < 4*maxSize; x += maxSize {
		if rune(lines[0][x]) != ' ' {
			return Position{x, 0, 0}
		}
	}
	fmt.Println("error not found")
	return Position{-1, -1, -1}
}

func print(board [][]rune) {
	for _, v := range board {
		fmt.Println(string(v))
	}
}
