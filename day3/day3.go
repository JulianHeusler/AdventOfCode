package day3

import (
	"strings"
)

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines), solvePart2(lines)
}

func solvePart1(lines []string) (sum int) {
	for _, line := range lines {
		middleIndex := len(line) / 2
		leftCompartment := line[0:middleIndex]
		rightCompartment := line[middleIndex:]
		leftSet := convertToSet(leftCompartment)
		rightSet := convertToSet(rightCompartment)

		for _, similarity := range findSimilaritis(leftSet, rightSet) {
			sum += convertToPrio(string(similarity))
		}
	}
	return sum
}

func solvePart2(lines []string) (sum int) {
	for i := 0; i < len(lines); i += 3 {
		elf1 := convertToSet(lines[i])
		elf2 := convertToSet(lines[i+1])
		elf3 := convertToSet(lines[i+2])

		for _, similarity := range findSimilaritis(findSimilaritis(elf1, elf2), elf3) {
			sum += convertToPrio(string(similarity))
		}
	}
	return sum
}

func convertToSet(s string) (set []string) {
	for _, letter := range strings.Split(s, "") {
		if !contains(set, letter) {
			set = append(set, letter)
		}
	}
	return set
}

func findSimilaritis(leftSet, rightSet []string) (similaritis []string) {
	for _, letter := range leftSet {
		if contains(rightSet, letter) {
			similaritis = append(similaritis, letter)
		}
	}
	return similaritis
}

func contains(target []string, candidate string) bool {
	for _, t := range target {
		if t == candidate {
			return true
		}
	}
	return false
}

func convertToPrio(letter string) int {
	index := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		if string(ch) == letter {
			return index
		}
		index++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		if string(ch) == letter {
			return index
		}
		index++
	}

	println("error")
	return -1
}
