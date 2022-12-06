package day6

func Solve(lines []string) (part1, parts2 int) {
	return findIndexOfDistinctLetterSequence(lines, 4), findIndexOfDistinctLetterSequence(lines, 14)
}

func findIndexOfDistinctLetterSequence(lines []string, length int) (indexOfLastLetter int) {
	for _, line := range lines {
		for i := 0; i < len(line)-length; i++ {
			if areDistinctLetters(line[i : i+length]) {
				return i + length
			}
		}
	}
	return -1
}

func areDistinctLetters(letters string) bool {
	for i := 0; i < len(letters)-1; i++ {
		if contains(letters[i+1:], letters[i]) {
			return false
		}
	}
	return true
}

func contains(target string, candidate byte) bool {
	for i := 0; i < len(target); i++ {
		if target[i] == candidate {
			return true
		}
	}
	return false
}
