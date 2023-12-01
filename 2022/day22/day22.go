package day22

import (
	"adventofcode/util"
	"regexp"
)

type Position struct {
	x int
	y int
}

type handleEdgeMethod func(current Position, direction int) (newPosition Position, newDirection int)

const (
	right = iota
	down
	left
	up
)

var (
	board        [][]rune
	instructions []string
	maxSize      int
)

func Solve(lines []string, cubeSize int) (part1 int, part2 int) {
	maxSize = cubeSize
	board, instructions = parse(lines)
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	return simulateInstructions(func(current Position, direction int) (Position, int) {
		var next Position

		switch direction {
		case right:
			next = Position{findEdgeIndexForSide(left, current), current.y}
		case down:
			next = Position{current.x, findEdgeIndexForSide(up, current)}
		case left:
			next = Position{findEdgeIndexForSide(right, current), current.y}
		case up:
			next = Position{current.x, findEdgeIndexForSide(down, current)}
		}

		if !isStone(next) {
			return next, direction
		} else {
			return current, direction
		}
	})
}

func solvePart2() int {
	return simulateInstructions(func(current Position, direction int) (Position, int) {
		var newPosition Position
		var newDirection int

		switch direction {
		case right:
			if isEmtpy(Position{current.x + 1, current.y}) {
				newPosition, newDirection = findFirstEdge(current, direction)
			}
		case down:
			if isEmtpy(Position{current.x, current.y + 1}) {
				newPosition, newDirection = findFirstEdge(current, direction)
			}
		case left:
			if isEmtpy(Position{current.x - 1, current.y}) {
				newPosition, newDirection = findFirstEdge(current, direction)
			}
		case up:
			if isEmtpy(Position{current.x, current.y - 1}) {
				newPosition, newDirection = findFirstEdge(current, direction)
			}
		}

		if !isStone(newPosition) {
			return newPosition, newDirection
		} else {
			return current, direction
		}
	})
}

func simulateInstructions(handleEdge handleEdgeMethod) int {
	direction := right
	position := Position{findEdgeIndexForSide(left, Position{0, 0}), 0}

	for _, instruction := range instructions {
		if instruction == "R" {
			direction = (direction + 1) % 4
		} else if instruction == "L" {
			direction = (direction - 1 + 4) % 4
		} else {
			for i := 0; i < util.GetInt(instruction); i++ {
				position, direction = move(position, direction, handleEdge)
			}
		}
	}

	return calculateResult(position, direction)
}

func calculateResult(position Position, direction int) int {
	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

func move(current Position, direction int, handleEdge handleEdgeMethod) (newPosition Position, newDirection int) {
	var next Position
	switch direction {
	case right:
		next = Position{current.x + 1, current.y}
	case down:
		next = Position{current.x, current.y + 1}
	case left:
		next = Position{current.x - 1, current.y}
	case up:
		next = Position{current.x, current.y - 1}
	}

	if isEmtpy(next) {
		current, direction = handleEdge(current, direction)
	} else if isFree(next) {
		current = next
	}

	return current, direction
}

func findEdgeIndexForSide(side int, current Position) int {
	switch side {
	case right:
		for x := len(board[current.y]) - 1; x >= 0; x-- {
			if !isEmtpy(Position{x, current.y}) {
				return x
			}
		}
	case down:
		for y := len(board) - 1; y >= 0; y-- {
			if len(board[y])-1 >= current.x {
				if !isEmtpy(Position{current.x, y}) {
					return y
				}
			}
		}
	case left:
		for x := 0; x < len(board[current.y])-1; x++ {
			if !isEmtpy(Position{x, current.y}) {
				return x
			}
		}
	case up:
		for y := 0; y < len(board)-1; y++ {
			if len(board[y])-1 >= current.x {
				if !isEmtpy(Position{current.x, y}) {
					return y
				}
			}
		}
	}
	return -1
}

func findFirstEdge(current Position, direction int) (newPosition Position, newDirection int) {
	var offsets []Position
	switch direction {
	case right:
		offsets = getOffsets(current.y % maxSize)
	case down:
		offsets = getOffsets(maxSize - 1 - (current.x % maxSize))
	case left:
		offsets = getOffsets(maxSize - 1 - (current.y % maxSize))
	case up:
		offsets = getOffsets(current.x % maxSize)
	}

	for i := 0; i < len(offsets)-1; i += 2 {
		var foundEdge bool
		switch direction {
		case right:
			newPosition = Position{current.x - offsets[i].y, current.y + offsets[i].x}
			foundEdge, newDirection = isEdge(newPosition, Position{current.x - offsets[i+1].y, current.y + offsets[i+1].x})
		case down:
			newPosition = Position{current.x - offsets[i].x, current.y - offsets[i].y}
			foundEdge, newDirection = isEdge(newPosition, Position{current.x - offsets[i+1].x, current.y - offsets[i+1].y})
		case left:
			newPosition = Position{current.x + offsets[i].y, current.y - offsets[i].x}
			foundEdge, newDirection = isEdge(newPosition, Position{current.x + offsets[i+1].y, current.y - offsets[i+1].x})
		case up:
			newPosition = Position{current.x + offsets[i].x, current.y + offsets[i].y}
			foundEdge, newDirection = isEdge(newPosition, Position{current.x + offsets[i+1].x, current.y + offsets[i+1].y})
		}

		if foundEdge {
			return newPosition, newDirection
		}
	}
	return Position{}, -1
}

func isEdge(positionA, positionB Position) (bool, int) {
	if !isOutOfBounds(positionA) {
		if !isEmtpy(positionA) {
			if isEmtpy(positionB) {
				return true, getDirectionFromSubtraction(positionA, positionB)
			}
		}
	}
	return false, -1
}

func getDirectionFromSubtraction(positionA, positionB Position) int {
	difference := Position{positionB.x - positionA.x, positionB.y - positionA.y}
	if difference.y == 0 {
		if difference.x == 1 {
			return left
		} else if difference.x == -1 {
			return right
		}
	} else if difference.x == 0 {
		if difference.y == 1 {
			return up
		} else if difference.y == -1 {
			return down
		}
	}
	return -1
}

func isEmtpy(position Position) bool {
	if isOutOfBounds(position) {
		return true
	}
	return board[position.y][position.x] == ' '
}

func isStone(position Position) bool {
	return board[position.y][position.x] == '#'
}

func isFree(position Position) bool {
	return board[position.y][position.x] == '.'
}

func isOutOfBounds(position Position) bool {
	return position.y < 0 || len(board) <= position.y || position.x < 0 || len(board[position.y]) <= position.x
}

func getOffsets(edgePosition int) []Position {
	return []Position{
		{-edgePosition - 1, -edgePosition - 1},
		{-edgePosition, -edgePosition - 1},

		{-edgePosition + maxSize, edgePosition - maxSize},
		{-edgePosition + maxSize - 1, edgePosition - maxSize},

		{-2*edgePosition - maxSize - 1, -maxSize},
		{-2*edgePosition - maxSize - 1, -maxSize - 1},

		{-2*edgePosition + 3*maxSize - 1, -maxSize},
		{-2*edgePosition + 3*maxSize - 1, -maxSize - 1},

		{-edgePosition - 3*maxSize, edgePosition - maxSize},
		{-edgePosition - 3*maxSize - 1, edgePosition - maxSize},

		{-edgePosition + 4*maxSize - 1, -edgePosition - 1},
		{-edgePosition + 4*maxSize, -edgePosition - 1},

		{-edgePosition - 3*maxSize, edgePosition + maxSize},
		{-edgePosition - 3*maxSize - 1, edgePosition + maxSize},

		{-edgePosition + 4*maxSize - 1, -edgePosition + 2*maxSize - 1},
		{-edgePosition + 4*maxSize, -edgePosition + 2*maxSize - 1},

		{-2*edgePosition - maxSize - 1, maxSize},
		{-2*edgePosition - maxSize - 1, maxSize - 1},

		{-2*edgePosition + 3*maxSize - 1, maxSize},
		{-2*edgePosition + 3*maxSize - 1, maxSize - 1},

		{-edgePosition - 2*maxSize, edgePosition + 2*maxSize},
		{-edgePosition - 2*maxSize - 1, edgePosition + 2*maxSize},

		{-edgePosition + 3*maxSize - 1, -edgePosition + 3*maxSize - 1},
		{-edgePosition + 3*maxSize, -edgePosition + 3*maxSize - 1},

		{-edgePosition - maxSize, edgePosition + 3*maxSize},
		{-edgePosition - maxSize - 1, edgePosition + 3*maxSize},

		{-edgePosition + 2*maxSize - 1, -edgePosition + 4*maxSize - 1},
		{-edgePosition + 2*maxSize, -edgePosition + 4*maxSize - 1},

		{-edgePosition - 1, -edgePosition + 4*maxSize - 1},
		{-edgePosition, -edgePosition + 4*maxSize - 1},

		{-edgePosition + maxSize, edgePosition + 3*maxSize},
		{-edgePosition + maxSize - 1, edgePosition + 3*maxSize},

		{-2*edgePosition + maxSize - 1, 4*maxSize - 1},
		{-2*edgePosition + maxSize - 1, 4 * maxSize},
	}
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
