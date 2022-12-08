package day7

import (
	"strconv"
	"strings"
)

type directory struct {
	name          string
	subDirectorys []directory
	files         []file
}

type file struct {
	name string
	size int
}

var filesystem []directory

func Solve(lines []string) (part1, part2 int) {

	return solvePart1(lines), 0
}

func solvePart1(lines []string) int {

	var root *directory
	root = &directory{
		name: "/",
	}

	var currentDir *directory
	var lastAddedDir *directory
	added := true

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if !added && currentDir != nil {
				currentDir.subDirectorys = append(currentDir.subDirectorys, *lastAddedDir)
				added = true
			}

			if strings.HasPrefix(line[2:], "cd") {
				currentDir = openDir(line[5:], root)
			}

		} else {
			if strings.HasPrefix(line, "dir") {
				if !added {
					currentDir.subDirectorys = append(currentDir.subDirectorys, *lastAddedDir)
					added = true
				}

				lastAddedDir = &directory{name: line[4:]}
				added = false
			} else {
				splitted := strings.Split(line, " ")
				fileName := splitted[1]
				fileSize, _ := strconv.Atoi(splitted[0])
				lastAddedDir.files = append(lastAddedDir.files, file{name: fileName, size: fileSize})
			}
		}
	}

	if !added {
		currentDir.subDirectorys = append(currentDir.subDirectorys, *lastAddedDir)
		added = true
	}

	return clacSize(root)
}

func parseInstruction(line string) {
	if strings.HasPrefix(line, "cd") {

	}
}

func openDir(name string, root *directory) *directory {
	if root.name == name {
		return root
	}

	_, d := findDir(root, name)
	return d
}

func findDir(current *directory, name string) (found bool, d *directory) {
	if current.name == name {
		return true, current
	}

	for _, sub := range current.subDirectorys {
		found2, d2 := findDir(&sub, name)
		if found2 {
			return true, d2
		}
	}

	return false, &directory{}
}

func clacSize(dir *directory) int {
	sum := 0
	for _, file := range dir.files {
		sum += file.size
	}

	for _, sub := range dir.subDirectorys {
		sum += clacSize(&sub)
	}
	return sum
}
