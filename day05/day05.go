package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readSeatID(path string) ([]uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ids := make([]uint64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()
		code = strings.ReplaceAll(code, "B", "1")
		code = strings.ReplaceAll(code, "F", "0")
		code = strings.ReplaceAll(code, "R", "1")
		code = strings.ReplaceAll(code, "L", "0")
		id, err := strconv.ParseUint(code, 2, 16)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func part1(ids []uint64) uint64 {
	max := uint64(0)
	for _, id := range ids {
		if id > max {
			max = id
		}
	}
	return max
}

func part2(ids []uint64) uint64 {
	sort.Slice(ids, func(i, j int) bool { return ids[i] <= ids[j] })
	for i := range ids {
		if i >= 1 && ids[i-1]+2 == ids[i] {
			return ids[i-1] + 1
		}
	}
	return 0
}

func main() {
	ids, _ := readSeatID("day05.txt")
	fmt.Printf("Part 1: %d\n", part1(ids))
	fmt.Printf("Part 2: %d\n", part2(ids))
}
