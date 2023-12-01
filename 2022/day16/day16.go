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

type Path struct {
	path   string
	volume int
}

var valveMap map[string]Valve

func Solve(lines []string) (part1, part2 int) {
	parseValveMap(lines)
	calcValveDistances()
	printDistances()
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	paths := getAllPossiblePaths([]Valve{valveMap["AA"]}, 30, 0)
	return getMaxVolume(paths)
}

func solvePart2() (maxTotalVolume int) {
	paths := getAllPossiblePaths([]Valve{valveMap["AA"]}, 26, 0)

	for _, person := range paths {
		for _, elephant := range paths {
			if isDisjoint(person.path, elephant.path) {
				totalVolume := person.volume + elephant.volume
				if totalVolume > maxTotalVolume {
					maxTotalVolume = totalVolume
				}
			}
		}
	}
	return maxTotalVolume
}

func getMaxVolume(paths []Path) int {
	max := 0
	for _, path := range paths {
		if path.volume > max {
			max = path.volume
		}
	}
	return max
}

func getAllPossiblePaths(path []Valve, time int, volume int) (possiblePaths []Path) {
	if time <= 0 {
		return possiblePaths
	}

	current := path[len(path)-1]
	newVolume := volume + time*current.rate
	possiblePaths = append(possiblePaths, Path{pathToString(path), newVolume})

	for name, distance := range current.distances {
		next := valveMap[name]
		remainingTime := time - distance - 1
		if !contains(path, next) {
			possiblePaths = append(possiblePaths, getAllPossiblePaths(append(path, next), remainingTime, newVolume)...)
		}
	}
	return possiblePaths
}

func isDisjoint(pathNameA, pathNameB string) bool {
	for i := 0; i < len(pathNameA); i += 2 {
		currentPathName := pathNameA[i : i+2]
		if currentPathName != "AA" && strings.Contains(pathNameB, currentPathName) {
			return false
		}
	}
	return true
}

var traveresed map[string]bool

func calcValveDistances() {
	for _, valve := range valveMap {
		traveresed = make(map[string]bool)
		if valve.rate > 0 || valve.name == "AA" {
			breadthFirstCalcDistances([]Valve{valve}, valve, 0)
		}
	}
}

func breadthFirstCalcDistances(currentValves []Valve, origin Valve, distance int) {
	var possibleNextValves []Valve
	for _, current := range currentValves {
		if !traveresed[current.name] {
			possibleNextValves = append(possibleNextValves, setDistanceIfValid(current, origin, distance)...)
		}
	}
	if len(possibleNextValves) > 0 {
		breadthFirstCalcDistances(possibleNextValves, origin, distance+1)
	}
}

func setDistanceIfValid(current Valve, origin Valve, distance int) (nextValves []Valve) {
	traveresed[current.name] = true
	if origin.name != current.name {
		if current.rate > 0 {
			if isUseful(origin, current, distance) || origin.name == "AA" {
				valveMap[origin.name].distances[current.name] = distance
			}
		}
	}
	for _, tunnel := range current.tunnels {
		nextValves = append(nextValves, valveMap[tunnel])
	}
	return nextValves
}

func isUseful(origin Valve, current Valve, distance int) bool {
	return isABAGreaterThanBA(origin.rate, current.rate, distance)
}

func isABAGreaterThanBA(a, b, distance int) bool {
	return a*(2*distance+1)+b*distance > b*(1+distance)
}

// util

func contains(path []Valve, candiate Valve) bool {
	for _, valve := range path {
		if valve.name == candiate.name {
			return true
		}
	}
	return false
}

func pathToString(path []Valve) (s string) {
	for _, v := range path {
		s += v.name
	}
	return s
}

func printDistances() {
	for _, v := range valveMap {
		if v.rate > 0 || v.name == "AA" {
			fmt.Printf("%s: %v\n", v.name, v.distances)
		}
	}
}

func parseValveMap(lines []string) {
	valveMap = make(map[string]Valve)
	for _, line := range lines {
		regex := util.FindStringSubmatch(line, `Valve (\w\w) has flow rate=(\d+); tunnels* leads* to valves* ([\w, ]*)`)
		name := regex[1]
		rate := util.GetInt(regex[2])
		tunnels := strings.Split(regex[3], ", ")
		valveMap[name] = Valve{name: name, rate: rate, tunnels: tunnels, distances: make(map[string]int)}
	}
}
