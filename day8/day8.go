package day8

import "strconv"

type position struct {
	x int
	y int
}

func Solve(lines []string) (part1, part2 int) {

	var visibleTrees []position

	for y, line := range lines {
		left := countVisibleTrees(line, y)
		for _, t := range left {
			visibleTrees = addTree(visibleTrees, t)
		}
		right := countVisibleTrees(reverse(line), y)
		for _, t := range right {
			t.x = len(line) + 1 - t.x
			visibleTrees = addTree(visibleTrees, t)
		}
	}
	for y, line := range RotateToLeft(lines) {
		up := countVisibleTrees(line, y)
		for _, t := range up {
			visibleTrees = addTree(visibleTrees, t)
		}
		down := countVisibleTrees(reverse(line), y)
		for _, t := range down {
			t.x = len(line) + 1 - t.x
			visibleTrees = addTree(visibleTrees, t)
		}
	}

	return len(visibleTrees), 0
}

func addTree(trees []position, tree position) []position {
	if !contains(trees, tree) {
		trees = append(trees, tree)
	}
	return trees
}

func contains(trees []position, candidate position) bool {
	for _, tree := range trees {
		if tree == candidate {
			return true
		}
	}
	return false
}

func countVisibleTrees(line string, y int) (treePositions []position) {
	highestTree := 0
	for x := 0; x < len(line); x++ {
		tree := getInt(line[x : x+1])
		if isVisible(tree, highestTree) {
			highestTree = tree
			treePositions = append(treePositions, position{x: x + 1, y: y + 1})
			if highestTree == 9 {
				break
			}
		}
	}
	return treePositions
}

func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isVisible(treeHeight, highestTree int) bool {
	return treeHeight > highestTree
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func reverseSlice(slice []string) (result []string) {
	for _, s := range slice {
		result = append(result, reverse(s))
	}
	return result
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

func Solve2(lines []string) (part1, part2 int) {
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if isVisibleTree(position{x: x, y: y}, lines) {
				part1++
			} else {
				part2++
			}
		}
	}
	return part1, 0
}

func isVisibleTree(tree position, lines []string) bool {
	treeHeight := getInt(lines[tree.y][tree.x : tree.x+1])
	//if treeHeight == 9 {
	//	return true
	//}

	lastX := len(lines[tree.y])
	// rightCovered
	rightCovered := isCovered(treeHeight, tree.x, lines[tree.y])
	leftCovered := isCovered(treeHeight, lastX-tree.x, reverse(lines[tree.y]))

	//rotated := RotateToLeft(lines)
	//t := len(rotated[tree.y])
	lineLength := len(lines[tree.y]) - 1
	downCovered := isCoveredVertical(treeHeight, tree.y, lines)
	upCovered := isCoveredVertical(treeHeight, lineLength-tree.y, reverseSlice(lines))

	result := !rightCovered || !leftCovered || !downCovered || !upCovered
	return result
}

func isCovered(treeHeight, x int, line string) bool {
	for i := x; i < len(line); i++ {
		if getInt(line[i:i+1]) >= treeHeight {
			return true
		}
	}
	return false
}

func isCoveredVertical(treeHeight, y int, lines []string) bool {
	for _, line := range lines {
		if y >= len(line) || y <0 {
			return false
		}
		if getInt(string(line[y])) >= treeHeight {
			return true
		}
	}
	return false
}
