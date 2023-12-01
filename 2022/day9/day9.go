package day9

import (
	"math"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func Solve(lines []string) (part1, part2 int) {
	return calcVisitedPositions(lines, 2), calcVisitedPositions(lines, 10)
}

func calcVisitedPositions(lines []string, ropeSize int) int {
	rope := createRope(ropeSize)
	var visitedTailPositions []Position
	visitedTailPositions = append(visitedTailPositions, rope[0]) // add inital position

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
	for step := 0; step < amount; step++ {
		movePosition(&rope[0], direction)
		visitedTailPositions = append(visitedTailPositions, moveKnots(rope, direction)...)
	}
	return visitedTailPositions
}

func moveKnots(rope []Position, direction string) (visitedTailPositions []Position) {
	var current *Position
	var previous *Position
	for knotIndex := 1; knotIndex < len(rope); knotIndex++ {
		current = &rope[knotIndex]
		previous = &rope[knotIndex-1]

		if IsNotTouching(*current, *previous) {
			clacNewPosition(current, previous)
			if knotIndex == len(rope)-1 { // is tail
				visitedTailPositions = append(visitedTailPositions, *current)
			}
		}
	}
	return visitedTailPositions
}

func clacNewPosition(curr, prev *Position) {
	dx := (prev.X - curr.X)
	dy := (prev.Y - curr.Y)

	curr.X += sgn(dx)
	curr.Y += sgn(dy)
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

func movePosition(position *Position, direction string) {
	switch direction {
	case "R":
		position.X++
	case "L":
		position.X--
	case "U":
		position.Y++
	case "D":
		position.Y--
	}
}

func IsNotTouching(head, tail Position) bool {
	return math.Abs(float64(head.X-tail.X)) > 1 || math.Abs(float64(head.Y-tail.Y)) > 1
}

func countUniquePositions(positions []Position) int {
	var set []Position
	for _, position := range positions {
		if !contains(set, position) {
			set = append(set, position)
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
