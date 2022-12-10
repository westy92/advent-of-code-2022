package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
)

func coordsToInt(a int, b int) int {
	return a*100_000 + b
}

func getItem(grid *arraylist.List, r int, c int) int {
	col, _ := grid.Get(c)
	col2 := col.(*arraylist.List)
	item, _ := col2.Get(r)
	itm, _ := item.(int)
	return itm
}

func day8() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var grid = arraylist.New()
	made := false
	for scanner.Scan() {
		text := scanner.Text()
		if !made {
			for i := 0; i < len(text); i++ {
				grid.Add(arraylist.New())
			}
			made = true
		}
		for i, ch := range text {
			num, _ := strconv.Atoi(string(ch))
			col, _ := grid.Get(i)
			col.(*arraylist.List).Add(num)
		}
	}

	width := grid.Size()
	corner, _ := grid.Get(0)
	height := corner.(*arraylist.List).Size()
	visible := hashset.New()
	// top
	for c := 0; c < width; c++ {
		prevMax := -1
		for r := 0; r < height; r++ {
			itm := getItem(grid, r, c)
			if itm > prevMax {
				prevMax = itm
				visible.Add(coordsToInt(r, c))
			}
		}
	}
	// right
	for r := 0; r < height; r++ {
		prevMax := -1
		for c := 0; c < width; c++ {
			itm := getItem(grid, r, c)
			if itm > prevMax {
				prevMax = itm
				visible.Add(coordsToInt(r, c))
			}
		}
	}

	// bottom
	for c := 0; c < width; c++ {
		prevMax := -1
		for r := height - 1; r > 0; r-- {
			itm := getItem(grid, r, c)
			if itm > prevMax {
				prevMax = itm
				visible.Add(coordsToInt(r, c))
			}
		}
	}
	// left
	for r := height - 1; r > 0; r-- {
		prevMax := -1
		for c := width - 1; c > 0; c-- {
			itm := getItem(grid, r, c)
			if itm > prevMax {
				prevMax = itm
				visible.Add(coordsToInt(r, c))
			}
		}
	}

	println("Part 1: " + strconv.Itoa(visible.Size()))

	// Part 2
	// permute all trees (ignore edges)
	best := 0
	for c := 1; c < width-1; c++ {
		for r := 1; r < height-1; r++ {
			itm := getItem(grid, r, c)

			top := 1
			for r2 := r - 1; r2 > 0; r2-- {
				itm2 := getItem(grid, r2, c)
				if itm2 >= itm {
					break
				} else {
					top++
				}
			}

			left := 1
			for c2 := c - 1; c2 > 0; c2-- {
				itm2 := getItem(grid, r, c2)
				if itm2 >= itm {
					break
				} else {
					left++
				}
			}

			right := 1
			for c2 := c + 1; c2 < width; c2++ {
				itm2 := getItem(grid, r, c2)
				if itm2 >= itm {
					break
				} else {
					right++
				}
			}

			bottom := 0
			for r2 := r + 1; r2 < height; r2++ {
				itm2 := getItem(grid, r2, c)
				if itm2 >= itm {
					break
				} else {
					bottom++
				}
			}

			subscore := top * bottom * left * right
			if subscore > best {
				best = subscore
			}
		}
	}

	println("Part 2: " + strconv.Itoa(best))
}
