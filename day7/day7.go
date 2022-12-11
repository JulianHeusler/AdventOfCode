package day7

import (
	"fmt"
	"math"
	"reflect"
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

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines)
}

func solvePart1(lines []string) (int, int) {
	var currentDir *directory
	root := &directory{name: "/"}
	currentDir = root

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line[2:], "cd") {
				if line == "$ cd .." {
					currentDir = currentDir.parent
				} else {
					if currentDir != nil {
						currentDir = openDirHere(currentDir, line[5:])
					}
				}
			}
			// $ ls - does nothing
		} else {
			if strings.HasPrefix(line, "dir") {
				subdir := &directory{name: line[4:], parent: currentDir}
				currentDir.subDirectorys = append(currentDir.subDirectorys, subdir)
				dirs = append(dirs, subdir)
			} else {
				splitted := strings.Split(line, " ")
				fileName := splitted[1]
				fileSize, _ := strconv.Atoi(splitted[0])
				currentDir.files = append(currentDir.files, &file{name: fileName, size: fileSize})
				//currentDir.size += fileSize
			}
		}
	}

	calcSize(root)
	fmt.Println(root)
	return part1(root), part2(root)
}

func openDirHere(current *directory, name string) *directory {
	for _, sub := range current.subDirectorys {
		if sub.name == name {
			return sub
		}
	}
	fmt.Println("error not found")
	return current
}

func contains(dirs []*directory, candidate *directory) bool {
	for _, dir := range dirs {
		if reflect.DeepEqual(dir, candidate) {
			fmt.Println(candidate)
			return true
		}
	}
	return false
}

func part1(dir *directory) (x int) {
	for _, sub := range dir.subDirectorys {
		x += part1(sub)
	}

	if dir.size <= 100000 {
		x += dir.size
	}
	return x
}

func part2(root *directory) (minimumFittingSize int) {
	const filesystemSize = 70000000
	const updateNeededSpace = 30000000
	freeSpace := filesystemSize - root.size
	neededSpace := updateNeededSpace - freeSpace
	minimumFittingSize = math.MaxInt

	for _, dir := range dirs {
		if dir.size >= neededSpace {
			if dir.size < minimumFittingSize {
				minimumFittingSize = dir.size
			}
		}
	}
	return minimumFittingSize
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

var dirs []*directory

func calcSize(dir *directory) (size int) {
	for _, sub := range dir.subDirectorys {
		size += calcSize(sub)
	}

	for _, file := range dir.files {
		size += file.size
	}

	dir.size = size

	if !contains(dirs, dir) {
		dirs = append(dirs, dir)
	}
	return size
}
