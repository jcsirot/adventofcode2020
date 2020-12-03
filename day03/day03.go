package main

import (
	"bufio"
	"fmt"
	"os"
)

type tree struct {
	X int
	Y int
}

type world struct {
	Trees  []tree
	Width  int
	Height int
}

func (w world) isTree(x int, y int) bool {
	if x >= w.Width {
		x = x % w.Width
	}
	for _, t := range w.Trees {
		if t.X == x && t.Y == y {
			return true
		}
	}
	return false
}

func parseFile(path string) (*world, error) {
	w := &world{}
	w.Trees = make([]tree, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if y == 0 {
			w.Width = len(line)
		}
		for x, char := range scanner.Text() {
			if char == '#' {
				w.Trees = append(w.Trees, tree{X: x, Y: y})
			}
		}
		y++
	}
	w.Height = y
	return w, nil
}

func main() {
	w, _ := parseFile("day03.txt")
	fmt.Println(part1(*w, 3, 1))
	fmt.Println(part2(*w))
}

func part1(w world, rightSift int, downShift int) int {
	x := 0
	y := 0
	encountered := 0

	for y < w.Height {
		if w.isTree(x, y) {
			encountered++
		}
		x += rightSift
		y += downShift
	}

	return encountered
}

func part2(w world) int {
	return part1(w, 1, 1) * part1(w, 3, 1) * part1(w, 5, 1) * part1(w, 7, 1) * part1(w, 1, 2)
}
