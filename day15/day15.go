package day15

import (
	"adventofcode/util"
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

type Sensor struct {
	position      Position
	closestBeacon Position
}

func Solve(lines []string) (part1, part2 int) {
	sensors := parse(lines)
	return chris(sensors, 2000000), 0
}

func chris(sensors []Sensor, lineNumber int) int {
	minX, maxX := findBounds(sensors)
	lenght := maxX - minX + 1
	line := make([]bool, lenght)

	for _, sensor := range sensors {
		closestBeaconDistance := TaxicabDistance(sensor.position, sensor.closestBeacon)
		distanceToLine := int(math.Abs(float64(sensor.position.Y - lineNumber)))

		if distanceToLine <= closestBeaconDistance {
			from, to := getInterval(sensor.position.X, closestBeaconDistance-distanceToLine)

			if from < minX {
				line = append(make([]bool, minX-from), line...)
				minX = from
			}

			if maxX < to {
				line = append(line, make([]bool, to-maxX)...)
				maxX = to
			}

			for i := from - minX; i <= to-minX; i++ {
				line[i] = true
			}
		}
	}

	return countAndPrint(line) - countAlreadyKnownBeacons(sensors, lineNumber)
}

func countAndPrint(line []bool) (count int) {
	for _, b := range line {
		if b {
			count++
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	return count
}

func countAlreadyKnownBeacons(sensors []Sensor, y int) (count int) {
	var temp []Position
	for _, sensor := range sensors {
		if sensor.closestBeacon.Y == y && !contains(temp, sensor.closestBeacon) {
			temp = append(temp, sensor.closestBeacon)
		}
	}
	return len(temp)
}

func contains(positions []Position, candidate Position) bool {
	for _, position := range positions {
		if position == candidate {
			return true
		}
	}
	return false
}

func getInterval(x, d int) (from, to int) {
	return x - d, x + d
}

func findBounds(sensors []Sensor) (minX, maxX int) {
	minX = math.MaxInt
	maxX = math.MinInt
	for _, sensor := range sensors {
		currentX := sensor.position.X
		if currentX < minX {
			minX = currentX
		}
		if currentX > maxX {
			maxX = currentX
		}
	}
	return minX, maxX
}

func solvePart1(sensors []Sensor, lineNumber int) (part1 int) {
	for x := 0; x < 25; x++ {
		currentPosition := Position{x, lineNumber}

		if canBeaconBePresent(sensors, currentPosition) {
			part1++
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}

	return part1
}

func canBeaconBePresent(sensors []Sensor, position Position) bool {
	closestSensor, distanceToSensor := getClosestSensor(sensors, position)
	lastMin := TaxicabDistance(closestSensor.closestBeacon, closestSensor.position)
	return distanceToSensor > lastMin
}

func getClosestSensor(sensors []Sensor, position Position) (closestSensor Sensor, distance int) {
	distance = math.MaxInt
	for _, sensor := range sensors {
		distanceToPosition := TaxicabDistance(sensor.position, position)
		if distanceToPosition < distance {
			closestSensor = sensor
			distance = distanceToPosition
		}
	}
	return closestSensor, distance
}

func TaxicabDistance(from Position, to Position) int {
	return int(math.Abs(float64(from.X-to.X)) + math.Abs(float64(from.Y-to.Y)))
}

func parse(lines []string) (sensors []Sensor) {
	for _, line := range lines {
		sensorRegex := util.FindStringSubmatch(line, `Sensor at x=(-{0,1}\d+), y=(-{0,1}\d+):`)
		sensorX := util.GetInt(sensorRegex[1])
		sensorY := util.GetInt(sensorRegex[2])

		beaconRegex := util.FindStringSubmatch(line, `closest beacon is at x=(-{0,1}\d+), y=(-{0,1}\d+)`)
		beaconX := util.GetInt(beaconRegex[1])
		beaconY := util.GetInt(beaconRegex[2])

		sensors = append(sensors, Sensor{Position{sensorX, sensorY}, Position{beaconX, beaconY}})
	}
	return sensors
}
