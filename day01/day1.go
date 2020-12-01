package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("part1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	expenses := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, v)
	}

	part1, err := Part1(expenses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1)

	part2, err := Part2(expenses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", part2)
}

// Part1 solve Day 1 Part 1
func Part1(expenses []int) (int, error) {
	sort.Ints(expenses)
	i := 0
	j := len(expenses) - 1

	for i < j {
		sum := expenses[i] + expenses[j]
		if sum == 2020 {
			return expenses[i] * expenses[j], nil
		} else if sum < 2020 {
			i++
		} else {
			j--
		}
	}
	return 0, fmt.Errorf("No solution found")
}

// Part2 solve Day 1 Part 1
func Part2(expenses []int) (int, error) {
	sort.Ints(expenses)

	k := 0
	for k < len(expenses) {
		i := k + 1
		j := len(expenses) - 1

		for i < j {
			sum := expenses[i] + expenses[j] + expenses[k]
			if sum == 2020 {
				return expenses[i] * expenses[j] * expenses[k], nil
			} else if sum < 2020 {
				i++
			} else {
				j--
			}
		}
		k++
	}

	return 0, fmt.Errorf("No solution found")
}
