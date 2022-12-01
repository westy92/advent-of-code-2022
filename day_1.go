package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
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

func part2() {
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
			min, _ := queue.Peek()
			intMin, _ := min.(int)
			if queue.Size() < num {
				queue.Enqueue(elfTotal)
			} else if min != nil && elfTotal > intMin {
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
		intNum, _ := num.(int)
		grandTotal += intNum
	}
	fmt.Println("Total:", grandTotal)
}
