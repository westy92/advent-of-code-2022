package main

import (
	"bufio"
	"fmt"
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

func day9() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	tailVisited := hashset.New()
	knots := 10
	knot := make([]Coords, knots, knots)
	prevState := make([]Coords, knots, knots)
	tailVisited.Add(knot[knots-1])

	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text[2:])
		switch text[0] {
		case 'L':
			for i := 0; i < num; i++ {
				prevTail := knot[knots-1]
				copy(prevState, knot)
				knot[0] = knot[0].Left()
				for s := 0; s < 1; s++ { // TODO remove
					if !knot[s+1].Touches(knot[s], true) {
						knot[s+1] = knot[s+1].Left()
						if !knot[s+1].Touches(knot[s], false) {
							if knot[s+1].Up().Touches(knot[s], false) {
								knot[s+1] = knot[s+1].Up()
							} else {
								knot[s+1] = knot[s+1].Down()
							}
						}
					}
				}
				// bubble rest
				/*for k := 1; k < knots; k++ {
					knot[k+1] = knot[k]
				}*/
				if knot[knots-1] != prevTail {
					tailVisited.Add(knot[knots-1])
				}
			}
			fmt.Println(knot)
		case 'R':
			for i := 0; i < num; i++ {
				prevTail := knot[knots-1]
				knot[0] = knot[0].Right()
				for s := 0; s < 1; s++ { // TODO remove
					if !knot[s+1].Touches(knot[s], true) {
						knot[s+1] = knot[s+1].Right()
						if !knot[s+1].Touches(knot[s], false) {
							if knot[s+1].Up().Touches(knot[s], false) {
								knot[s+1] = knot[s+1].Up()
							} else {
								knot[s+1] = knot[s+1].Down()
							}
						}
					}
				}
				if knot[knots-1] != prevTail {
					tailVisited.Add(knot[knots-1])
				}
			}
			fmt.Println(knot)
		case 'U':
			for i := 0; i < num; i++ {
				prevTail := knot[knots-1]
				knot[0] = knot[0].Up()
				for s := 0; s < 1; s++ { // TODO remove
					if !knot[s+1].Touches(knot[s], true) {
						knot[s+1] = knot[s+1].Up()
						if !knot[s+1].Touches(knot[s], false) {
							if knot[s+1].Left().Touches(knot[s], false) {
								knot[s+1] = knot[s+1].Left()
							} else {
								knot[s+1] = knot[s+1].Right()
							}
						}
					}
				}
				if knot[knots-1] != prevTail {
					tailVisited.Add(knot[knots-1])
				}
			}
			fmt.Println(knot)
		case 'D':
			for i := 0; i < num; i++ {
				prevTail := knot[knots-1]
				knot[0] = knot[0].Down()
				for s := 0; s < 1; s++ { // TODO remove
					if !knot[s+1].Touches(knot[s], true) {
						knot[s+1] = knot[s+1].Down()
						if !knot[s+1].Touches(knot[s], false) {
							if knot[s+1].Left().Touches(knot[s], false) {
								knot[s+1] = knot[s+1].Left()
							} else {
								knot[s+1] = knot[s+1].Right()
							}
						}
					}
				}
				if knot[knots-1] != prevTail {
					tailVisited.Add(knot[knots-1])
				}
			}
			fmt.Println(knot)
		}
	}

	println("Part 1: " + strconv.Itoa(tailVisited.Size()))

	fmt.Println(tailVisited)

	//println("Part 2: " + strconv.Itoa(part2))
}
