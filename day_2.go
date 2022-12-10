package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day2() {
	day2part1()
	day2part2()
}

var m = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
	"A": "R",
	"B": "P",
	"C": "S",
}

func determineOutcome(p1 string, p2 string) int {
	play1 := m[p1]
	play2 := m[p2]
	if play1 == play2 {
		return pts[play2] + 3
	} else if win[play1] == play2 {
		return pts[play2] + 6
	} else {
		return pts[play2]
	}
}

var lose = map[string]string{
	"R": "S",
	"P": "R",
	"S": "P",
}
var win = map[string]string{
	"S": "R",
	"R": "P",
	"P": "S",
}
var pts = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

func determineOutcome2(p1 string, end string) int {
	if end == "X" {
		return pts[lose[m[p1]]]
	} else if end == "Y" {
		return pts[m[p1]] + 3
	} else {
		return pts[win[m[p1]]] + 6
	}
}

func day2part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		plays := strings.Split(text, " ")
		total += determineOutcome(plays[0], plays[1])
	}
	fmt.Println("Score:", total)
}

func day2part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		plays := strings.Split(text, " ")
		total += determineOutcome2(plays[0], plays[1])
	}
	fmt.Println("Score:", total)
}
