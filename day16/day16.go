package day16

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

type Valve struct {
	name      string
	rate      int
	tunnels   []string
	distances map[string]int
}

func Solve(lines []string) (part1, part2 int) {
	valves := parse(lines)
	return solvePart1(valves), 0
}

func solvePart1(valves []Valve) int {
	initValveMap(valves)
	calcDistances(valves)
	printDistances()
	//fmt.Println(calcPaths())
	start()

	return getMaxVolume()
}

func getMaxVolume() int {
	max := 0
	for _, volume := range volumes {
		if volume > max {
			max = volume
		}
	}
	return max
}

func start() {
	//for _, tunnel := range valvesMap["AA"].tunnels {
	//	traveresed = make(map[string]bool)
	//	breadthFirstPaths22([]Valve{valvesMap["AA"], valvesMap[tunnel]}, 30, 0)
	//}

	breadthFirstPaths22([]Valve{valvesMap["AA"]}, 30, 0)
}

var volumes []int

func breadthFirstPaths22(path []Valve, time int, volume int) {
	test := pathToString(path)
	if test != "" {

	}

	if time <= 0 {
		printPath(path)
		volumes = append(volumes, volume)
		return
	}

	current := path[len(path)-1]

	for name, distance := range current.distances {
		next := valvesMap[name]
		remaining := time - distance - 1
		newVolume := time * current.rate
		if !contains(path, next) {
			breadthFirstPaths22(append(path, next), remaining, volume+newVolume)
		}
	}

	volumes = append(volumes, volume+time*current.rate)
}

func printPath(path []Valve) {
	fmt.Println(pathToString(path))
}

func pathToString(path []Valve) (s string) {
	for _, v := range path {
		s += v.name
	}
	return s
}

func printDistances() {
	for _, v := range valvesMap {
		if v.rate > 0 || v.name == "AA" {
			fmt.Printf("%s: %v\n", v.name, v.distances)
		}
	}
}

func initValveMap(valves []Valve) {
	valvesMap = make(map[string]Valve)
	for _, valve := range valves {
		valvesMap[valve.name] = valve
	}
}

var valvesMap map[string]Valve

func calcDistances(valves []Valve) {
	for _, valve := range valves {
		traveresed = make(map[string]bool)
		if valve.rate > 0 || valve.name == "AA" {
			breadthFirstSearch([]Valve{valve}, valve, 0)
		}
	}
}

func calcPaths() (x []int) {
	for _, tunnel := range valvesMap["AA"].tunnels {
		x = append(x, breadthFirstPaths([]Valve{valvesMap["AA"], valvesMap[tunnel]}, 30))
	}
	return x
}

func temp2(current []Valve) (path []Valve) {
	if len(current)+1 == 6 {
		return path
	}

	for _, tunnel := range current[len(current)-1].tunnels {
		t := valvesMap[tunnel]
		if !contains(current, t) {
			path = append(path, temp2(append(current, t))...)
		}
	}
	return path
}

func breadthFirstPaths(currentValves []Valve, depth int) (pressure int) {
	if depth <= 0 {
		return pressure
	}

	var possibleNextValves []Valve
	for _, current := range currentValves {
		if !traveresed[current.name] {
			possibleNextValves = append(possibleNextValves, temp(current, depth)...)
		}
	}
	if len(possibleNextValves) > 0 {
		breadthFirstPaths(possibleNextValves, depth-1)
	}

	return pressure
}

func temp(current Valve, depth int) (nextValves []Valve) {
	traveresed[current.name] = true

	for _, tunnel := range current.tunnels {
		nextValves = append(nextValves, valvesMap[tunnel])
	}
	return nextValves
}

func contains(path []Valve, candiate Valve) bool {
	for _, valve := range path {
		if valve.name == candiate.name {
			return true
		}
	}
	return false
}

func breadthFirstSearch(currentValves []Valve, origin Valve, distance int) {
	var possibleNextValves []Valve
	for _, current := range currentValves {
		if !traveresed[current.name] {
			possibleNextValves = append(possibleNextValves, getNextValves(current, origin, distance)...)
		}
	}
	if len(possibleNextValves) > 0 {
		breadthFirstSearch(possibleNextValves, origin, distance+1)
	}
}

var traveresed map[string]bool

func getNextValves(current Valve, origin Valve, distance int) (nextValves []Valve) {
	traveresed[current.name] = true

	if origin.name != current.name {
		if current.rate > 0 {
			if lohntEsSich(origin, current, distance) || origin.name == "AA" {
				valvesMap[origin.name].distances[current.name] = distance
			}
		}
	}

	for _, tunnel := range current.tunnels {
		nextValves = append(nextValves, valvesMap[tunnel])
	}
	return nextValves
}

func lohntEsSich(origin Valve, current Valve, distance int) bool {
	return compare(origin.rate, current.rate, distance)
}

func compare(a, b, distance int) bool {
	return a*(2*distance+1)+b*distance > b*(1+distance)
}

func parse(lines []string) (valves []Valve) {
	for _, line := range lines {
		regex := util.FindStringSubmatch(line, `Valve (\w\w) has flow rate=(\d+); tunnels* leads* to valves* ([\w, ]*)`)
		name := regex[1]
		rate := util.GetInt(regex[2])
		tunnels := strings.Split(regex[3], ", ")
		valves = append(valves, Valve{name: name, rate: rate, tunnels: tunnels, distances: make(map[string]int)})
	}
	return valves
}
