package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type inst struct {
	op    string
	value int
}

func parseFile(path string) ([]*inst, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var instructions []*inst
	rInst := regexp.MustCompile("^(\\w+) ([+\\-]\\d+)")
	for _, line := range lines {
		match := rInst.FindStringSubmatch(line)
		v, err := strconv.ParseInt(match[2], 10, 16)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, &inst{op: match[1], value: int(v)})
	}
	return instructions, nil
}

func part2(instructions []*inst) int {
	for _, inst := range instructions {
		if inst.op == "nop" {
			inst.op = "jmp"
			acc, loop := part1(instructions)
			if !loop {
				return acc
			}
			inst.op = "nop"
		} else if inst.op == "jmp" {
			inst.op = "nop"
			acc, loop := part1(instructions)
			if !loop {
				return acc
			}
			inst.op = "jmp"
		}
	}
	return 0
}

func part1(instructions []*inst) (int, bool) {
	acc := 0
	pc := 0
	executed := make(map[int]int)
	for {
		// fmt.Printf("pc=%d, acc=%d\n", pc, acc)
		if pc >= len(instructions) {
			return acc, false
		}
		_, visited := executed[pc]
		if visited {
			return acc, true
		}
		executed[pc] = 1
		inst := instructions[pc]
		switch inst.op {
		case "nop":
			pc++
		case "acc":
			pc++
			acc += inst.value
		case "jmp":
			pc += inst.value
		}
	}
}

func main() {
	instructions, _ := parseFile("day08.txt")
	acc, _ := part1(instructions)
	fmt.Printf("Part 1: %d\n", acc)
	fmt.Printf("Part 2: %d\n", part2(instructions))
}
