package day18

import (
	"adventofcode/util"
	"math"
	"strings"
)

type Block struct {
	x int
	y int
	z int
}

var (
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
)

func Solve(lines []string) (part1, part2 int) {
	lavaBlocks := parse(lines)
	minX, minY, minZ = findMinCoords(lavaBlocks)
	maxX, maxY, maxZ = findMaxCoords(lavaBlocks)
	return getUncoveredSides(lavaBlocks), solvePart2(lavaBlocks)
}

func solvePart2(lavaBlocks []Block) int {
	return getUncoveredSides(lavaBlocks) - getUncoveredSides(getInnerAirBlocks(lavaBlocks))
}

func getUncoveredSides(blocks []Block) (uncoveredSides int) {
	for _, block := range blocks {
		uncoveredSides += 6 - countNeighbors(blocks, block)
	}
	return uncoveredSides
}

func getInnerAirBlocks(blocks []Block) (innerAirBlocks []Block) {
	allAirBlocks := getAllAirBlocks(blocks)
	outsideAirBlocks := getOutsideAirBlocks(allAirBlocks)

	for _, airBlock := range allAirBlocks {
		if !contains(outsideAirBlocks, airBlock) {
			innerAirBlocks = append(innerAirBlocks, airBlock)
		}
	}
	return innerAirBlocks
}

func getOutsideAirBlocks(allAirBlocks []Block) (outsideAirBlocks []Block) {
	outsideAirBlocks = append(outsideAirBlocks, Block{minX - 1, minY - 1, minZ - 1})
	oldLength := 0

	for oldLength != len(outsideAirBlocks) {
		oldLength = len(outsideAirBlocks)
		for _, airBlock := range allAirBlocks {
			if !contains(outsideAirBlocks, airBlock) && countNeighbors(outsideAirBlocks, airBlock) > 0 {
				outsideAirBlocks = append(outsideAirBlocks, airBlock)
			}
		}
	}
	return outsideAirBlocks
}

func getAllAirBlocks(lavaBlocks []Block) (allAirBlocks []Block) {
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				currentBlock := Block{x, y, z}
				if !contains(lavaBlocks, currentBlock) {
					allAirBlocks = append(allAirBlocks, currentBlock)
				}
			}
		}
	}
	return allAirBlocks
}

func countNeighbors(blocks []Block, block Block) (neighbors int) {
	candidates := []Block{
		{block.x + 1, block.y, block.z},
		{block.x - 1, block.y, block.z},
		{block.x, block.y + 1, block.z},
		{block.x, block.y - 1, block.z},
		{block.x, block.y, block.z + 1},
		{block.x, block.y, block.z - 1},
	}

	for _, candiate := range candidates {
		if contains(blocks, candiate) {
			neighbors++
		}
	}
	return neighbors
}

func contains(blocks []Block, candiate Block) bool {
	for _, point := range blocks {
		if candiate == point {
			return true
		}
	}
	return false
}

func findMinCoords(blocks []Block) (int, int, int) {
	minX := math.MaxInt
	minY := math.MaxInt
	minZ := math.MaxInt
	for _, point := range blocks {
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.z < minZ {
			minZ = point.z
		}
	}
	return minX, minY, minZ
}

func findMaxCoords(blocks []Block) (maxX, maxY, maxZ int) {
	for _, point := range blocks {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.z > maxZ {
			maxZ = point.z
		}
	}
	return maxX, maxY, maxZ
}

func parse(lines []string) (points []Block) {
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x := util.GetInt(splitted[0])
		y := util.GetInt(splitted[1])
		z := util.GetInt(splitted[2])
		points = append(points, Block{x, y, z})
	}
	return points
}
