package day17

import (
	"fmt"
	"math"
	"strings"
)

const maxLength = 7

var jetIndex int

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines[0], 2022), solvePart1(lines[0], 1000000000000)
}

func solvePart1(jetInput string, rounds int) int {
	var chamber []string
	jetIndex = 0
	rocks := getDefaultRocks()
	score := 0

	test := 5
	f := 5 * len(jetInput)
	temp := 0

	for i := 0; i < rounds; i++ {
		if i%10000 == 0 {
			percent := int(float64(i) / float64(rounds) * 100)
			fmt.Printf("%d%% | iteration: %d, score:=%d\n", percent, i, score)
		}

		chamber = simulateRock(chamber, jetInput, rocks[i%5])

		newChamber, removedLines := removeBlockedLines(chamber)
		chamber = newChamber
		score += removedLines

		//newChamber2, tetrisScore := tetris(chamber)
		//chamber = newChamber2
		//score += tetrisScore
		if i%f == 0 {
			fmt.Println(len(chamber) + score - temp)
			temp = len(chamber) + score
			test--
		}
	}
	return len(chamber) + score
}

func tetris(chamber []string) ([]string, int) {
	maxY := 0
	for y, line := range chamber {
		if line == "#######" {
			maxY = y
		}
	}
	return chamber[maxY:], maxY
}

func removeBlockedLines(chamber []string) ([]string, int) {
	for y := range chamber {
		pos := canNotSeeSky(chamber, y)
		if pos != -1 {
			return chamber[pos:], pos
		}
	}
	return chamber, 0
}

func canNotSeeSky(chamber []string, startY int) int {
	var levels []int
	for x := 0; x < maxLength; x++ {
		b := blocked(chamber, startY, x)
		if b != -1 {
			levels = append(levels, b)
		}
	}

	lowest := math.MaxInt
	for _, y := range levels {
		if y == -1 {
			return -1
		}
		if y < lowest {
			lowest = y
		}
	}
	return lowest
}

func blocked(chamber []string, startY, startX int) int {
	for y := startY; y < len(chamber); y++ {
		if chamber[startY][startX] == '#' {
			if chamber[y][startX] == '#' {
				return y
			}
		}
	}
	return -1
}

func simulateRock(chamber []string, jetInput string, rock []string) []string {
	air := []string{
		".......",
		".......",
		".......",
	}
	chamber = append(chamber, air...)
	chamber = append(chamber, rock...)
	//printChamber(chamber)

	for {
		direction := string(jetInput[jetIndex%len(jetInput)])
		jetIndex++
		isFinished, newChamber := simulatePushingAndFalling(chamber, direction)
		chamber = newChamber
		if !isFinished {
			return removeEmptyLines(newChamber)
		}
	}
}

func simulatePushingAndFalling(oldChamber []string, direction string) (isResting bool, newChamber []string) {
	pushingWasValid, pushedChamber := simulatePushing(oldChamber, direction)
	if pushingWasValid {
		oldChamber = pushedChamber
	}

	fallingWasValid, newChamber := simulateFalling(oldChamber)
	if fallingWasValid {
		return true, newChamber
	} else {
		return false, replaceAll(oldChamber, "@", "#")
	}
}

func simulatePushing(oldChamber []string, direction string) (isValid bool, newChamber []string) {
	newChamber = append(newChamber, oldChamber...)
	for y, line := range newChamber {
		var newLine string
		if direction == ">" {
			newLine = reverse(pushToTheLeft(reverse(line)))
		} else {
			newLine = pushToTheLeft(line)
		}
		if newLine == "" {
			return false, []string{}
		} else {
			newChamber[y] = newLine
		}
	}
	return true, newChamber
}

func pushToTheLeft(line string) string {
	for x := 0; x < maxLength-1; x++ {
		isMoveValid, newValue := move(rune(line[x]), rune(line[x+1]), x)
		if !isMoveValid {
			return ""
		}
		line = line[:x] + string(newValue) + line[x+1:]
	}

	if line[maxLength-1] == '@' {
		return line[:maxLength-1] + "."
	}
	return line
}

func simulateFalling(oldChamber []string) (isValid bool, newChamber []string) {
	if len(oldChamber) <= 1 {
		return false, []string{}
	}
	newChamber = append(newChamber, oldChamber...)

	for y := 0; y < len(newChamber)-1; y++ {
		onTop := newChamber[y+1]
		current := newChamber[y]
		for x := 0; x < maxLength; x++ {
			isMoveValid, newValue := move(rune(current[x]), rune(onTop[x]), y)
			if isMoveValid {
				newChamber[y] = newChamber[y][:x] + string(newValue) + newChamber[y][x+1:]
			} else {
				return false, []string{}
			}
		}
	}

	newChamber[len(newChamber)-1] = strings.ReplaceAll(newChamber[len(newChamber)-1], "@", ".")
	return true, newChamber
}

func move(current, next rune, currentPosition int) (isValid bool, newValue rune) {
	if currentPosition == 0 && current == '@' {
		return false, '0'
	}

	if next == '@' {
		if current == '.' || current == '@' {
			return true, '@'
		} else {
			return false, '0'
		}
	} else { // t == '.' || t == '#'
		if current == '#' {
			return true, current
		} else {
			return true, '.'
		}
	}
}

// util

func replaceAll(chamber2 []string, old, new string) []string {
	for i := range chamber2 {
		chamber2[i] = strings.ReplaceAll(chamber2[i], old, new)
	}
	return chamber2
}

func removeEmptyLines(chamber []string) (newChamber []string) {
	for _, line := range chamber {
		if line != "......." {
			newChamber = append(newChamber, line)
		}
	}
	return newChamber
}

func reverse(input string) (rev string) {
	for _, letter := range input {
		rev = string(letter) + rev
	}
	return rev
}

func printChamber(chamber []string) {
	fmt.Println("-------")
	for y := len(chamber) - 1; y >= 0; y-- {
		fmt.Println(chamber[y])
	}
}

func getDefaultRocks() [][]string {
	vertical := []string{
		"..@@@@.",
	}
	plus := []string{
		"...@...",
		"..@@@..",
		"...@...",
	}
	l := []string{
		"..@@@..",
		"....@..",
		"....@..",
	}
	horizontal := []string{
		"..@....",
		"..@....",
		"..@....",
		"..@....",
	}
	block := []string{
		"..@@...",
		"..@@...",
	}
	return [][]string{vertical, plus, l, horizontal, block}
}
