package day11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items               []int
	operation           string
	divisibilityTest    int
	targetTrue          int
	targetFalse         int
	inspectedItemsCount int
}

func Solve(lines []string) (part1, part2 int) {
	monkeys := parseMonkeys(lines)
	fmt.Println(monkeys)

	for round := 0; round < 20; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				newWorryLevel := executeOperation(m.operation, item)
				newWorryLevel = int(newWorryLevel / 3)
				if newWorryLevel%m.divisibilityTest == 0 {
					monkeys[m.targetTrue].items = append(monkeys[m.targetTrue].items, newWorryLevel)
				} else {
					monkeys[m.targetFalse].items = append(monkeys[m.targetFalse].items, newWorryLevel)
				}
				monkeys[i].inspectedItemsCount++
			}
			monkeys[i].items = []int{}
		}
	}

	return getResult(monkeys), 0
}

func getResult(monkeys []monkey) int {
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItemsCount > monkeys[j].inspectedItemsCount
	})

	return monkeys[0].inspectedItemsCount * monkeys[1].inspectedItemsCount
}

func parseMonkeys(lines []string) (monkeys []monkey) {
	for lineNumber := 0; lineNumber < len(lines); lineNumber += 7 {
		newMonkey := monkey{
			items:            parseStartingItems(lines[lineNumber+1]),
			operation:        parseOperation(lines[lineNumber+2]),
			divisibilityTest: parseDivisibilityTest(lines[lineNumber+3]),
			targetTrue:       parseTargetTrue(lines[lineNumber+4]),
			targetFalse:      parseTargetFalse(lines[lineNumber+5]),
		}
		monkeys = append(monkeys, newMonkey)
	}

	return monkeys
}

func parseStartingItems(line string) (startingItems []int) {
	regexStartingItems := regexp.MustCompile(`Starting items: (.+)`)
	items := strings.Split(regexStartingItems.FindStringSubmatch(line)[1], ", ")
	for _, item := range items {
		startingItems = append(startingItems, getInt(item))
	}
	return startingItems
}

func parseOperation(line string) string {
	regex := regexp.MustCompile(`Operation: (.+)`)
	return regex.FindStringSubmatch(line)[1]
}

func parseDivisibilityTest(line string) int {
	regex := regexp.MustCompile(`Test: divisible by (\d+)`)
	return getInt(regex.FindStringSubmatch(line)[1])
}

func parseTargetTrue(line string) int {
	regex := regexp.MustCompile(`If true: throw to monkey (\d+)`)
	return getInt(regex.FindStringSubmatch(line)[1])
}

func parseTargetFalse(line string) int {
	regex := regexp.MustCompile(`If false: throw to monkey (\d+)`)
	return getInt(regex.FindStringSubmatch(line)[1])
}

func executeOperation(operation string, worryLevel int) int {
	regex := regexp.MustCompile(`new = old (.) (\d+)`)
	match := regex.FindStringSubmatch(operation)

	if len(match) != 0 {
		operator := match[1]
		value := getInt(match[2])
		if operator == "*" {
			return worryLevel * value
		} else {
			return worryLevel + value
		}
	}

	regex2 := regexp.MustCompile(`new = old (.) old`)
	match2 := regex2.FindStringSubmatch(operation)
	operator := match2[1]

	if operator == "*" {
		return worryLevel * worryLevel
	} else {
		return worryLevel + worryLevel
	}
}

func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
