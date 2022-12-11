package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func day10() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	x := 1
	cycles := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "addx ") {
			cycles = append(cycles, x)
			cycles = append(cycles, x)
			param := text[5:]
			num, _ := strconv.Atoi(param)
			x += num
		} else if text == "noop" {
			cycles = append(cycles, x)
		}
	}

	cycle := 20
	result := 0
	for i, item := range cycles {
		if i+1 == cycle {
			result += (i + 1) * item
			cycle += 40
		}
	}

	println("Part 1: " + strconv.Itoa(result))

	//fmt.Printf("%v\n", cycles)

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
