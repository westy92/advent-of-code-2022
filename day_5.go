package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func day5() {
	day5part1()
	day5part2()
}

func day5part1() {
	println("Part 1:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	stacks := []lls.Stack{*lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New()}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		for i := 0; i < len(text); i += 4 {
			block := text[i : i+3]
			if block[0] == '[' {
				stacks[i/4].Push(block[1])
			}
		}
	}

	// reverse stacks
	for i, stack := range stacks {
		newStack := lls.New()
		for !stack.Empty() {
			item, _ := stack.Pop()
			newStack.Push(item)
		}
		stacks[i] = *newStack
	}

	// directions
	regex := *regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		text := scanner.Text()
		res := regex.FindStringSubmatch(text)
		move, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])
		for i := 0; i < move; i++ {
			item, _ := stacks[from-1].Pop()
			stacks[to-1].Push(item)
		}
	}

	// print results
	for _, stack := range stacks {
		item, ok := stack.Peek()
		if ok {
			print(string(item.(uint8)))
		}
	}
	println()
}

func day5part2() {
	println("Part 2:")
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	stacks := []lls.Stack{*lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New(), *lls.New()}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		for i := 0; i < len(text); i += 4 {
			block := text[i : i+3]
			if block[0] == '[' {
				stacks[i/4].Push(block[1])
			}
		}
	}

	// reverse stacks
	for i, stack := range stacks {
		newStack := lls.New()
		for !stack.Empty() {
			item, _ := stack.Pop()
			newStack.Push(item)
		}
		stacks[i] = *newStack
	}

	// directions
	regex := *regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		text := scanner.Text()
		res := regex.FindStringSubmatch(text)
		move, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])
		intermediate := lls.New()
		for i := 0; i < move; i++ {
			item, _ := stacks[from-1].Pop()
			intermediate.Push(item)
		}
		for i := 0; i < move; i++ {
			item, _ := intermediate.Pop()
			stacks[to-1].Push(item)
		}
	}

	// print results
	for _, stack := range stacks {
		item, ok := stack.Peek()
		if ok {
			print(string(item.(uint8)))
		}
	}
	println()
}
