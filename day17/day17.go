package day17

import (
	"fmt"
	"strings"
)

type Position struct {
	x int
	y int
}

type Rock struct {
	positions []Position
}

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines[0]), 0
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

func solvePart1(jetInput string) int {
	var chamber []string
	jetIndex = 0

	for i := 0; i < 2022; i++ {
		rockNumber := i % 5
		rocks := getDefaultRocks()
		chamber = simulateRock(chamber, jetInput, rocks[rockNumber])
	}
	return len(chamber)
}

const maxLength = 7

var jetIndex int

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
		dir := string(jetInput[jetIndex%len(jetInput)])
		jetIndex++
		isFinished, newChamber := simFallingAndPushing(chamber, dir)
		chamber = newChamber
		if !isFinished {
			return removeEmptyLines(newChamber)
		}
	}
}

func printChamber(chamber []string) {
	fmt.Println("-------")
	for i := len(chamber) - 1; i >= 0; i-- {
		fmt.Println(chamber[i])
	}
}

func simFallingAndPushing(oldChamber []string, direction string) (bool, []string) {
	var newChamber []string
	pushingWasValid, pushedChamber := simPushing(oldChamber, direction)
	var fallingWasValid bool
	if pushingWasValid {
		oldChamber = pushedChamber
	}
	fallingWasValid, newChamber = simFalling(oldChamber)
	if fallingWasValid {
		return true, newChamber
	} else {
		return false, replaceAll(oldChamber, "@", "#")
	}
}

func replaceAll(chamber2 []string, old, new string) []string {
	for i := range chamber2 {
		chamber2[i] = strings.ReplaceAll(chamber2[i], old, new)
	}
	return chamber2
}

func simPushing(chamber2 []string, direction string) (isValid bool, newChamber []string) {
	newChamber = append(newChamber, chamber2...)
	for i, line := range newChamber {
		var newLine string
		if direction == ">" {
			newLine = reverse(simJetToLeft(reverse(line)))
		} else {
			newLine = simJetToLeft(line)
		}
		if newLine == "" {
			return false, []string{}
		} else {
			newChamber[i] = newLine
		}
	}

	return true, newChamber
}

func simJetToLeft(line string) string {
	for j := 0; j < maxLength-1; j++ {
		isValid, newC := move(rune(line[j]), rune(line[j+1]), j)
		if !isValid {
			return ""
		}
		line = line[:j] + string(newC) + line[j+1:]
	}

	if line[maxLength-1] == '@' {
		return line[:maxLength-1] + "."
	}

	return line
}

func simFalling(oldChamber []string) (isValid bool, newChamber []string) {
	if len(oldChamber) == 1 {
		return false, []string{}
	}

	newChamber = append(newChamber, oldChamber...)

	for y := 0; y < len(newChamber)-1; y++ {
		onTop := newChamber[y+1]
		current := newChamber[y]

		for x := 0; x < maxLength; x++ {
			isValid2, newC := move(rune(current[x]), rune(onTop[x]), y)
			if isValid2 {
				newChamber[y] = newChamber[y][:x] + string(newC) + newChamber[y][x+1:]
			} else {
				return false, []string{}
			}
		}
	}

	newChamber[len(newChamber)-1] = strings.ReplaceAll(newChamber[len(newChamber)-1], "@", ".")
	return true, newChamber
}

func removeEmptyLines(chamber []string) (newChamber []string) {
	counter := 0
	for _, line := range chamber {
		if line != "......." {
			newChamber = append(newChamber, line)
		} else {
			counter++
		}
	}
	// fmt.Println("Delted: ", (counter))
	return newChamber
}

func move(c, t rune, cPos int) (isValid bool, newC rune) {
	if cPos == 0 && c == '@' {
		return false, c
	}

	if t == '@' {
		if c == '.' || c == '@' {
			return true, '@'
		} else {
			return false, c
		}
	} else { // t == '.' || t == '#'
		if c == '#' {
			return true, c
		} else {
			return true, '.'
		}
	}
}

func reverse(input string) (rev string) {
	for _, letter := range input {
		rev = string(letter) + rev
	}
	return rev
}

func isRockResting() bool {
	return true
}
