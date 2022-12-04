package day1

import (
	"sort"
	"strconv"
)

func Solve(lines []string) (part1, part2 int) {
	var elfes []int
	counter := 0
	for _, line := range lines {
		if line != "" {
			number, _ := strconv.Atoi(line)
			counter += number
		} else {
			elfes = append(elfes, counter)
			counter = 0
		}
	}
	elfes = append(elfes, counter)

	sort.Ints(elfes)
	i := len(elfes) - 1
	return elfes[i], elfes[i] + elfes[i-1] + elfes[i-2]
}
