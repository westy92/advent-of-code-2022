package main

import (
	"bufio"
	"os"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	part1()
	part2()
}

func containsDuplicates(s string) bool {
	set := hashset.New()
	for _, ch := range s {
		set.Add(ch)
	}

	return len(s) != set.Size()
}

func part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		for i, _ := range text {
			if i < 4 || containsDuplicates(text[(i-4):i]) {
				continue
			} else {
				println(i)
				break
			}
		}

	}
}

func part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		for i, _ := range text {
			if i < 14 || containsDuplicates(text[(i-14):i]) {
				continue
			} else {
				println(i)
				break
			}
		}

	}
}
