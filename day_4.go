package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
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
		pairs := strings.Split(text, ",")
		nums1 := strings.Split(pairs[0], "-")
		nums2 := strings.Split(pairs[1], "-")
		a1, _ := strconv.Atoi(nums1[0])
		a2, _ := strconv.Atoi(nums1[1])
		b1, _ := strconv.Atoi(nums2[0])
		b2, _ := strconv.Atoi(nums2[1])
		if a1 <= b1 && a2 >= b2 || b1 <= a1 && b2 >= a2 {
			total++
		}
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
	for scanner.Scan() {
		text := scanner.Text()
		pairs := strings.Split(text, ",")
		nums1 := strings.Split(pairs[0], "-")
		nums2 := strings.Split(pairs[1], "-")
		a1, _ := strconv.Atoi(nums1[0])
		a2, _ := strconv.Atoi(nums1[1])
		b1, _ := strconv.Atoi(nums2[0])
		b2, _ := strconv.Atoi(nums2[1])
		if a1 <= b1 && a2 >= b1 || b1 <= a2 && b2 >= a2 || b1 <= a1 && b2 >= a1 || a1 <= b2 && a2 >= b2 {
			total++
		}
	}
	fmt.Println("Score:", total)
}
