package day22

import (
	"adventofcode/util"
	"fmt"
	"regexp"
)

type Position struct {
	x int
	y int
}

const (
	right = iota
	down
	left
	up
)

func Solve(lines []string) (part1 int, part2 int) {
	board, instructions := parse(lines)
	print(board)
	return solvePart1(board, instructions), 0
}

func solvePart1(board [][]rune, instructions []string) int {
	direction := 0
	position := Position{findEdge(left, Position{0, 0}, board), 0}

	for _, instruction := range instructions {
		if instruction == "R" {
			direction = (direction + 1) % 4
		} else if instruction == "L" {
			if direction == 0 {
				direction = 3
			} else {
				direction--
			}
		} else {
			position = move(position, direction, util.GetInt(instruction), board)
			fmt.Println(position)
		}
	}
	return 1000*(position.y+1) + 4*(position.x+1) + direction
}

func move(current Position, direction, amount int, board [][]rune) Position {
	for i := 0; i < amount; i++ {
		switch direction {
		case right:
			rightEdge := findEdge(right, current, board)
			if rightEdge < current.x+1 {
				leftEdge := findEdge(left, current, board)
				if board[current.y][leftEdge] != '#' {
					current.x = leftEdge
				}
			} else if board[current.y][current.x+1] == '.' {
				current.x++
			}
		case down:
			bottomEdge := findEdge(down, current, board)
			if bottomEdge < current.y+1 {
				topEdge := findEdge(up, current, board)
				if board[topEdge][current.x] != '#' {
					current.y = topEdge
				}
			} else if board[current.y+1][current.x] == '.' {
				current.y++
			}
		case left:
			leftEdge := findEdge(left, current, board)
			if leftEdge > current.x-1 {
				rightEdge := findEdge(right, current, board)
				if board[current.y][rightEdge] != '#' {
					current.x = rightEdge
				}
			} else if board[current.y][current.x-1] == '.' {
				current.x--
			}
		case up:
			topEdge := findEdge(up, current, board)
			if topEdge > current.y-1 {
				bottomEdge := findEdge(down, current, board)
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

func findEdge(side int, position Position, board [][]rune) int {
	switch side {
	case right:
		for x := len(board[position.y]) - 1; x >= 0; x-- {
			if board[position.y][x] != ' ' {
				return x
			}
		}
	case down:
		for y := len(board) - 1; y >= 0; y-- {
			if len(board[y])-1 >= position.x {
				if board[y][position.x] != ' ' {
					return y
				}
			}
		}
	case left:
		for x := 0; x < len(board[position.y])-1; x++ {
			if board[position.y][x] != ' ' {
				return x
			}
		}
	case up:
		for y := 0; y < len(board)-1; y++ {
			if len(board[y])-1 >= position.x {
				if board[y][position.x] != ' ' {
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

	instructions = regexp.MustCompile(`\d+|R|L`).FindAllString(lines[lastLineIndex], -1)

	return board, instructions
}

func print(board [][]rune) {
	for _, v := range board {
		fmt.Println(string(v))
	}
}
