package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type directory struct {
	name          string
	subDirectorys []*directory
	files         []*file
	parent        *directory
	size          int
}

type file struct {
	name string
	size int
}

var uniqueDirs []*directory

func Solve(lines []string) (part1, part2 int) {
	root := parseDirectorys(lines)
	return solvePart1(), solvePart2(root.size)
}

func parseDirectorys(lines []string) directory {
	root := &directory{name: "/"}
	currentDir := root

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line[2:], "cd") {
				if line == "$ cd .." {
					currentDir = currentDir.parent
				} else {
					currentDir = openSubDirectory(currentDir, line[5:])
				}
			}
			// $ ls - does nothing
		} else {
			if strings.HasPrefix(line, "dir") {
				subdir := &directory{name: line[4:], parent: currentDir}
				currentDir.subDirectorys = append(currentDir.subDirectorys, subdir)
				uniqueDirs = append(uniqueDirs, subdir)
			} else {
				currentDir.files = append(currentDir.files, parseFile(line))
			}
		}
	}

	calcDirectorySizes(root)
	return *root
}

func parseFile(line string) *file {
	splitted := strings.Split(line, " ")
	fileName := splitted[1]
	fileSize, _ := strconv.Atoi(splitted[0])
	return &file{name: fileName, size: fileSize}
}

func openSubDirectory(current *directory, name string) *directory {
	for _, sub := range current.subDirectorys {
		if sub.name == name {
			return sub
		}
	}
	fmt.Println("error not found")
	return current
}

func solvePart1() (x int) {
	for _, dir := range uniqueDirs {
		if dir.size <= 100000 {
			x += dir.size
		}
	}
	return x
}

func solvePart2(rootSize int) (minimumFittingSize int) {
	freeSpace := 70000000 - rootSize
	neededSpace := 30000000 - freeSpace
	minimumFittingSize = math.MaxInt

	for _, dir := range uniqueDirs {
		if dir.size >= neededSpace {
			if dir.size < minimumFittingSize {
				minimumFittingSize = dir.size
			}
		}
	}
	return minimumFittingSize
}

func calcDirectorySizes(dir *directory) (size int) {
	for _, sub := range dir.subDirectorys {
		size += calcDirectorySizes(sub)
	}
	for _, file := range dir.files {
		size += file.size
	}
	dir.size = size
	return size
}
