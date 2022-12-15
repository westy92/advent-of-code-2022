package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/emirpasic/gods/sets/hashset"
)

func day15() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	regex := *regexp.MustCompile(`Sensor at x=([+-]?\d+), y=([+-]?\d+): closest beacon is at x=([+-]?\d+), y=([+-]?\d+)`)
	crossSection := 2000000
	set := hashset.New()
	beaconRowSet := hashset.New()
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)
		sensorX, _ := strconv.Atoi(res[1])
		sensorY, _ := strconv.Atoi(res[2])
		beaconX, _ := strconv.Atoi(res[3])
		beaconY, _ := strconv.Atoi(res[4])
		if beaconY == crossSection {
			beaconRowSet.Add(beaconX)
		}
		distance := abs(sensorX-beaconX) + abs(sensorY-beaconY)
		println(distance)
		if crossSection <= sensorY {
			dY := sensorY - crossSection
			dX := (distance - dY) // at crossSection
			if dY < distance {
				for x := sensorX - dX; x <= sensorX+dX; x++ {
					set.Add(x)
				}
			}
		} else {
			dY := crossSection - sensorY
			dX := (distance - dY) // at crossSection
			if dY < distance {
				for x := sensorX - dX; x <= sensorX+dX; x++ {
					set.Add(x)
				}
			}
		}
	}

	println("Part 1: " + strconv.Itoa(set.Difference(beaconRowSet).Size()))

	println("Part 2: " + strconv.Itoa(1))
}
