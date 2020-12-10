package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseFile(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ints []int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func part1(adapters []int) int {
	sort.Ints(adapters)
	diff1 := 0
	diff3 := 1
	for i, v := range adapters {
		prev := 0
		if i > 0 {
			prev = adapters[i-1]
		}
		if v-prev == 1 {
			diff1++
		} else if v-prev == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

func count(adapters []int, counts map[int]int) int {
	if len(adapters) == 1 {
		return 1
	}
	idx := len(adapters) - 1
	last := adapters[idx]
	v, ok := counts[last]
	if ok {
		return v
	}
	cnt := 0
	for i := 1; i <= 3; i++ {
		if idx-i < 0 {
			continue
		}
		n := adapters[idx-i]
		if last-n > 3 {
			continue
		}
		c := count(adapters[0:idx-i+1], counts)
		cnt += c
		counts[n] = c
	}
	return cnt
}

func part2(adapters []int) int {
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	counts := make(map[int]int)
	return count(adapters, counts)
}

func main() {
	adapters, _ := parseFile("day10.txt")
	fmt.Printf("Part 1: %d\n", part1(adapters))
	fmt.Printf("Part 2: %d\n", part2(adapters))
}
