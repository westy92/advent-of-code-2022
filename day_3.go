package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	part1()
	part2()
}

func points(char rune) int {
	if 'a' <= char && char <= 'z' {
		return int(char - 'a' + 1)
	}
	if 'A' <= char && char <= 'Z' {
		return int(char - 'A' + 27)
	}
	return 0
}

func do1(text string) int {
	length := len(text)
	comp1 := text[:(length / 2)]
	comp2 := text[(length / 2):]
	set := hashset.New()
	for _, ch := range comp1 {
		set.Add(ch)
	}
	for _, ch := range comp2 {
		if set.Contains(ch) {
			return points(ch)
		}
	}
	return 0
}

func part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		total += do1(text)
	}
	fmt.Println("Score:", total)
}

func part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	total := 0
	i := 0
	hset := hashset.New()
	for scanner.Scan() {
		text := scanner.Text()
		set := hashset.New()
		for _, ch := range text {
			set.Add(ch)
		}
		switch i {
		case 0:
			hset = set
		case 1:
			hset = set.Intersection(hset)
		case 2:
			hset = set.Intersection(hset)
			total += points(hset.Values()[0].(rune))
		}
		i = (i + 1) % 3
	}
	fmt.Println("Score:", total)
}
