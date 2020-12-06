package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	eof := false
	for !eof {
		eof = !scanner.Scan()
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func part1(lines []string) int {
	sum := 0
	set := make(map[rune]interface{})
	for _, line := range lines {
		if len(line) == 0 {
			sum += len(set)
			set = make(map[rune]interface{})
		} else {
			for _, c := range line {
				set[c] = 1
			}
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	size := 0
	set := make(map[rune]int)
	for _, line := range lines {
		if len(line) == 0 {
			for _, v := range set {
				if v == size {
					sum++
				}
			}
			set = make(map[rune]int)
			size = 0
		} else {
			size++
			for _, c := range line {
				set[c]++
			}
		}
	}
	return sum
}

func main() {
	lines, _ := readFile("day06.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
