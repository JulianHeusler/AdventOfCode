package day13

import (
	"adventofcode/util"
)

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(lines), 0
}

func solvePart1(lines []string) (sum int) {
	for i := 0; i < len(lines); i += 3 {
		leftPacket := lines[i]
		rightPacket := lines[i+1]
		if isInRightOrder(leftPacket, rightPacket) {
			pairNumber := (i / 3) + 1
			sum += pairNumber
		}
	}
	return sum
}

func isInRightOrder(left, right string) bool {
	for {
		if len(left) == 0 {
			return true
		}
		if len(right) == 0 {
			return false
		}

		l := left[0]
		r := right[0]
		if l == ']' && r != ']' {
			return true
		}
		if l != ']' && r == ']' {
			return false
		}

		if areEqualNonDigits(l, r) {
			left = left[1:]
			right = right[1:]
			continue
		}

		if IsDigit(l) && IsDigit(r) {
			leftNumber, leftNumberLength := getFirtNumber(left)
			rightNumber, rightNumberLength := getFirtNumber(right)
			if leftNumber < rightNumber {
				return true
			}
			if leftNumber > rightNumber {
				return false
			}
			left = left[leftNumberLength:]
			right = right[rightNumberLength:]
			continue
		}

		if IsDigit(l) && r == '[' {
			left = warpFirstNumberInBrackets(left)
			continue
		}
		if l == '[' && IsDigit(r) {
			right = warpFirstNumberInBrackets(right)
			continue
		}

		panic(left + " | " + right)
	}
}

func warpFirstNumberInBrackets(s string) string {
	_, numberLength := getFirtNumber(s)
	return "[" + s[0:numberLength] + "]" + s[numberLength:]
}

func areEqualNonDigits(l, r byte) bool {
	return !IsDigit(l) && !IsDigit(r) && l == r
}

func getFirtNumber(s string) (firstNumber, numberLength int) {
	for ; IsDigit(s[numberLength]) && numberLength < len(s); numberLength++ {
		firstNumber *= 10
		firstNumber += util.GetInt(string(s[numberLength]))
	}
	return firstNumber, numberLength
}

func IsDigit(d byte) bool {
	return '0' <= d && d <= '9'
}
