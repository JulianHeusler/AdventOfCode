package day18

import (
	"adventofcode/util"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

func Solve(lines []string) (part1, part2 int) {
	points := parse(lines)
	return solvePart1(points), 0
}

func solvePart1(points []Point) (surfaceArea int) {
	for _, point := range points {
		surfaceArea += countUncoveredSides(points, point)
	}
	return surfaceArea
}

func countUncoveredSides(points []Point, point Point) (uncoveredSides int) {
	candidates := []Point{
		{point.x + 1, point.y, point.z},
		{point.x - 1, point.y, point.z},
		{point.x, point.y + 1, point.z},
		{point.x, point.y - 1, point.z},
		{point.x, point.y, point.z + 1},
		{point.x, point.y, point.z - 1},
	}

	for _, candiate := range candidates {
		if !contains(points, candiate) {
			uncoveredSides++
		}
	}
	return uncoveredSides
}

func contains(points []Point, candiate Point) bool {
	for _, point := range points {
		if candiate == point {
			return true
		}
	}
	return false
}

func parse(lines []string) (points []Point) {
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x := util.GetInt(splitted[0])
		y := util.GetInt(splitted[1])
		z := util.GetInt(splitted[2])
		points = append(points, Point{x, y, z})
	}
	return points
}
