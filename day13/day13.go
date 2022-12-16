package day13

import (
	"adventofcode/util"
	"math"
	"strings"
)

type Packet struct {
	values []int
	sub    []*Packet
}

type Bracket struct {
	value int
	depth int
}

func Solve(lines []string) (part1, part2 int) {
	packets := parse2(lines)
	return solvePart1(packets), 0
}

func solvePart1(packets [][]Bracket) (part1 int) {
	for i := 0; i < len(packets); i += 2 {
		if rightOrder2(packets[i], packets[i+1]) {
			part1 += i + 1
		}
	}
	return part1
}

func rightOrder2(left, right []Bracket) bool {
	x := int(math.Min(float64(len(left)), float64(len(right))))
	for i := 0; i < x; i++ {
		currL := left[i]
		currR := right[i]
		if left[i].depth == right[i].depth {
			if left[i].value > right[i].value {
				return false
			}
		} else {
			if currL.depth < currR.depth {
				currL.depth++
				rightOrder2(append([]Bracket{currL}, left[i:]...), right[i:])
			} else {
				currR.depth++
				rightOrder2(left[i:], append([]Bracket{currR}, right[i:]...))
			}
			break
		}
	}

	return true
}

func rightOrder(left, right []Bracket) bool {
	i := 0
	if left[i].depth == right[i].depth {
		if left[i].value > right[i].value {
			return false
		}
	} else if left[i].depth > right[i].depth {
		return false
	}

	if rightOrder(left[i+1:], right[i+1:]) {
		return true
	}

	return true
}

var packets []Packet

func parse(lines []string) {
	for i := 0; i < len(lines); i += 3 {
		packets = append(packets, *parseLine(lines[i]))
		packets = append(packets, *parseLine(lines[+1]))
	}
}

func parse2(lines []string) (pakets [][]Bracket) {
	for _, line := range lines {
		if line != "" {
			pakets = append(pakets, parseLine2(line))
		}
	}
	return pakets
}

func parseLine2(line string) (packet []Bracket) {
	currentDepth := 0
	for i := 0; i < len(line)-1; i++ {
		current := line[i]
		next := line[i+1]
		if current == '[' {
			currentDepth++
			if next == ']' {
				packet = append(packet, Bracket{-1, currentDepth})
			}
		} else if current == ']' {
			currentDepth--
		} else if current != ',' {
			packet = append(packet, Bracket{util.GetInt(string(current)), currentDepth})
		}
	}
	return packet
}

func parseLine(line string) (p *Packet) {
	p = &Packet{}
	s := util.FindStringSubmatch(line, `\[([^\]]*)\]`)
	if len(s) != 0 {
		p.sub = append(p.sub, parseLine(s[1]))
		return p
	}

	ints := strings.Split(line, ",")
	for _, v := range ints {
		p.values = append(p.values, util.GetInt(v))
	}
	return p
}
