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
	file := parseFile(lines)
	return solvePart1(file), 0
}

func solvePart1(file []Code) int {
	mixedFile := mixing(file)

	l := len(mixedFile)
	i := indexOfZero(mixedFile)

	a := mixedFile[(i+1000)%l].value
	b := mixedFile[(i+2000)%l].value
	c := mixedFile[(i+3000)%l].value
	return a + b + c
}

func mixing(originalfile []Code) (mixedFile []Code) {
	mixedFile = make([]Code, len(originalfile))
	copy(mixedFile, originalfile)

	for _, code := range originalfile {
		removeIndex := indexOf(code.startIndex, mixedFile)
		insertIndex := (removeIndex + code.value) % (len(mixedFile) - 1)
		if insertIndex < 0 {
			insertIndex = len(mixedFile) - 1 + insertIndex

		} else if insertIndex == 0 {
			insertIndex = len(mixedFile) - 1 // append at the end
		}

		mixedFile = remove(mixedFile, removeIndex)
		mixedFile = insert(mixedFile, insertIndex, code)
	}

	return mixedFile
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

func parseFile(lines []string) (file []Code) {
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

func indexOfZero(file []Code) int {
	for index, code := range file {
		if code.value == 0 {
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
