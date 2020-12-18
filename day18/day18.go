package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Operator int

const (
	Plus = iota
	Multiply
	LeftParenthesis
	RightParenthesis
)

var precedenceP1 = map[Operator]int{
	Plus:             1,
	Multiply:         1,
	LeftParenthesis:  99,
	RightParenthesis: 99,
}

var precedenceP2 = map[Operator]int{
	Plus:             1,
	Multiply:         10,
	LeftParenthesis:  99,
	RightParenthesis: 99,
}

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *stack) Peek() int {
	r := (*s)[len(*s)-1]
	return r
}

func (s *stack) Size() int {
	return len(*s)
}

func apply(op Operator, v1, v2 int) int {
	switch op {
	case Plus:
		return v1 + v2
	case Multiply:
		return v1 * v2
	default:
		panic("Unexpected operator")
	}
}

func eval(expr string, precedence map[Operator]int) int {
	expr = strings.ReplaceAll(expr, " ", "")
	var values stack
	var ops stack
	for _, c := range expr {
		if c == '(' {
			ops.Push(LeftParenthesis)
		} else if c == ')' {
			for {
				op := ops.Pop()
				if op == LeftParenthesis {
					break
				}
				v1 := values.Pop()
				v2 := values.Pop()
				v := apply(Operator(op), v1, v2)
				values.Push(v)
			}
		} else if c == '+' {
			for {
				if ops.Size() == 0 || precedence[Operator(ops.Peek())] > precedence[Plus] {
					ops.Push(Plus)
					break
				} else {
					op := ops.Pop()
					v1 := values.Pop()
					v2 := values.Pop()
					v := apply(Operator(op), v1, v2)
					values.Push(v)
				}
			}
		} else if c == '*' {
			for {
				if ops.Size() == 0 || precedence[Operator(ops.Peek())] > precedence[Multiply] {
					ops.Push(Multiply)
					break
				} else {
					op := ops.Pop()
					v1 := values.Pop()
					v2 := values.Pop()
					v := apply(Operator(op), v1, v2)
					values.Push(v)
				}
			}
		} else if c >= '0' && c <= '9' {
			v, _ := strconv.Atoi(string(c))
			values = append(values, v)
		}
	}
	for values.Size() > 1 {
		op := ops.Pop()
		v1 := values.Pop()
		v2 := values.Pop()
		v := apply(Operator(op), v1, v2)
		values.Push(v)
	}
	return values[0]
}

func parseFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func main() {
	expressions, _ := parseFile("day18.txt")
	sum := 0
	for _, expr := range expressions {
		v := eval(expr, precedenceP1)
		sum += v
	}
	fmt.Printf("Part1: %d\n", sum)

	sum = 0
	for _, expr := range expressions {
		v := eval(expr, precedenceP2)
		sum += v
	}

	fmt.Printf("Part2: %d\n", sum)
}
