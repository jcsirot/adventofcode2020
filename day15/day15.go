package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseData(input string) ([]int, error) {
	tmp := strings.Split(input, ",")
	var numbers []int
	for _, i := range tmp {
		i, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, i)
	}
	return numbers, nil
}

func solve(input []int, rank int) int {
	numbers := make(map[int]int)
	var last int
	turn := 1
	for ; turn < len(input); turn++ {
		numbers[input[turn-1]] = turn
	}
	last = input[turn-1]
	turn++
	for ; turn <= rank; turn++ {
		t, found := numbers[last]
		numbers[last] = turn - 1
		if !found {
			last = 0
		} else {
			last = turn - 1 - t
		}
	}
	return last
}

func main() {
	input, _ := parseData("18,8,0,5,4,1,20")
	fmt.Printf("Part 1: %d\n", solve(input, 2020))
	fmt.Printf("Part 2: %d\n", solve(input, 30000000))
}
