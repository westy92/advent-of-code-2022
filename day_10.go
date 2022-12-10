package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type History struct {
	startCycle int
	endCycle   int
	value      int
}

func main2() {
	day10()
}

func day10() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	x := 1
	cycle := 0
	history := make([]*History, 0)
	cycles := make([]int, 0)
	history = append(history, &History{
		startCycle: 0,
		endCycle:   0,
		value:      x,
	})
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "addx ") {
			cycles = append(cycles, x)
			cycles = append(cycles, x)
			cycle += 2
			last := history[len(history)-1]
			last.endCycle = cycle
			param := text[5:]
			num, _ := strconv.Atoi(param)
			x += num
			history = append(history, &History{
				startCycle: cycle,
				endCycle:   0,
				value:      x,
			})
		} else if text == "noop" {
			// listing files, will pick up contents in other lines
			cycle += 1
			cycles = append(cycles, x)
		} else {
			// nothing else is supported
		}
	}

	// close loop? last end = cycle
	history[len(history)-1].endCycle = cycle

	cycle = 20
	result := 0
	for _, item := range history {
		if item.startCycle <= cycle && cycle <= item.endCycle {
			result += cycle * item.value
			cycle += 40
		}
	}

	println("Part 1: " + strconv.Itoa(result))

	//fmt.Printf("%v\n", cycles)

	cycle = 1
	for i, item := range cycles {
		crtPosition := i % 40
		if item-1 <= crtPosition && crtPosition <= item+1 {
			print("X")
		} else {
			print(".")
		}
		if crtPosition == 39 {
			println()
		}
	}
}
