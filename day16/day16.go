package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	Name string
	Min1 int
	Max1 int
	Min2 int
	Max2 int
}

func (rule rule) IsValid(value int) bool {
	return (value >= rule.Min1 && value <= rule.Max1) || (value >= rule.Min2 && value <= rule.Max2)
}

type ticket []int

func parseTicket(data string) (ticket, error) {
	numbers := strings.Split(data, ",")
	var t ticket
	for _, nb := range numbers {
		v, err := strconv.Atoi(nb)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		t = append(t, v)
	}
	return t, nil
}

func parseFile(rulesPath, myTicketPath, nearbyTicketsPath string) ([]rule, ticket, []ticket, error) {
	data, err := ioutil.ReadFile(rulesPath)
	if err != nil {
		return nil, nil, nil, err
	}
	lines := strings.Split(string(data), "\n")
	r := regexp.MustCompile(`^([a-zA-Z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	var rules []rule
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		min1, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, nil, nil, err
		}
		max1, err := strconv.Atoi(match[3])
		if err != nil {
			return nil, nil, nil, err
		}
		min2, err := strconv.Atoi(match[4])
		if err != nil {
			return nil, nil, nil, err
		}
		max2, err := strconv.Atoi(match[5])
		if err != nil {
			return nil, nil, nil, err
		}
		rule := rule{Name: match[1], Min1: min1, Max1: max1, Min2: min2, Max2: max2}
		rules = append(rules, rule)
	}

	data, err = ioutil.ReadFile(myTicketPath)
	if err != nil {
		return nil, nil, nil, err
	}
	my, err := parseTicket(string(data))
	if err != nil {
		return nil, nil, nil, err
	}

	data, err = ioutil.ReadFile(nearbyTicketsPath)
	if err != nil {
		return nil, nil, nil, err
	}
	lines = strings.Split(string(data), "\n")
	var tickets []ticket
	for _, line := range lines {
		t, err := parseTicket(line)
		if err != nil {
			return nil, nil, nil, err
		}
		tickets = append(tickets, t)
	}

	return rules, my, tickets, nil
}

func allRules(vs []rule, f func(rule) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func allTickets(vs []ticket, f func(ticket) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filterRules(vs []rule, f func(rule) bool) []rule {
	vsf := make([]rule, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func part1(rules []rule, nearby []ticket) int {
	sum := 0
	for _, ticket := range nearby {
		for _, value := range ticket {
			if allRules(rules, func(r rule) bool { return !r.IsValid(value) }) {
				sum += value
			}
		}
	}
	return sum
}

func part2(rules []rule, my ticket, nearby []ticket) int {
	valid := []ticket{my}
	for _, ticket := range nearby {
		invalid := false
		for _, value := range ticket {
			invalid = invalid || allRules(rules, func(r rule) bool { return !r.IsValid(value) })
		}
		if !invalid {
			valid = append(valid, ticket)
		}
	}

	flen := len(my)
	fields := make([][]rule, flen)
	for field := 0; field < flen; field++ {
		fields[field] = make([]rule, 0)
		for _, rule := range rules {
			if allTickets(valid, func(t ticket) bool { return rule.IsValid(t[field]) }) {
				fields[field] = append(fields[field], rule)
			}
		}
	}

	match := make([]rule, flen)
	for !allRules(match, func(rule rule) bool { return rule.Name != "" }) {
		for _, r := range match {
			for i, f := range fields {
				fields[i] = filterRules(f, func(rule rule) bool { return r != rule })
			}
		}
		for i, f := range fields {
			if len(f) == 1 {
				match[i] = f[0]
			}
		}
	}

	prod := 1
	for i, r := range match {
		if strings.HasPrefix(r.Name, "departure") {
			prod *= my[i]
		}
	}

	return prod
}

func main() {
	rules, my, nearby, err := parseFile("rules.txt", "my.txt", "nearby.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1: %d\n", part1(rules, nearby))
	fmt.Printf("Part2: %d\n", part2(rules, my, nearby))
}
