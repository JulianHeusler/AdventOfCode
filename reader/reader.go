package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadExampleInput(day int) (lines []string) {
	return readFile(getFilePath(day, "example"))
}

func ReadInput(day int) (lines []string) {
	return readFile(getFilePath(day, "input"))
}

func getFilePath(day int, filename string) string {
	return fmt.Sprintf("%s\\day%d\\%s.txt", getProjectDirectoryPath(), day, filename)
}

func getProjectDirectoryPath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(currentFile), "../")
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
