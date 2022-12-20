package day20

import (
	"adventofcode/util"
	"fmt"
)

func Solve(lines []string) (part1 int, part2 int) {
	file := parse(lines)
	return solvePart1(file), 0
}

func solvePart1(file []int) int {
	mixedFile := mixing(file)

	l := len(mixedFile)
	i := indexOf(0, mixedFile)[0]

	a := mixedFile[(i+1000)%l]
	b := mixedFile[(i+2000)%l]
	c := mixedFile[(i+3000)%l]
	return a + b + c
}

func mixing(file []int) []int {
	var fileCopy = make([]int, len(file))
	copy(fileCopy, file)

	for i, number := range file {
		t := indexOf(number, fileCopy)
		

		newIndex := (i + number) % (len(file))
		if newIndex < 0 {
			newIndex = len(file) + newIndex
		}
		l := file[newIndex]
		r := file[newIndex+1]
		insertIndex := findIndexBetween(fileCopy, l, r)

		fileCopy = remove(fileCopy, t[0])
		fileCopy = insert(fileCopy, insertIndex, number)
	}

	return fileCopy
}

func findIndexBetween(file []int, leftNumber, rightNumber int) int {
	for i := range file {
		if file[i] == leftNumber && file[(i+1)%len(file)] == rightNumber {
			return i
		}
	}
	fmt.Println("Error not found")
	return -1
}

func temp(file []int, number, oldIndex int) []int {
	if number == 0 {
		return file
	}

	var fileWithoutNumber = make([]int, len(file))
	copy(fileWithoutNumber, file)

	//oldIndex := indexOf(number, file)

	fileWithoutNumber = remove(fileWithoutNumber, oldIndex)

	newIndex := (oldIndex + number) % (len(fileWithoutNumber) + 1)

	if newIndex < 0 {
		newIndex = len(fileWithoutNumber) + newIndex
	}

	// fileWithNumber := append(fileWithoutNumber[:newIndex+1], fileWithoutNumber[newIndex:]...)
	//	fileWithNumber[newIndex] = number
	fileWithNumber := insert(fileWithoutNumber, newIndex, number)
	return fileWithNumber
}

func remove(file []int, index int) []int {
	if len(file) == index {
		return file[:index]
	}
	return append(file[:index], file[index+1:]...)
}

// 0 <= index <= len(file)
func insert(file []int, index int, value int) []int {
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

func parse(lines []string) (file []int) {
	for _, line := range lines {
		file = append(file, util.GetInt(line))
	}
	return file
}

func indexOf(element int, file []int) (indexList []int) {
	for index, number := range file {
		if element == number {
			indexList = append(indexList, index)
		}
	}
	return indexList
}

func contains(i []int, x int) bool {
	for _, v := range i {
		if v == x {
			return true
		}
	}
	return false
}
