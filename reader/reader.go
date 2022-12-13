package reader

import (
	"fmt"
	"os"
	"strings"
)

func ReadExampleInput(day int) (lines []string) {
	return readFile(getFilePath(day, false, "example"))
}

func ReadInput(day int, isRootDir bool) (lines []string) {
	return readFile(getFilePath(day, isRootDir, "input"))
}

func getFilePath(day int, isRootDir bool, filename string) string {
	if isRootDir {
		return fmt.Sprintf("day%d/%s.txt", day, filename)
	} else {
		return fmt.Sprintf("../day%d/%s.txt", day, filename)
	}
}

func readFile(filepath string) (lines []string) {
	input, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	parsedInput := strings.ReplaceAll(string(input), "\r", "")

	if parsedInput[len(parsedInput)-1:] == "\n" {
		return strings.Split(parsedInput[:len(parsedInput)-1], "\n")
	}
	return strings.Split(parsedInput, "\n")
}
