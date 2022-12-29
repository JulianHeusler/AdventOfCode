package day22

import (
	"adventofcode/util"
	"regexp"
)

type Position struct {
	x int
	y int
}

type moveStep func(current Position, direction int) (newPosition Position, newDirection int)

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
		switch direction {
		case right:
			if findEdgeIndexForSide(right, current) < current.x+1 {
				leftEdge := findEdgeIndexForSide(left, current)
				if board[current.y][leftEdge] != '#' {
					current.x = leftEdge
				}
			} else if board[current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			bottomEdge := findEdgeIndexForSide(down, current)
			if bottomEdge < current.y+1 {
				topEdge := findEdgeIndexForSide(up, current)
				if board[topEdge][current.x] != '#' {
					current.y = topEdge
				}
			} else if board[current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			leftEdge := findEdgeIndexForSide(left, current)
			if leftEdge > current.x-1 {
				rightEdge := findEdgeIndexForSide(right, current)
				if board[current.y][rightEdge] != '#' {
					current.x = rightEdge
				}
			} else if board[current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			topEdge := findEdgeIndexForSide(up, current)
			if topEdge > current.y-1 {
				bottomEdge := findEdgeIndexForSide(down, current)
				if board[bottomEdge][current.x] != '#' {
					current.y = bottomEdge
				}
			} else if board[current.y-1][current.x] == '.' {
				current.y--
			}
		}
		return current, direction
	})
}

func solvePart2() int {
	return simulateInstructions(func(current Position, direction int) (Position, int) {
		switch direction {
		case right:
			if isEmtpy(Position{current.x + 1, current.y}) {
				next, newDirection := findFirstEdge(current, direction)
				if !isStone(next) {
					current = next
					direction = newDirection
				}
			} else if board[current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			if isEmtpy(Position{current.x, current.y + 1}) {
				next, newDirection := findFirstEdge(current, direction)
				if !isStone(next) {
					current = next
					direction = newDirection
				}
			} else if board[current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			if isEmtpy(Position{current.x - 1, current.y}) {
				next, newDirection := findFirstEdge(current, direction)
				if !isStone(next) {
					current = next
					direction = newDirection
				}
			} else if board[current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			if isEmtpy(Position{current.x, current.y - 1}) {
				next, newDirection := findFirstEdge(current, direction)
				if !isStone(next) {
					current = next
					direction = newDirection
				}
			} else if board[current.y-1][current.x] == '.' {
				current.y--
			}
		}
		return current, direction
	})
}

func simulateInstructions(move moveStep) int {
	direction := right
	position := Position{findEdgeIndexForSide(left, Position{0, 0}), 0}

	for _, instruction := range instructions {
		if instruction == "R" {
			direction = (direction + 1) % 4
		} else if instruction == "L" {
			direction = (direction - 1 + 4) % 4
		} else {
			for i := 0; i < util.GetInt(instruction); i++ {
				position, direction = move(position, direction)
			}
		}
	}

	return calculateResult(position, direction)
}

func calculateResult(position Position, direction int) int {
	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

func findEdgeIndexForSide(side int, current Position) int {
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
	if difference.x == 1 && difference.y == 0 {
		return left
	} else if difference.x == -1 && difference.y == 0 {
		return right
	} else if difference.x == 0 && difference.y == 1 {
		return up
	} else if difference.x == 0 && difference.y == -1 {
		return down
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
