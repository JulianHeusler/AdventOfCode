package day1

import (
	"sort"
	"strconv"
)

func Solve(lines []string) {

	var elfes []int
	counter := 0
	for _, line := range lines {
		if line != "" {
			i, _ := strconv.Atoi(line)
			counter += i
		} else {
			elfes = append(elfes, counter)
			counter = 0
		}
	}

	sort.Ints(elfes)
	i := len(elfes) - 1

	println("Day 1 - Part 1: ", elfes[i])
	println("Day 1 - Part 2: ", elfes[i]+elfes[i-1]+elfes[i-2])
}
