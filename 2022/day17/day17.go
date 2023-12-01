package day17

import (
	"fmt"
	"strings"
)

const (
	maxLength       = 7
	loopStartOffset = 2000
	minLoopLength   = 100
)

var jetIndex int

func Solve(lines []string) (part1, part2 int) {
	return calculateRockTowerHeight(lines[0], 2022), calculateRockTowerHeight(lines[0], 1000000000000)
}

func calculateRockTowerHeight(jetInput string, rounds int) (score int) {
	var chamber []string
	jetIndex = 0
	rocks := getDefaultRocks()
	rockCount := len(rocks)
	oldHeight := 0
	var loop []int

	for i := 0; i < rounds; i++ {
		if i%rockCount == 1 {
			oldHeight = score + len(chamber)
		}

		chamber = simulateRock(chamber, jetInput, rocks[i%rockCount])
		chamber, score = trimChamber(chamber, score)

		if i > loopStartOffset && i%rockCount == 0 {
			delta := score + len(chamber) - oldHeight
			loop = append(loop, delta)
		}
		if i > loopStartOffset+minLoopLength && i%(2*rockCount) == 0 {
			if isRepeating(loop) {
				loopDelta := 0
				for _, v := range loop {
					loopDelta += v
				}
				x := (rounds - i) / (len(loop) * 5)
				score += x * loopDelta
				i += x * len(loop) * 5
				fmt.Println(i)
			}
		}
	}
	return score + len(chamber)
}

func trimChamber(chamber []string, score int) ([]string, int) {
	if len(chamber) > 200 {
		return chamber[150:], score + 150
	}
	return chamber, score
}

func isRepeating(list []int) bool {
	if len(list)%2 != 0 {
		return false
	}
	mid := len(list) / 2
	for i := 0; i < mid; i++ {
		if list[i] != list[mid+i] {
			return false
		}
	}
	return true
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
