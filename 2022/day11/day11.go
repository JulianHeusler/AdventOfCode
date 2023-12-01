package day11

import (
	"adventofcode/util"
	"sort"
)

type Monkey struct {
	items               []int
	operation           Operation
	dividesBy           int
	targetTrue          int
	targetFalse         int
	inspectedItemsCount int
}

type Operation func(input int) int

func Solve(lines []string) (part1, part2 int) {
	monkeys := parseMonkeys(lines)
	monkeysCopy := make([]Monkey, len(monkeys))
	copy(monkeysCopy, monkeys)

	return simulateMonkeyInTheMiddle(monkeys, 20,
			func(level int) int {
				return int(level / 3)
			}),
		simulateMonkeyInTheMiddle(monkeysCopy, 10000,
			func(level int) int {
				return level % calcLeastCommonMultiple(monkeys)
			})
}

func simulateMonkeyInTheMiddle(monkeys []Monkey, rounds int, reduceWorryLevel Operation) int {
	for round := 0; round < rounds; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newWorryLevel := monkey.operation(item)
				newWorryLevel = reduceWorryLevel(newWorryLevel)
				if newWorryLevel%monkey.dividesBy == 0 {
					monkeys[monkey.targetTrue].items = append(monkeys[monkey.targetTrue].items, newWorryLevel)
				} else {
					monkeys[monkey.targetFalse].items = append(monkeys[monkey.targetFalse].items, newWorryLevel)
				}
				monkeys[i].inspectedItemsCount++
			}
			monkeys[i].items = []int{}
		}
	}
	return calcMonkeyBusiness(monkeys)
}

func calcLeastCommonMultiple(monkeys []Monkey) int {
	kgv := 1
	for _, monkey := range monkeys {
		kgv *= monkey.dividesBy
	}
	return kgv
}

func calcMonkeyBusiness(monkeys []Monkey) int {
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItemsCount > monkeys[j].inspectedItemsCount
	})
	return monkeys[0].inspectedItemsCount * monkeys[1].inspectedItemsCount
}

func parseMonkeys(lines []string) (monkeys []Monkey) {
	for lineNumber := 0; lineNumber < len(lines); lineNumber += 7 {
		newMonkey := Monkey{
			items:       util.FindIntSlice(lines[lineNumber+1], `Starting items: (.+)`, ", "),
			operation:   parseOperation(util.FindStringSubmatch(lines[lineNumber+2], `Operation: (.+)`)[1]),
			dividesBy:   util.FindFirstInt(lines[lineNumber+3], `Test: divisible by (\d+)`),
			targetTrue:  util.FindFirstInt(lines[lineNumber+4], `If true: throw to monkey (\d+)`),
			targetFalse: util.FindFirstInt(lines[lineNumber+5], `If false: throw to monkey (\d+)`),
		}
		monkeys = append(monkeys, newMonkey)
	}
	return monkeys
}

func parseOperation(operationString string) Operation {
	match := util.FindStringSubmatch(operationString, `new = old (.) (\d+)`)
	if match != nil {
		operator := match[1]
		value := util.GetInt(match[2])
		if operator == "*" {
			return func(worryLevel int) int { return worryLevel * value }
		} else {
			return func(worryLevel int) int { return worryLevel + value }
		}
	}
	match = util.FindStringSubmatch(operationString, `new = old (.) old`)
	operator := match[1]
	if operator == "*" {
		return func(worryLevel int) int { return worryLevel * worryLevel }
	} else {
		return func(worryLevel int) int { return worryLevel + worryLevel }
	}
}
