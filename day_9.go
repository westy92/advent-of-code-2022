package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/emirpasic/gods/sets/hashset"
)

type Coords struct {
	Row    int
	Column int
}

func (c Coords) Touches(other Coords, includeDiagonal bool) bool {
	return c == other || // same spot
		c == other.Up() || c == other.Down() || // same col
		c == other.Left() || c == other.Right() || // same row
		// diagonal
		includeDiagonal && (c == other.Up().Left() ||
			c == other.Up().Right() ||
			c == other.Down().Left() ||
			c == other.Down().Right())
}

func (c Coords) Up() Coords {
	c.Row--
	return c
}

func (c Coords) Down() Coords {
	c.Row++
	return c
}

func (c Coords) Left() Coords {
	c.Column--
	return c
}

func (c Coords) Right() Coords {
	c.Column++
	return c
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	tailVisited := hashset.New()
	head := Coords{
		Row:    0,
		Column: 0,
	}
	tail := Coords{
		Row:    0,
		Column: 0,
	}
	tailVisited.Add(head)

	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text[2:])
		switch text[0] {
		case 'L':
			for i := 0; i < num; i++ {
				head = head.Left()
				if !tail.Touches(head, true) {
					tail = tail.Left()
					if !tail.Touches(head, false) {
						if tail.Up().Touches(head, false) {
							tail = tail.Up()
						} else {
							tail = tail.Down()
						}
					}
					tailVisited.Add(tail)
				}
			}
			println()
		case 'R':
			for i := 0; i < num; i++ {
				head = head.Right()
				if !tail.Touches(head, true) {
					tail = tail.Right()
					if !tail.Touches(head, false) {
						if tail.Up().Touches(head, false) {
							tail = tail.Up()
						} else {
							tail = tail.Down()
						}
					}
					tailVisited.Add(tail)
				}
			}
			println()
		case 'U':
			for i := 0; i < num; i++ {
				head = head.Up()
				if !tail.Touches(head, true) {
					tail = tail.Up()
					if !tail.Touches(head, false) {
						if tail.Right().Touches(head, false) {
							tail = tail.Right()
						} else {
							tail = tail.Left()
						}
					}
					tailVisited.Add(tail)
				}
			}
			println()
		case 'D':
			for i := 0; i < num; i++ {
				head = head.Down()
				if !tail.Touches(head, true) {
					tail = tail.Down()
					if !tail.Touches(head, false) {
						if tail.Right().Touches(head, false) {
							tail = tail.Right()
						} else {
							tail = tail.Left()
						}
					}
					tailVisited.Add(tail)
				}
			}
			println()
		}

	}

	println("Part 1: " + strconv.Itoa(tailVisited.Size()))
	//println(tailVisited)
	//s, _ := json.MarshalIndent(tailVisited, "", "\t")
	//fmt.Print(string(s))

	//println("Part 2: " + strconv.Itoa(part2))
}
