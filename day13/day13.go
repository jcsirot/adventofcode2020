package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

func readFile(path string) (int, map[int]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, nil, err
	}
	lines := strings.Split(string(data), "\n")
	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}
	busLines := make(map[int]int)
	for i, l := range strings.Split(lines[1], ",") {
		if l != "x" {
			tmp, err := strconv.Atoi(l)
			if err != nil {
				return 0, nil, err
			}
			busLines[i] = tmp
		} else {
			busLines[i] = 0
		}
	}
	return timestamp, busLines, nil
}

func part1(timestamp int, schedule map[int]int) int {
	var lines []int
	for _, v := range schedule {
		if v != 0 {
			lines = append(lines, v)
		}
	}
	for i := 0; ; i++ {
		for _, line := range lines {
			if (timestamp+i)%line == 0 {
				return i * line
			}
		}
	}
}

func part2(timestamp int, schedule map[int]int) int {
	var lines []int
	var offsets []int
	m := 1
	for k, v := range schedule {
		if v != 0 {
			lines = append(lines, v)
			offsets = append(offsets, k)
			m *= v
		}
	}
	r := 0
	for i, line := range lines {
		mi := m / line
		y := big.NewInt(0).ModInverse(big.NewInt(int64(mi)), big.NewInt(int64(line)))
		r += (line - offsets[i]) * mi * int(y.Int64())
	}

	return r % m
}

func main() {
	timestamp, schedule, _ := readFile("day13.txt")
	fmt.Printf("Part 1: %d\n", part1(timestamp, schedule))
	fmt.Printf("Part 2: %d\n", part2(timestamp, schedule))
}
