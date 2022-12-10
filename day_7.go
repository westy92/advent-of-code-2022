package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/hashmap"
)

const (
	File      = 0
	Directory = 1
)

type Item struct {
	Parent   *Item
	Type     int // file/dir
	Size     int
	Children *hashmap.Map
}

func day7() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	root := Item{
		Type:     Directory,
		Children: hashmap.New(),
	}
	pwd := &root
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "$ cd ") {
			dir := text[5:]
			if dir == "/" {
				pwd = &root
			} else if dir == ".." {
				pwd = pwd.Parent
			} else {
				item, found := pwd.Children.Get(dir)
				if found && item.(*Item).Type == Directory {
					pwd, _ = item.(*Item)
				}
			}
		} else if text == "$ ls" {
			// listing files, will pick up contents in other lines
		} else {
			if strings.HasPrefix(text, "dir ") {
				dir := text[4:]
				pwd.Children.Put(dir, &Item{
					Parent:   pwd,
					Type:     Directory,
					Children: hashmap.New(),
				})
			} else {
				info := strings.Split(text, " ")
				size, _ := strconv.Atoi(info[0])
				name := info[1]
				pwd.Children.Put(name, &Item{
					Parent: pwd,
					Type:   File,
					Size:   size,
				})
			}
		}
	}

	var dirSizes []int
	calculateDirSizes(&root, &dirSizes)
	part1 := 0
	for _, dirSize := range dirSizes {
		if dirSize < 100000 {
			part1 += dirSize
		}
	}
	println("Part 1: " + strconv.Itoa(part1))

	neededSpace := 30_000_000 - (70_000_000 - root.Size)
	smallest := math.MaxInt
	for _, dirSize := range dirSizes {
		if dirSize >= neededSpace && dirSize < smallest {
			smallest = dirSize
		}
	}
	println("Part 2: " + strconv.Itoa(smallest))
}

func calculateDirSizes(root *Item, dirSizes *[]int) int {
	sizes := 0
	for _, child := range root.Children.Values() {
		chld, _ := child.(*Item)
		if chld.Type == File {
			sizes += chld.Size
		} else {
			dirSize := calculateDirSizes(chld, dirSizes)
			sizes += dirSize
			*dirSizes = append(*dirSizes, dirSize)
		}
	}
	root.Size = sizes
	return sizes
}
