package day21

import (
	"adventofcode/util"
	"regexp"
	"strings"
)

type OperationMonkey struct {
	name      string
	operation string
}

func Solve(lines []string) (part1 int, part2 int) {
	operationMonkeys, numberMonkeys := parse(lines)
	return solvePart1(operationMonkeys, numberMonkeys), 0
}

func solvePart1(operationMonkeys []OperationMonkey, numberMonkeys map[string]int) int {
	index := 0
	for len(operationMonkeys) != 0 {
		current := operationMonkeys[index]
		isExecutable, result := tryOperation(current.operation, numberMonkeys)

		if isExecutable {
			numberMonkeys[current.name] = result
			operationMonkeys = remove(operationMonkeys, current)
			if index > 0 {
				index--
			}
		} else {
			index = (index + 1) % len(operationMonkeys)
		}
	}
	return numberMonkeys["root"]
}

func remove(operationMonkeys []OperationMonkey, operationMonkey OperationMonkey) []OperationMonkey {
	var index int
	for i, monkey := range operationMonkeys {
		if monkey == operationMonkey {
			index = i
			break
		}
	}

	if len(operationMonkeys) == index {
		return operationMonkeys[:index]
	}
	return append(operationMonkeys[:index], operationMonkeys[index+1:]...)
}

func tryOperation(operation string, numberMonkeys map[string]int) (bool, int) {
	regex := util.FindStringSubmatch(operation, `(\w+) (\+|-|\/|\*) (\w+)`)
	if len(regex) == 0 {
		return false, 0
	}

	first := numberMonkeys[regex[1]]
	second := numberMonkeys[regex[3]]
	if first == 0 || second == 0 {
		return false, 0
	}

	switch regex[2] {
	case "+":
		return true, first + second
	case "-":
		return true, first - second
	case "/":
		return true, first / second
	case "*":
		return true, first * second
	default:
		return false, 0
	}
}

func parse(lines []string) (operationMonkeys []OperationMonkey, numberMonkeys map[string]int) {
	numberMonkeys = map[string]int{}
	for _, line := range lines {
		splitted := strings.Split(line, ": ")
		name := splitted[0]
		number := regexp.MustCompile(`\d+`).FindString(splitted[1])

		if number != "" {
			numberMonkeys[name] = util.GetInt(number)
		} else {
			operationMonkeys = append(operationMonkeys, OperationMonkey{name, splitted[1]})
		}
	}
	return operationMonkeys, numberMonkeys
}
