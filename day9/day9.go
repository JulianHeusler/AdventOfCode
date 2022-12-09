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

	return solvePart1(lines), 0
}

func solvePart1(lines []string) int {
	rope := Rope{Position{0, 0}, Position{0, 0}}
	var visitedTailPositions []Position
	visitedTailPositions = append(visitedTailPositions, rope.tail) // inital position

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		direction := splitted[0]
		amount := splitted[1]
		visitedTailPositions = append(visitedTailPositions, moveRope(&rope, direction, getInt(amount))...)
	}

	return countUniquePositions(visitedTailPositions)
}

func moveRope(rope *Rope, direction string, amount int) (visitedTailPositions []Position) {
	for i := 0; i < amount; i++ {
		oldPosition := rope.head
		movePosition(&rope.head, direction, 1)
		if IsNotTouching(rope.head, rope.tail) {
			//movePosition(&rope.tail, direction, 1)
			//if getDistance(rope.head, rope.tail) > math.Sqrt(2) {
			// TODO digonal
			rope.tail = Position{oldPosition.X, oldPosition.Y}
			visitedTailPositions = append(visitedTailPositions, rope.tail)
		}
	}
	return visitedTailPositions
}

func movePosition(position *Position, direction string, amount int) {
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
}

func IsNotTouching(head, tail Position) bool {
	return math.Abs(float64(head.X-tail.X)) > 1 || math.Abs(float64(head.Y-tail.Y)) > 1
}

func getDistance(position1, position2 Position) float64 {
	return math.Sqrt(math.Pow(float64(position2.X-position1.X), 2) + math.Pow(float64(position2.Y-position1.Y), 2))
}

func countUniquePositions(positions []Position) (count int) {
	for index, position := range positions {
		if !contains(positions[index+1:], position) {
			count++
		}
	}
	return count
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
