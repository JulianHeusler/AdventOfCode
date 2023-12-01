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
	position         Position
	closestBeacon    Position
	distanceToBeacon int
}

const THREADS = 4

var maxSize int

func Solve(lines []string, lineNumber int, maximumSize int) (part1 int, part2 int) {
	sensors := parse(lines)
	maxSize = maximumSize
	return chris(sensors, lineNumber), solvePart2(sensors)
}

var cave []bool

func solvePart2(sensors []Sensor) int {
	sensorsPerThread := len(sensors) / THREADS
	channel := make(chan int)

	for y := 0; y < len(sensors); y += sensorsPerThread {
		var curr []Sensor
		if y+sensorsPerThread >= len(sensors) { // last iteration
			curr = sensors[y:]
		} else {
			curr = sensors[y : y+sensorsPerThread]
		}
		go findDistressBeaconParalell(channel, sensors, curr)
	}

	return <-channel
}

func findDistressBeaconParalell(channel chan int, allSensors []Sensor, sensorsToCheck []Sensor) {
	for _, sensor := range sensorsToCheck {
		candidates := getDiamondOutline(sensor.position, sensor.distanceToBeacon+1)
		for _, candidate := range candidates {
			if isInBounds(candidate) {
				if isDistressBeacon(allSensors, candidate) {
					channel <- candidate.X*4000000 + candidate.Y
				}
			}
		}
	}
}

func getDiamondOutline(origin Position, radius int) (outlinePositions []Position) {
	for y := origin.Y - radius; y < origin.Y+radius; y++ {
		distanceFromOrigin := TaxicabDistance(origin, Position{origin.X, y})
		offsetX := radius - distanceFromOrigin
		if offsetX == 0 {
			outlinePositions = append(outlinePositions, Position{origin.X, y})
		} else {
			outlinePositions = append(outlinePositions, Position{origin.X - offsetX, y})
			outlinePositions = append(outlinePositions, Position{origin.X + offsetX, y})
		}
	}
	return outlinePositions
}

func isInBounds(position Position) bool {
	return 0 <= position.X && position.X <= maxSize &&
		0 <= position.Y && position.Y <= maxSize
}

func isDistressBeacon(sensors []Sensor, current Position) bool {
	for _, sensor := range sensors {
		distanceToSensor := TaxicabDistance(current, sensor.position)
		if current == sensor.closestBeacon {
			return false
		}
		if distanceToSensor <= sensor.distanceToBeacon {
			return false
		}
	}
	return true
}

func solvePart3(sensors []Sensor, maxSize int) int64 {
	n := maxSize / THREADS

	channel := make(chan int64)
	for y := 0; y < maxSize; y += n {
		curr := createIntSlice(y, y+n)
		if y+n == maxSize {
			curr = append(curr, y+n)
		}
		go para(channel, curr, sensors)
	}

	i := <-channel
	return i
}

func createIntSlice(from, to int) (slice []int) {
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}
	return slice
}




func para(myChannel chan int64, yValues []int, sensors []Sensor) {
	for _, y := range yValues {
		if y%100000 == 0 {
			fmt.Println(y)
		}
		for x := 0; x <= maxSize; x++ {
			currentPosition := Position{x, y}
			if isDistressBeacon(sensors, currentPosition) {
				myChannel <- int64(currentPosition.X)*int64(4000000) + int64(currentPosition.Y)
			}
		}
	}
}



func solvePart2Old(sensors []Sensor, maxSize int) int64 {
	cave = make([]bool, int(math.Pow(float64(maxSize+1), 2)))

	for _, sensor := range sensors {
		drawDiamond(sensor.position, sensor.distanceToBeacon, maxSize)
	}

	for _, sensor := range sensors {
		drawItem(sensor.position, 'S', maxSize)
		drawItem(sensor.closestBeacon, 'B', maxSize)
	}

	return findDistressBeacon(maxSize)
}

func findDistressBeacon(maxSize int) (tuningFrequency int64) {
	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			if !cave[y*(maxSize+1)+x] {
				return int64(x)*int64(4000000) + int64(y)
			}
		}
	}
	return -1
}

func drawItem(position Position, item rune, maxSize int) {
	if 0 <= position.X && position.X <= maxSize &&
		0 <= position.Y && position.Y <= maxSize {
		cave[position.Y*(maxSize+1)+position.X] = true
	}
}

func drawDiamond(origin Position, radius int, maxSize int) {
	for y := max(origin.Y-radius, 0); y <= min(origin.Y+radius, maxSize); y++ {
		for x := max(origin.X-radius, 0); x <= min(origin.X+radius, maxSize); x++ {
			if TaxicabDistance(Position{x, y}, origin) <= radius {
				cave[y*(maxSize+1)+x] = true
			}
		}
	}
}

func min(value, min int) int {
	return int(math.Min(float64(value), float64(min)))
}

func max(value, max int) int {
	return int(math.Max(float64(value), float64(max)))
}

func chris(sensors []Sensor, lineNumber int) int {
	minX, maxX := findBounds(sensors)
	lenght := maxX - minX + 1
	line := make([]bool, lenght)

	for _, sensor := range sensors {
		distanceToLine := int(math.Abs(float64(sensor.position.Y - lineNumber)))

		if distanceToLine <= sensor.distanceToBeacon {
			from, to := getInterval(sensor.position.X, sensor.distanceToBeacon-distanceToLine)

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

	return count(line) - countAlreadyKnownBeacons(sensors, lineNumber)
}

func count(line []bool) (count int) {
	for _, b := range line {
		if b {
			count++
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
	return distanceToSensor > closestSensor.distanceToBeacon
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
		sensorPosition := Position{sensorX, sensorY}

		beaconRegex := util.FindStringSubmatch(line, `closest beacon is at x=(-{0,1}\d+), y=(-{0,1}\d+)`)
		beaconX := util.GetInt(beaconRegex[1])
		beaconY := util.GetInt(beaconRegex[2])
		beaconPosition := Position{beaconX, beaconY}

		sensors = append(sensors, Sensor{sensorPosition, beaconPosition, TaxicabDistance(sensorPosition, beaconPosition)})
	}
	return sensors
}
