package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func findPieceEnd(s string) int {
	var braces int = 0
	for i, ch := range s {
		if ch == ',' && braces == 0 {
			return i
		} else if ch == '[' {
			braces++
		} else if ch == ']' {
			braces--
		}
	}
	if braces == 0 {
		return len(s)
	}

	return -1
}

const (
	valid   = 1
	next    = 0
	invalid = -1
)

func comparePackets(a string, b string) int {
	if strings.HasPrefix(a, "[") && strings.HasSuffix(a, "]") {
		a = a[1 : len(a)-1]
	}
	if strings.HasPrefix(b, "[") && strings.HasSuffix(b, "]") {
		b = b[1 : len(b)-1]
	}
	firstEnd := findPieceEnd(a)
	secondEnd := findPieceEnd(b)
	first := a[:firstEnd]
	second := b[:secondEnd]
	isInt := regexp.MustCompile(`^\d+$`)

	for {
		if first == "" && second != "" {
			return valid
		} else if first != "" && second == "" {
			return invalid
		} else if first == "" && second == "" {
			return next
		}
		isFirstInt := isInt.Match([]byte(first))
		isSecondInt := isInt.Match([]byte(second))
		if isFirstInt && isSecondInt {
			firstInt, _ := strconv.Atoi(first)
			secondInt, _ := strconv.Atoi(second)
			if firstInt < secondInt {
				return valid
			} else if firstInt > secondInt {
				return invalid
			}
		} else if !isFirstInt && !isSecondInt {
			result := comparePackets(first, second)
			if result != next {
				return result
			}
		} else if isFirstInt && !isSecondInt {
			result := comparePackets("["+first+"]", second)
			if result != next {
				return result
			}
		} else if !isFirstInt && isSecondInt {
			result := comparePackets(first, "["+second+"]")
			if result != next {
				return result
			}
		}
		if firstEnd == len(a) && secondEnd == len(b) {
			return next
		} else if firstEnd == len(a) {
			return valid
		} else if secondEnd == len(b) {
			return invalid
		}
		oldFirstEnd := firstEnd
		oldSecondEnd := secondEnd
		firstEnd += findPieceEnd(a[firstEnd+1:]) + 1
		secondEnd += findPieceEnd(b[secondEnd+1:]) + 1
		first = a[oldFirstEnd+1 : firstEnd]
		second = b[oldSecondEnd+1 : secondEnd]
	}
}

func day13() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	result := 0
	i := 1
	var packets []string
	packets = append(packets, "[[2]]", "[[6]]")
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		packets = append(packets, line1, line2)
		scanner.Scan() // blank
		if comparePackets(line1, line2) == valid {
			result += i
		}
		i++
	}

	println("Part 1: " + strconv.Itoa(result))

	sort.Slice(packets, func(i, j int) bool {
		return comparePackets(packets[i], packets[j]) > 0
	})

	result = 1
	for i, packet := range packets {
		if packet == "[[2]]" || packet == "[[6]]" {
			result *= i + 1
		}
	}

	println("Part 2: " + strconv.Itoa(result))
}
