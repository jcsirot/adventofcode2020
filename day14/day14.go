package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type maskInstr struct {
	Value string
}

type setInstr struct {
	Address uint64
	Value   uint64
}

func readFile(path string) ([]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var instructions []interface{}
	r1 := regexp.MustCompile("^mask = ([X01]+)$")
	r2 := regexp.MustCompile("^mem\\[(\\d+)] = (\\d+)$")
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			matches := r1.FindStringSubmatch(line)
			instructions = append(instructions, maskInstr{Value: matches[1]})
		} else if strings.HasPrefix(line, "mem") {
			matches := r2.FindStringSubmatch(line)
			a, err := strconv.ParseUint(matches[1], 10, 36)
			v, err := strconv.ParseUint(matches[2], 10, 36)
			if err != nil {
				return nil, err
			}
			instructions = append(instructions, setInstr{Address: a, Value: v})
		}
	}

	return instructions, nil
}

func applyValueMask(mask string, value uint64) uint64 {
	bin := fmt.Sprintf("%036b", value)
	var r strings.Builder
	for i, m := range mask {
		if m == rune('X') {
			fmt.Fprintf(&r, "%s", string(bin[i]))
		} else if m == rune('0') {
			fmt.Fprintf(&r, "%s", "0")
		} else if m == rune('1') {
			fmt.Fprintf(&r, "%s", "1")
		}
	}
	u, _ := strconv.ParseUint(r.String(), 2, 36)
	return u
}

func applyAddressMask(mask string, address uint64) []uint64 {
	bin := fmt.Sprintf("%036b", address)
	var addresses []uint64
	index := strings.Index(mask, "X")
	if index != -1 {
		m := strings.Replace(mask, "X", "0", 1)
		a0 := address &^ (1 << (35 - index))
		addresses = append(addresses, applyAddressMask(m, a0)...)
		a1 := address | (1 << (35 - index))
		addresses = append(addresses, applyAddressMask(m, a1)...)
	} else {
		var r strings.Builder
		for i, c := range mask {
			if c == rune('0') {
				fmt.Fprintf(&r, "%s", string(bin[i]))
			} else if c == rune('1') {
				fmt.Fprintf(&r, "%s", "1")
			}
		}
		u, _ := strconv.ParseUint(r.String(), 2, 36)
		addresses = append(addresses, u)
	}
	return addresses
}

func part1(instructions []interface{}) uint64 {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	mem := make(map[uint64]uint64)
	for _, instr := range instructions {
		switch instr.(type) {
		case maskInstr:
			mask = instr.(maskInstr).Value
		case setInstr:
			v := applyValueMask(mask, instr.(setInstr).Value)
			mem[instr.(setInstr).Address] = v
		}
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func part2(instructions []interface{}) uint64 {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	mem := make(map[uint64]uint64)
	for _, instr := range instructions {
		switch instr.(type) {
		case maskInstr:
			mask = instr.(maskInstr).Value
		case setInstr:
			//v := applyValueMask(mask, instr.(setInstr).Value)
			for _, a := range applyAddressMask(mask, instr.(setInstr).Address) {
				mem[a] = instr.(setInstr).Value
			}
		}
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func main() {
	instructions, _ := readFile("day14.txt")
	fmt.Printf("Part 1: %d\n", part1(instructions))
	fmt.Printf("Part 2: %d\n", part2(instructions))
}
