package day5

import (
	"regexp"
	"strconv"
)

var stacks [][]string

type instruction struct {
	amount int
	from   int
	to     int
}

func Solve(lines []string) (part1, part2 string) {
	InitStacks((len(lines[0]) + 1) / 4)

	stackLineIndex := 0
	for _, line := range lines {
		stackLineIndex++
		if line == "" {
			break
		}
	}

	stackLines := lines[:stackLineIndex-2] // remove filler lines
	instructions := parseInstructions(lines[stackLineIndex:])

	return solvePart1(stackLines, instructions), solvePart2(stackLines, instructions)
}

func solvePart1(stackLines []string, instructions []instruction) string {
	parseStacks(stackLines)
	executeInstructionsSolo(instructions)
	return readTopLine()
}

func solvePart2(stackLines []string, instructions []instruction) string {
	parseStacks(stackLines)
	executeInstructionsStacked(instructions)
	return readTopLine()
}

func parseStacks(lines []string) {
	for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]
		for i := 1; i < len(line); i += 4 {
			val := line[i : i+1]
			if val != " " {
				stacks[i/4] = append(stacks[i/4], val)
			}
		}
	}
}

func executeInstructionsSolo(instructions []instruction) {
	for _, instruction := range instructions {
		for i := 0; i < instruction.amount; i++ {
			moveCrate(instruction.from, instruction.to)
		}
	}
}

func moveCrate(from, to int) {
	crate := Pop(from)
	Push(to, crate)
}

func executeInstructionsStacked(instructions []instruction) {
	for _, instruction := range instructions {
		var movedCrates []string
		for i := 0; i < instruction.amount; i++ {
			movedCrates = append(movedCrates, Pop(instruction.from))
		}

		for i := len(movedCrates) - 1; i >= 0; i-- {
			Push(instruction.to, movedCrates[i])
		}
	}
}

func readTopLine() (topLine string) {
	for stackIndex := range stacks {
		topLine += Peek(stackIndex + 1)
	}
	return topLine
}

// stacks

func InitStacks(stackCount int) {
	stacks = make([][]string, stackCount)
}

func Pop(stackNumber int) (removedCrate string) {
	stackNumber--
	topOfStack := len(stacks[stackNumber]) - 1
	removedCrate = stacks[stackNumber][topOfStack]
	stacks[stackNumber] = stacks[stackNumber][:topOfStack]
	return removedCrate
}

func Push(stackNumber int, crate string) {
	stackNumber--
	stacks[stackNumber] = append(stacks[stackNumber], crate)
}

func Peek(stackNumber int) (topCrate string) {
	stackNumber--
	topOfStack := len(stacks[stackNumber]) - 1
	if topOfStack < 0 {
		return ""
	}
	return stacks[stackNumber][topOfStack]
}

func parseInstructions(lines []string) (instructions []instruction) {
	for _, line := range lines {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

func parseInstruction(line string) instruction {
	r := regexp.MustCompile(`move (?P<amount>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	s := r.FindStringSubmatch(line)

	return instruction{
		amount: stringToInt(s[1]),
		from:   stringToInt(s[2]),
		to:     stringToInt(s[3]),
	}
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
