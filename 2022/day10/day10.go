package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(lines []string) (sumSigalStrength int, crt string) {
	cpu := parseCPUStack(lines)
	x := 1

	for cycleIndex, instruction := range cpu {
		cycleNumber := cycleIndex + 1
		if IsSignalStrengthCycle(cycleNumber) {
			sigalStrength := x * (cycleNumber)
			sumSigalStrength += sigalStrength
		}

		if isVisible(cycleIndex%40, x) {
			crt += "#"
		} else {
			crt += "."
		}
		x += instruction
	}

	print(crt)
	return sumSigalStrength, crt
}

func IsSignalStrengthCycle(cycleNumber int) bool {
	return (cycleNumber+20)%40 == 0
}

func isVisible(position, sprite int) bool {
	return sprite-1 <= position && position <= sprite+1
}

func parseCPUStack(lines []string) (cpu []int) {
	for _, line := range lines {
		cpu = append(cpu, 0) // for each noop and addx
		if line != "noop" {
			number, _ := strconv.Atoi(strings.Split(line, " ")[1])
			cpu = append(cpu, number) // add addx value
		}
	}
	return cpu
}

func print(crt string) {
	for i := 0; i < len(crt); i += 40 {
		fmt.Println(crt[i : i+40])
	}
}
