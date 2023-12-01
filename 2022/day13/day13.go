package day13

import (
	"adventofcode/util"
)

func Solve(lines []string) (part1, part2 int) {
	packets := parse(lines)
	return solvePart1(packets), solvePart2(packets)
}

func solvePart1(packets []string) (sum int) {
	for i := 0; i < len(packets); i += 2 {
		leftPacket := packets[i]
		rightPacket := packets[i+1]
		if isInRightOrder(leftPacket, rightPacket) {
			pairNumber := (i / 2) + 1
			sum += pairNumber
		}
	}
	return sum
}

func solvePart2(packets []string) int {
	smallerThan2 := 1
	smallerThan6 := 2
	for _, packet := range packets {
		if isInRightOrder(packet, "[[2]]") {
			smallerThan2++
			smallerThan6++
		} else if isInRightOrder(packet, "[[6]]") {
			smallerThan6++
		}
	}
	return smallerThan2 * smallerThan6
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

func parse(lines []string) (packets []string) {
	for _, line := range lines {
		if line != "" {
			packets = append(packets, line)
		}
	}
	return packets
}
