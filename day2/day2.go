package day2

import "log"

func Solve(lines []string) (part1, part2 int) {
	for _, round := range lines {
		result, token := calcScore(string(round[0]), string(round[2]))
		part1 += result + token
		part2 += getScoreForExpectedResult(string(round[0]), string(round[2]))
	}
	return part1, part2
}

func getScoreForExpectedResult(input1, input2 string) int {
	expectedResult := convertResultToScore(input2)

	resultY, tokenY := calcScore(input1, "Y")
	if expectedResult == resultY {
		return resultY + tokenY
	}
	resultX, tokenX := calcScore(input1, "X")
	if expectedResult == resultX {
		return resultX + tokenX
	}
	resultZ, tokenZ := calcScore(input1, "Z")
	if expectedResult == resultZ {
		return resultZ + tokenZ
	}

	log.Println("Error")
	return -1
}

func calcScore(enemy string, player string) (result, token int) {
	valueEnemy := convertToValue(enemy)
	valuePlayer := convertToValue(player)

	if valueEnemy == valuePlayer {
		return 3, valuePlayer // draw
	}

	if valueEnemy == (valuePlayer%3)+1 {
		return 0, valuePlayer // loss
	}

	return 6, valuePlayer // win
}

func convertResultToScore(s string) int {
	switch s {
	case "X":
		return 0 // loss
	case "Y":
		return 3 // draw
	case "Z":
		return 6 // win
	default:
		log.Println("Error")
		return -1
	}
}

func convertToValue(letter string) int {
	switch letter {
	case "X":
		fallthrough
	case "A":
		return 1 // rock
	case "Y":
		fallthrough
	case "B":
		return 2 // paper
	case "Z":
		fallthrough
	case "C":
		return 3 // scissors
	default:
		log.Println("Error")
		return -1
	}
}
