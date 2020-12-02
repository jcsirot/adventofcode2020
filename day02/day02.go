package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Policy is a password policy
type Policy struct {
	Value1 int
	Value2 int
	Char   rune
}

func (p Policy) checkPart1(password string) bool {
	count := 0
	for _, c := range password {
		if c == p.Char {
			count++
		}
	}
	return count <= p.Value2 && count >= p.Value1
}

func (p Policy) checkPart2(password string) bool {
	return rune(password[p.Value1-1]) == p.Char && rune(password[p.Value2-1]) != p.Char ||
		rune(password[p.Value1-1]) != p.Char && rune(password[p.Value2-1]) == p.Char
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseLine(line string) (Policy, string) {
	r := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)")
	s := r.FindStringSubmatch(line)
	p := Policy{
		Value1: atoi(s[1]),
		Value2: atoi(s[2]),
		Char:   []rune(s[3])[0],
	}
	return p, s[4]
}

func part1(passwords []string) int {
	valid := 0
	for _, line := range passwords {
		p, password := parseLine(line)
		if p.checkPart1(password) {
			valid++
		}
	}
	return valid
}

func part2(passwords []string) int {
	valid := 0
	for _, line := range passwords {
		p, password := parseLine(line)
		if p.checkPart2(password) {
			valid++
		}
	}
	return valid
}

func main() {
	passwords, err := readLines("day02.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(passwords))
	fmt.Printf("Part 2: %d\n", part2(passwords))
}
