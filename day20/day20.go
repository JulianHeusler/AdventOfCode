package day20

import (
	"adventofcode/util"
	"fmt"
)

type Code struct {
	value      int
	moveCount  int
	startIndex int
}

var originalFile []Code

const decryptionKey = 811589153

func Solve(lines []string) (part1 int, part2 int64) {
	parseOriginalFile(lines)
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	return getSumOfGroveCoordinates(mixing(originalFile))
}

func solvePart2() int64 {
	for i := range originalFile {
		originalFile[i].moveCount = originalFile[i].value * (decryptionKey % (len(originalFile) - 1))
	}

	var mixedFile = make([]Code, len(originalFile))
	copy(mixedFile, originalFile)

	for i := 0; i < 10; i++ {
		mixedFile = mixing(mixedFile)
	}

	return int64(getSumOfGroveCoordinates(mixedFile)) * int64(decryptionKey)
}

func getSumOfGroveCoordinates(mixedFile []Code) int {
	l := len(mixedFile)
	i := indexOfZero(mixedFile)

	a := mixedFile[(i+1000)%l].value
	b := mixedFile[(i+2000)%l].value
	c := mixedFile[(i+3000)%l].value
	return a + b + c
}

func mixing(file []Code) (mixedFile []Code) {
	mixedFile = make([]Code, len(file))
	copy(mixedFile, file)

	for _, code := range originalFile {
		removeIndex := indexOf(code.startIndex, mixedFile)
		insertIndex := (removeIndex + code.moveCount) % (len(mixedFile) - 1)
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

func parseOriginalFile(lines []string) {
	for i, line := range lines {
		originalFile = append(originalFile, Code{util.GetInt(line), util.GetInt(line), i})
	}
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
