package day9

import (
	"math"
	"strconv"
	"strings"
)

type Rope struct {
	head Position
	tail Position
}

type Position struct {
	X int
	Y int
}

type MoveAction struct {
	direction string
	amount    int
}

func Solve(lines []string) (part1, part2 int) {
	return calcVisitedPositions(lines, 2), calcVisitedPositions(lines, 10)
}

func calcVisitedPositions(lines []string, ropeSize int) int {
	rope := createRope(ropeSize)
	var visitedTailPositions []Position
	visitedTailPositions = append(visitedTailPositions, rope[0]) // inital position

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		direction := splitted[0]
		amount := splitted[1]
		visitedTailPositions = append(visitedTailPositions, moveRope(rope, direction, getInt(amount))...)
	}

	return countUniquePositions(visitedTailPositions)
}

func createRope(size int) (rope []Position) {
	for i := 0; i < size; i++ {
		rope = append(rope, Position{0, 0})
	}
	return rope
}

func moveRope(rope []Position, direction string, amount int) (visitedTailPositions []Position) {
	for i := 0; i < amount; i++ {
		for knotIndex := 0; knotIndex < len(rope); knotIndex++ {
			if knotIndex == 0 {
				rope[0] = movePosition(rope[0], direction, 1)
			} else {
				if IsNotTouching(rope[knotIndex-1], rope[knotIndex]) {
					//if getDistance(rope[knotIndex-1], rope[knotIndex]) == 2 {
					//	rope[knotIndex] = movePosition(rope[knotIndex], direction, 1)
					//} else {
						rope[knotIndex] = temp(rope, knotIndex) // diagonal
					//}
					if knotIndex == len(rope)-1 {
						visitedTailPositions = append(visitedTailPositions, rope[knotIndex])
					}
				}
			}
		}
	}
	return visitedTailPositions
}

func temp(rope []Position, index int) Position {
	prev := rope[index-1]
	curr := rope[index]

	dx := (prev.X - curr.X)
	dy := (prev.Y - curr.Y)

	return Position{X: curr.X + sgn(dx), Y: curr.Y + sgn(dy)}
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

func movePosition(position Position, direction string, amount int) Position {
	switch direction {
	case "R":
		position.X += amount
	case "L":
		position.X -= amount
	case "U":
		position.Y += amount
	case "D":
		position.Y -= amount
	}
	return position
}

func IsNotTouching(head, tail Position) bool {
	return math.Abs(float64(head.X-tail.X)) > 1 || math.Abs(float64(head.Y-tail.Y)) > 1
}

func getDistance(position1, position2 Position) float64 {
	return math.Sqrt(math.Pow(float64(position2.X-position1.X), 2) + math.Pow(float64(position2.Y-position1.Y), 2))
}

func countUniquePositions(positions []Position) (count int) {
	var set []Position
	for i := 0; i < len(positions); i++ {
		if !contains(set, positions[i]) {
			set = append(set, positions[i])
		}
	}
	return len(set)
}

func contains(positions []Position, candidate Position) bool {
	for _, position := range positions {
		if position.X == candidate.X && position.Y == candidate.Y {
			return true
		}
	}
	return false
}

func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
