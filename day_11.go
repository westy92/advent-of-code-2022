package main

import (
	"sort"
	"strconv"
)

type Monkey struct {
	Items        []int
	Operation    func(x int) int
	Test         func(x int) int
	InspectCount int
}

func day11() {
	monkeys := []Monkey{
		{
			Items: []int{91, 58, 52, 69, 95, 54},
			Operation: func(x int) int {
				return x * 13
			},
			Test: func(x int) int {
				if x%7 == 0 {
					return 1
				} else {
					return 5
				}
			},
		},
		{
			Items: []int{80, 80, 97, 84},
			Operation: func(x int) int {
				return x * x
			},
			Test: func(x int) int {
				if x%3 == 0 {
					return 3
				} else {
					return 5
				}
			},
		},
		{
			Items: []int{86, 92, 71},
			Operation: func(x int) int {
				return x + 7
			},
			Test: func(x int) int {
				if x%2 == 0 {
					return 0
				} else {
					return 4
				}
			},
		},
		{
			Items: []int{96, 90, 99, 76, 79, 85, 98, 61},
			Operation: func(x int) int {
				return x + 4
			},
			Test: func(x int) int {
				if x%11 == 0 {
					return 7
				} else {
					return 6
				}
			},
		},
		{
			Items: []int{60, 83, 68, 64, 73},
			Operation: func(x int) int {
				return x * 19
			},
			Test: func(x int) int {
				if x%17 == 0 {
					return 1
				} else {
					return 0
				}
			},
		},
		{
			Items: []int{96, 52, 52, 94, 76, 51, 57},
			Operation: func(x int) int {
				return x + 3
			},
			Test: func(x int) int {
				if x%5 == 0 {
					return 7
				} else {
					return 3
				}
			},
		},
		{
			Items: []int{75},
			Operation: func(x int) int {
				return x + 5
			},
			Test: func(x int) int {
				if x%13 == 0 {
					return 4
				} else {
					return 2
				}
			},
		},
		{
			Items: []int{83, 75},
			Operation: func(x int) int {
				return x + 1
			},
			Test: func(x int) int {
				if x%19 == 0 {
					return 2
				} else {
					return 6
				}
			},
		},
	}

	for i := 0; i < 20; i++ {
		for curMonkey, monkey := range monkeys {
			for _, item := range monkey.Items {
				newItem := monkey.Operation(item) / 3
				newMonkey := monkey.Test(newItem)
				monkeys[newMonkey].Items = append(monkeys[newMonkey].Items, newItem)
				monkeys[curMonkey].InspectCount++
			}
			monkeys[curMonkey].Items = []int{}
		}
	}

	InspectCounts := make([]int, len(monkeys))
	for _, monkey := range monkeys {
		InspectCounts = append(InspectCounts, monkey.InspectCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(InspectCounts)))

	println("Part 1: " + strconv.Itoa(InspectCounts[0]*InspectCounts[1]))
}
