package main

import (
	"bufio"
	"os"
	"strconv"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

type Cell struct {
	Row      int
	Column   int
	Distance int
}

func findStarts(matrix [][]rune, includeAs bool) []Cell {
	var cells []Cell
	for i, row := range matrix {
		for j, ch := range row {
			if ch == 'S' || includeAs && ch == 'a' {
				cells = append(cells, Cell{Row: i, Column: j})
			}
		}
	}
	return cells
}

func isValid(row int, col int, matrix [][]rune, visited [][]bool, current rune) bool {
	valid := 0 <= row && 0 <= col &&
		row < len(matrix) && col < len(matrix[0]) &&
		!visited[row][col]
	if !valid {
		return false
	}
	if current == 'S' {
		current = 'a'
	}
	checking := matrix[row][col]
	if checking == 'E' {
		checking = 'z'
	}
	return checking <= current+1
}

func findShortestPath(matrix [][]rune, includeAs bool) int {
	// find starting point(s)
	starting := findStarts(matrix, includeAs)

	var results []int
	for _, start := range starting {
		// init visited matrix
		visited := make([][]bool, len(matrix))
		for i := range visited {
			visited[i] = make([]bool, len(matrix[i]))
		}

		queue := llq.New()
		queue.Enqueue(start)
		visited[start.Row][start.Column] = true

		for !queue.Empty() {
			item, _ := queue.Dequeue()
			cell, _ := item.(Cell)
			if matrix[cell.Row][cell.Column] == 'E' {
				results = append(results, cell.Distance)
				break
			}

			current := matrix[cell.Row][cell.Column]

			// up
			if isValid(cell.Row-1, cell.Column, matrix, visited, current) {
				queue.Enqueue(Cell{Row: cell.Row - 1, Column: cell.Column, Distance: cell.Distance + 1})
				visited[cell.Row-1][cell.Column] = true
			}
			// down
			if isValid(cell.Row+1, cell.Column, matrix, visited, current) {
				queue.Enqueue(Cell{Row: cell.Row + 1, Column: cell.Column, Distance: cell.Distance + 1})
				visited[cell.Row+1][cell.Column] = true
			}
			// left
			if isValid(cell.Row, cell.Column-1, matrix, visited, current) {
				queue.Enqueue(Cell{Row: cell.Row, Column: cell.Column - 1, Distance: cell.Distance + 1})
				visited[cell.Row][cell.Column-1] = true
			}
			// right
			if isValid(cell.Row, cell.Column+1, matrix, visited, current) {
				queue.Enqueue(Cell{Row: cell.Row, Column: cell.Column + 1, Distance: cell.Distance + 1})
				visited[cell.Row][cell.Column+1] = true
			}
		}
	}

	min := results[0]
	for _, result := range results {
		if result < min {
			min = result
		}
	}
	return min
}

func day12() {

	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// read lines
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// build matrix
	matrix := make([][]rune, len(lines))
	for i := range matrix {
		matrix[i] = make([]rune, len(lines[i]))
	}

	// populate matrix
	for i, line := range lines {
		for j, ch := range line {
			matrix[i][j] = ch
		}
	}

	shortestPath := findShortestPath(matrix, false)

	println("Part 1: " + strconv.Itoa(shortestPath))

	shortestPath2 := findShortestPath(matrix, true)

	println("Part 2: " + strconv.Itoa(shortestPath2))
}
