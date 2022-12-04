package reader

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(day int, rootDir bool) (lines []string) {
	var filepath string
	if rootDir {
		filepath = fmt.Sprintf("day%d/input.txt", day)
	} else {
		filepath = fmt.Sprintf("../day%d/input.txt", day)
	}

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
