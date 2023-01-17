package day13

import (
	"adventofcode/util"
)

func Solve(lines []string) (part1, part2 int) {
	return solvePart1(parse(lines)), 0
}

func solvePart1(packets []string) (sum int) {
	for i := 0; i < len(packets); i += 3 {
		left := packets[i]
		right := packets[i+1]
		if isRightOrder(left, right) {
			pairNumber := (i / 3) + 1
			sum += pairNumber
		}
	}
	return sum
}

func isRightOrder(left, right string) bool {
	for {
		if len(left) == 0 {
			return true
		}
		if len(right) == 0 {
			return false
		}

		if left[0] == ']' && right[0] != ']' {
			return true
		}
		if left[0] != ']' && right[0] == ']' {
			return false
		}

		if areEqualNonDigits(left[0], right[0]) {
			left = left[1:]
			right = right[1:]
		} else if IsDigit(left[0]) && IsDigit(right[0]) {
			leftNumber, leftLength := getFirtNumber(left)
			rightNumber, rightLength := getFirtNumber(right)
			if leftNumber < rightNumber {
				return true
			}
			if leftNumber > rightNumber {
				return false
			}
			left = left[leftLength:]
			right = right[rightLength:]
		} else if IsDigit(left[0]) && right[0] == '[' {
			left = warpFirstNumberInBrackets(left)
		} else if left[0] == '[' && IsDigit(right[0]) {
			right = warpFirstNumberInBrackets(right)
		} else {
			panic(left + " | " + right)
		}
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

func IsDigit(r byte) bool {
	return '0' <= r && r <= '9'
}

func parse(lines []string) (packets []string) {
	for _, line := range lines {
		packets = append(packets, line)
	}
	return packets
}
