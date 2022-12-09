package day8

import "strconv"

type position struct {
	x int
	y int
}

func Solve(lines []string) (part1, part2 int) {
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			tree := position{x: x, y: y}
			treeIsVisible, senicScore := isTreeVisibleAndCalcScenicScore(tree, lines)
			if treeIsVisible {
				part1++
			}
			if part2 < senicScore {
				part2 = senicScore
			}
		}
	}
	return part1, part2
}

func isTreeVisibleAndCalcScenicScore(tree position, lines []string) (bool, int) {
	treeHeight := getInt(lines[tree.y][tree.x : tree.x+1])
	lineLength := len(lines[tree.y])
	column := getColumn(lines, tree.x)
	columnLength := len(column)

	rightCovered, visibleTreesRight := visibleTreesInLine(treeHeight, tree.x, lines[tree.y])
	leftCovered, visibleTreesLeft := visibleTreesInLine(treeHeight, lineLength-1-tree.x, reverse(lines[tree.y]))
	downCovered, visibleTreesDown := visibleTreesInLine(treeHeight, tree.y, column)
	upCovered, visibleTreesUp := visibleTreesInLine(treeHeight, columnLength-1-tree.y, reverse(column))

	treeIsVisible := !rightCovered || !leftCovered || !downCovered || !upCovered
	scenicScore := visibleTreesRight * visibleTreesLeft * visibleTreesDown * visibleTreesUp
	return treeIsVisible, scenicScore
}

func visibleTreesInLine(treeHeight, linePosition int, line string) (isCovered bool, count int) {
	for i := linePosition + 1; i < len(line); i++ {
		currentTree := getInt(line[i : i+1])
		count++

		if currentTree >= treeHeight {
			return true, count
		}
	}
	return false, count
}

func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func getColumn(lines []string, y int) (column string) {
	for _, line := range lines {
		column += line[y : y+1]
	}
	return column
}

func RotateToLeft(lines []string) (rotated []string) {
	lineLength := len(lines[0])
	for x := lineLength; x > 0; x-- {
		var temp string
		for y := 0; y < len(lines); y++ {
			temp += lines[y][x-1 : x]
		}
		rotated = append(rotated, temp)
	}
	return rotated
}
