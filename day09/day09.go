package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const preambleSize = 25

func parseFile(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var numbers []int
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func canBeSum(sum int, numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j]+numbers[i] == sum {
				return true
			}
		}
	}
	return false
}

func part1(numbers []int) int {
	for i := preambleSize; i < len(numbers); i++ {
		prev := numbers[i-preambleSize : i]
		if !canBeSum(numbers[i], prev) {
			return numbers[i]
		}
	}
	return -1
}

func part2(target int, numbers []int) int {
LOOP:
	for i := 0; i < len(numbers)-1; i++ {
		sum := numbers[i]
		smallest := numbers[i]
		largest := numbers[i]
		for j := 1; i+j < len(numbers); j++ {
			n := numbers[i+j]
			if n < smallest {
				smallest = n
			}
			if n > largest {
				largest = n
			}
			sum += n
			if sum > target {
				continue LOOP
			} else if sum == target {
				return smallest + largest
			}
		}
	}
	return -1
}

func main() {
	numbers, _ := parseFile("day09.txt")
	target := part1(numbers)
	fmt.Printf("Part 1: %d\n", target)
	fmt.Printf("Part 2: %d\n", part2(target, numbers))
}
