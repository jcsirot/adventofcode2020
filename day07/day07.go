package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	Color string
	Rules []*rule
}

type rule struct {
	Count int
	Bag   *bag
}

func find(bags []*bag, color string) *bag {
	for _, bag := range bags {
		if bag.Color == color {
			return bag
		}
	}
	return nil
}

func (bag *bag) String() string {
	return fmt.Sprintf("%s (%s)", bag.Color, bag.Rules)
}

func (rule *rule) String() string {
	return fmt.Sprintf("%d %s", rule.Count, rule.Bag)
}

func parseRules(path string) ([]*bag, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var bags []*bag
	rColor := regexp.MustCompile("^(\\w+ \\w+)")
	rContent := regexp.MustCompile("(\\d+) (\\w+ \\w+)")
	for _, line := range lines {
		color := rColor.FindString(line)
		bag := &bag{
			Color: color,
		}
		bags = append(bags, bag)
	}
	for _, line := range lines {
		color := rColor.FindString(line)
		bag := find(bags, color)
		content := rContent.FindAllStringSubmatch(line, -1)
		for _, r := range content {
			count, _ := strconv.Atoi(r[1])
			b := find(bags, r[2])
			bag.Rules = append(bag.Rules, &rule{
				Count: count,
				Bag:   b,
			})
		}
	}
	return bags, nil
}

func contains(bags []*bag, bag *bag) bool {
	for _, b := range bags {
		if b == bag {
			return true
		}
	}
	return false
}

func walkPart1(bag *bag, found []*bag) bool {
	if bag.Color == "shiny gold" {
		return true
	}
	for _, rule := range bag.Rules {
		if contains(found, bag) || walkPart1(rule.Bag, found) {
			return true
		}
	}
	return false
}

func part1(bags []*bag) int {
	var found []*bag
	for _, bag := range bags {
		if walkPart1(bag, found) {
			found = append(found, bag)
		}
	}
	return len(found) - 1
}

func walkPart2(bag *bag, subSums map[string]int) int {
	sum := 0
	for _, rule := range bag.Rules {
		if subSums[rule.Bag.Color] == 0 {
			subSums[rule.Bag.Color] = walkPart2(rule.Bag, subSums)
		}
		sum += (1 + subSums[rule.Bag.Color]) * rule.Count
	}
	return sum
}

func part2(bags []*bag) int {
	subSums := make(map[string]int)
	myBag := find(bags, "shiny gold")
	return walkPart2(myBag, subSums)
}

func main() {
	bags, _ := parseRules("day07.txt")
	fmt.Printf("Part 1: %d\n", part1(bags))
	fmt.Printf("Part 2: %d\n", part2(bags))
}
