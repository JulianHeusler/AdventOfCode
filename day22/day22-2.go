package day22

import (
	"adventofcode/util"
	"fmt"
)

type Pos struct {
	x int
	y int
}

func solvePart22(board [][]rune, instructions []string) int {
	direction := 0
	position := Pos{findEdgeIndexForSide(left, Position{0, 0, 0}, board), 0}

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
			position, direction = move3(position, direction, util.GetInt(instruction), board)
			// fmt.Println("now here = ", position)
		}
	}

	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

// kantenOffset = p
func getOffsets(p int) []Pos {
	return []Pos{
		{-p - 1, -p - 1},
		{-p, -p - 1},

		{-p + maxSize, p - maxSize},
		{-p + maxSize - 1, p - maxSize},

		{-2*p - maxSize - 1, -maxSize},
		{-2*p - maxSize - 1, -maxSize - 1},

		{-2*p + 3*maxSize - 1, -maxSize},
		{-2*p + 3*maxSize - 1, -maxSize - 1},

		{-p - 3*maxSize, p - maxSize},
		{-p - 3*maxSize - 1, p - maxSize},

		{-p + 4*maxSize - 1, -p - 1},
		{-p + 4*maxSize, -p - 1},

		{-p - 3*maxSize, p + maxSize},
		{-p - 3*maxSize - 1, p + maxSize},

		{-p + 4*maxSize - 1, -p + 2*maxSize - 1},
		{-p + 4*maxSize, -p + 2*maxSize - 1},

		{-2*p - maxSize - 1, maxSize},
		{-2*p - maxSize - 1, maxSize - 1},

		{-2*p + 3*maxSize - 1, maxSize},
		{-2*p + 3*maxSize - 1, maxSize - 1},

		{-p - 2*maxSize, p + 2*maxSize},
		{-p - 2*maxSize - 1, p + 2*maxSize},

		{-p + 3*maxSize - 1, -p + 3*maxSize - 1},
		{-p + 3*maxSize, -p + 3*maxSize - 1},

		{-p - maxSize, p + 3*maxSize},
		{-p - maxSize - 1, p + 3*maxSize},

		{-p + 2*maxSize - 1, -p + 4*maxSize - 1},
		{-p + 2*maxSize, -p + 4*maxSize - 1},

		{-p - 1, -p + 4*maxSize - 1},
		{-p, -p + 4*maxSize - 1},

		{-p + maxSize, p + 3*maxSize},
		{-p + maxSize - 1, p + 3*maxSize},

		{-2*p + maxSize - 1, 4*maxSize - 1},
		{-2*p + maxSize - 1, 4 * maxSize},
	}
}

func move3(current Pos, direction, amount int, board [][]rune) (Pos, int) {
	for i := 0; i < amount; i++ {
		switch direction {
		case right:
			if isEmtpy(Pos{current.x + 1, current.y}, board) {
				next, newDirection := findFirstEdge(current, direction, board)
				if !isStone(next, board) {
					current = next
					direction = newDirection
				}
			} else if board[current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			if isEmtpy(Pos{current.x, current.y + 1}, board) {
				next, newDirection := findFirstEdge(current, direction, board)
				if !isStone(next, board) {
					current = next
					direction = newDirection
				}
			} else if board[current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			if isEmtpy(Pos{current.x - 1, current.y}, board) {
				next, newDirection := findFirstEdge(current, direction, board)
				if !isStone(next, board) {
					current = next
					direction = newDirection
				}
			} else if board[current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			if isEmtpy(Pos{current.x, current.y - 1}, board) {
				next, newDirection := findFirstEdge(current, direction, board)
				if !isStone(next, board) {
					current = next
					direction = newDirection
				}
			} else if board[current.y-1][current.x] == '.' {
				current.y--
			}
		}
	}
	return current, direction
}

func findFirstEdge(current Pos, direction int, board [][]rune) (newPosition Pos, newDirection int) {
	var offsets []Pos
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
			newPosition = Pos{current.x - offsets[i].y, current.y + offsets[i].x}
			foundEdge, newDirection = isEdge(newPosition, Pos{current.x - offsets[i+1].y, current.y + offsets[i+1].x}, board)
		case down:
			newPosition = Pos{current.x - offsets[i].x, current.y - offsets[i].y}
			foundEdge, newDirection = isEdge(newPosition, Pos{current.x - offsets[i+1].x, current.y - offsets[i+1].y}, board)
		case left:
			newPosition = Pos{current.x + offsets[i].y, current.y - offsets[i].x}
			foundEdge, newDirection = isEdge(newPosition, Pos{current.x + offsets[i+1].y, current.y - offsets[i+1].x}, board)
		case up:
			newPosition = Pos{current.x + offsets[i].x, current.y + offsets[i].y}
			foundEdge, newDirection = isEdge(newPosition, Pos{current.x + offsets[i+1].x, current.y + offsets[i+1].y}, board)
		}

		if foundEdge {
			if i > 3 {
				fmt.Printf("from: %v to: %v\n", current, newPosition)
			}
			return newPosition, newDirection
		}
	}

	panic("")
}

func isEmtpy(pos Pos, board [][]rune) bool {
	if !isInBounds2(pos, board) {
		return true
	}
	return board[pos.y][pos.x] == ' '
}

func isStone(pos Pos, board [][]rune) bool {
	return board[pos.y][pos.x] == '#'
}

func isInBounds2(pos Pos, board [][]rune) bool {
	return 0 <= pos.y && pos.y < len(board) && 0 <= pos.x && pos.x < len(board[pos.y])
}

func isEdge(posA, posB Pos, board [][]rune) (bool, int) {
	if !isInBounds2(posA, board) {
		return false, -1
	}

	if !isEmtpy(posA, board) {
		if isEmtpy(posB, board) {
			return true, getDirectionFromSubtraction(posA, posB)
		}
	}

	return false, -1
}

func getDirectionFromSubtraction(a, b Pos) int {
	difference := Pos{b.x - a.x, b.y - a.y}
	if difference.x == 1 && difference.y == 0 {
		return left
	} else if difference.x == -1 && difference.y == 0 {
		return right
	} else if difference.x == 0 && difference.y == 1 {
		return up
	} else if difference.x == 0 && difference.y == -1 {
		return down
	}
	panic("here")
	return -1
}
