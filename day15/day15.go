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
	return solvePart1(sensors, 10), 0
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
		sensorRegex := util.FindStringSubmatch(line, `Sensor at x=(-*\d+), y=(-*\d+):`)
		sensorX := util.GetInt(sensorRegex[1])
		sensorY := util.GetInt(sensorRegex[2])

		beaconRegex := util.FindStringSubmatch(line, `closest beacon is at x=(-*\d+), y=(-*\d+)`)
		beaconX := util.GetInt(beaconRegex[1])
		beaconY := util.GetInt(beaconRegex[2])

		sensors = append(sensors, Sensor{Position{sensorX, sensorY}, Position{beaconX, beaconY}})
	}
	return sensors
}
