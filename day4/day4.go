package day4

import (
	"strconv"
	"strings"
)

type interval struct {
	from int
	to   int
}

func Solve(lines []string) (part1, part2 int) {
	for _, line := range lines {
		interval1, interval2 := parseIntervals(line)
		if fullyContained(interval1, interval2) {
			part1++
		}
		if overlapping(interval1, interval2) {
			part2++
		}
	}
	return part1, part2
}

func fullyContained(a, b interval) bool {
	if isInside(a.from, b) && isInside(a.to, b) {
		return true // a is inside b
	}
	if isInside(b.from, a) && isInside(b.to, a) {
		return true // b is inside a
	}
	return false
}

func overlapping(a, b interval) bool {
	return isInside(a.from, b) || isInside(a.to, b) || isInside(b.from, a) || isInside(b.to, a)
}

func isInside(x int, i interval) bool {
	return i.from <= x && x <= i.to
}

func parseIntervals(line string) (left, right interval) {
	splitted := strings.Split(line, ",")
	return parseInterval(splitted[0]), parseInterval(splitted[1])
}

func parseInterval(side string) interval {
	splitted := strings.Split(side, "-")
	leftLimit, _ := strconv.Atoi(splitted[0])
	rightLimit, _ := strconv.Atoi(splitted[1])
	return interval{
		from: leftLimit,
		to:   rightLimit,
	}
}
