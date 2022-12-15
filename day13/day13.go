package day13

import (
	"adventofcode/util"
	"strings"
)

type Packet struct {
	values []int
	sub    []*Packet
}

func Solve(lines []string) (part1, part2 int) {
	parse(lines)
	return 0, 0
}

var packets []Packet

func parse(lines []string) {
	for i := 0; i < len(lines); i += 3 {
		packets = append(packets, *parseLine(lines[i]))
		packets = append(packets, *parseLine(lines[+1]))
	}
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
