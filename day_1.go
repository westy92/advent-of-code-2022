package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func day1() {
	day1part1()
	day1part2()
}

func day1part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	max := 0
	elfTotal := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if elfTotal > max {
				max = elfTotal
			}
			elfTotal = 0
			continue
		}
		num, _ := strconv.Atoi(text)
		elfTotal += num
	}
	fmt.Println("Max:", max)
}

func day1part2() {
	println("Part 2:")
	num := 3
	queue := pq.NewWith(utils.IntComparator)

	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	elfTotal := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if queue.Size() < num {
				queue.Enqueue(elfTotal)
			} else if min, _ := queue.Peek(); elfTotal > min.(int) {
				queue.Dequeue()
				queue.Enqueue(elfTotal)
			}
			elfTotal = 0
			continue
		}
		num, _ := strconv.Atoi(text)
		elfTotal += num
	}
	grandTotal := 0
	for queue.Size() > 0 {
		num, _ := queue.Dequeue()
		grandTotal += num.(int)
	}
	fmt.Println("Total:", grandTotal)
}
