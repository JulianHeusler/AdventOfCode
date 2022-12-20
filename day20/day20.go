package day20

import (
	"adventofcode/util"
	"fmt"
)

type Code struct {
	value      int
	startIndex int
}

func Solve(lines []string) (part1 int, part2 int) {
	file := parse(lines)
	return solvePart1(file), 0
}

func solvePart1(file []Code) int {
	mixedFile := mixing(file)

	l := len(mixedFile)
	i := indexOf(0, mixedFile)

	a := mixedFile[(i+1000)%l].value
	b := mixedFile[(i+2000)%l].value
	c := mixedFile[(i+3000)%l].value
	return a + b + c
}

func mixing(file []Code) []Code {
	var fileCopy = make([]Code, len(file))
	copy(fileCopy, file)

	for _, code := range file {
		newIndex := (indexOf(code.startIndex, fileCopy) + code.value) % (len(file))
		if newIndex < 0 {
			newIndex = len(file) + newIndex - 1
			if newIndex == 0 {
				newIndex = len(fileCopy)
			}
		}
		insertIndex := newIndex

		toRemoveIndex := indexOf(code.startIndex, fileCopy)
		fileCopy = remove(fileCopy, toRemoveIndex)
		fileCopy = insert(fileCopy, insertIndex, code)
	}

	return fileCopy
}

func findIndexBetween(file []Code, leftNumber, rightNumber Code) int {
	for i := range file {
		if file[i] == leftNumber && file[(i+1)%len(file)] == rightNumber {
			return i
		}
	}
	fmt.Println("Error not found")
	return -1
}

func remove(file []Code, index int) []Code {
	if len(file) == index {
		return file[:index]
	}
	return append(file[:index], file[index+1:]...)
}

// 0 <= index <= len(file)
func insert(file []Code, index int, value Code) []Code {
	if len(file) == index { // nil or empty slice or after last element
		return append(file, value)
	}
	file = append(file[:index+1], file[index:]...) // index < len(a)
	file[index] = value
	return file
}

func swap(file []int, a, b int) []int {
	file[a], file[b] = file[b], file[a]
	return file
}

func parse(lines []string) (file []Code) {
	for i, line := range lines {
		file = append(file, Code{util.GetInt(line), i})
	}
	return file
}

func indexOf(startPos int, file []Code) int {
	for index, code := range file {
		if code.startIndex == startPos {
			return index
		}
	}
	fmt.Println("not found")
	return -1
}

func contains(i []int, x int) bool {
	for _, v := range i {
		if v == x {
			return true
		}
	}
	return false
}
