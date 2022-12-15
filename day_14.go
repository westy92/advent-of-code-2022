package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	start Point
	end   Point
}

type Point struct {
	x int
	y int
}

func day14() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	lines := make([]Line, 0)

	minX := 500
	maxX := 500
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, " -> ")
		points := make([]Point, 0)
		for _, pair := range pairs {
			ints := strings.Split(pair, ",")
			x, _ := strconv.Atoi(ints[0])
			y, _ := strconv.Atoi(ints[1])
			minX = min(minX, x)
			maxX = max(maxX, x)
			maxY = max(maxY, y)
			points = append(points, Point{
				x: x,
				y: y,
			})
		}

		for i := range points {
			if i == 0 {
				continue
			}
			lines = append(lines, Line{
				start: points[i-1],
				end:   points[i],
			})
		}
	}

	width := min(maxX-minX+1, 2*(maxY+1+2))
	grid := make([][]bool, width)

	for i := range grid {
		grid[i] = make([]bool, maxY+1+2)
	}

	// fill grid with lines
	for _, line := range lines {
		if line.start.x == line.end.x {
			for y := min(line.start.y, line.end.y); y <= max(line.start.y, line.end.y); y++ {
				grid[line.start.x-minX][y] = true
			}
		} else if line.start.y == line.end.y {
			for x := min(line.start.x, line.end.x); x <= max(line.start.x, line.end.x); x++ {
				grid[x-minX][line.start.y] = true
			}
		}
	}

	part1 := fillSandPart1(minX, maxY, grid, maxX)
	println("Part 1: " + strconv.Itoa(part1))

	println("Part 2: " + strconv.Itoa(1))
}

func fillSandPart1(minX int, maxY int, grid [][]bool, maxX int) int {
	count := 0
	for {
		curX := 500 - minX
		curY := 0
		for {
			if curY+1 > maxY {
				return count
			} else if !grid[curX][curY+1] {
				curY++
			} else if curY+1 > maxY || curX-1 < 0 {
				return count
			} else if !grid[curX-1][curY+1] {
				curX--
				curY++
			} else if curY+1 > maxY || curX+1 > maxX-minX {
				return count
			} else if !grid[curX+1][curY+1] {
				curX++
				curY++
			} else if curY > maxY || curX < 0 || curX > maxX-minX {
				return count
			} else {
				grid[curX][curY] = true
				count++
				break
			}
		}
	}
}
