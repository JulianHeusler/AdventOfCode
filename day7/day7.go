package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type directory struct {
	name          string
	subDirectorys []*directory
	files         []*file
	parentName    string
}

type file struct {
	name string
	size int
}

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines), 0
}

func solvePart1(lines []string) int {
	var currentDir *directory
	root := &directory{name: "/"}

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line[2:], "cd") {
				if line == "$ cd .." {
					currentDir = openDir(currentDir.parentName, root)
				} else {
					currentDir = openDir(line[5:], root)
				}
			}
			// $ ls - does nothing
		} else {
			if strings.HasPrefix(line, "dir") {
				subdir := &directory{name: line[4:], parentName: currentDir.name}
				currentDir.subDirectorys = append(currentDir.subDirectorys, subdir)
			} else {
				splitted := strings.Split(line, " ")
				fileName := splitted[1]
				fileSize, _ := strconv.Atoi(splitted[0])
				currentDir.files = append(currentDir.files, &file{name: fileName, size: fileSize})
			}
		}
	}

	return part1(root)
}

func part1(root *directory) (x int) {
	size := clacSize(root)
	if size <= 100000 {
		x += size
	}

	for _, sub := range root.subDirectorys {
		x += part1(sub)
	}

	return x
}

func openDir(name string, root *directory) *directory {
	b, d := findDir(root, name)
	if !b {
		fmt.Println("Error")
	}
	return d
}

func findDir(current *directory, name string) (found bool, d *directory) {
	if current.name == name {
		return true, current
	}

	for _, sub := range current.subDirectorys {
		found2, d2 := findDir(sub, name)
		if found2 {
			return true, d2
		}
	}

	return false, &directory{}
}

func clacSize(dir *directory) (size int) {
	for _, file := range dir.files {
		size += file.size
	}

	for _, sub := range dir.subDirectorys {
		size += clacSize(sub)
	}
	return size
}
