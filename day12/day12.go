package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/adam-lavrik/go-imath/ix"
)

type move struct {
	Action string
	Value  int
}

type ship struct {
	Heading int
	X       int
	Y       int
}

type waypoint struct {
	Rx int
	Ry int
}

func (ship *ship) move(heading int, value int) {
	switch heading {
	case 0:
		ship.Y -= value
	case 90:
		ship.X -= value
	case 180:
		ship.Y += value
	case 270:
		ship.X += value
	}
}

func (ship *ship) turn(value int) {
	ship.Heading = (ship.Heading + value) % 360
}

func (ship *ship) moveForwardWaypoint(wp waypoint, value int) {
	ship.X += wp.Rx * value
	ship.Y += wp.Ry * value
}

func (wp *waypoint) move(dx int, dy int) {
	wp.Rx += dx
	wp.Ry += dy
}

func (wp *waypoint) turn(value int) {
	rx := 0
	ry := 0
	switch value {
	case 90:
		rx = -wp.Ry
		ry = wp.Rx
	case 180:
		rx = -wp.Rx
		ry = -wp.Ry
	case 270:
		rx = wp.Ry
		ry = -wp.Rx
	}
	wp.Rx = rx
	wp.Ry = ry
}

func parseFile(path string) ([]move, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var moves []move
	for _, line := range lines {
		v, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		moves = append(moves, move{Action: line[0:1], Value: v})
	}
	return moves, nil
}

func part1(moves []move) int {
	ship := &ship{Heading: 90}
	for _, move := range moves {
		switch move.Action {
		case "F":
			ship.move(ship.Heading, move.Value)
		case "N":
			ship.move(0, move.Value)
		case "E":
			ship.move(90, move.Value)
		case "S":
			ship.move(180, move.Value)
		case "W":
			ship.move(270, move.Value)
		case "R":
			ship.turn(move.Value)
		case "L":
			ship.turn(360 - move.Value)
		}
	}
	return ix.Abs(ship.X) + ix.Abs(ship.Y)
}

func part2(moves []move) int {
	ship := &ship{}
	wp := &waypoint{Rx: 10, Ry: -1}
	for _, move := range moves {
		switch move.Action {
		case "F":
			ship.moveForwardWaypoint(*wp, move.Value)
		case "N":
			wp.move(0, -move.Value)
		case "E":
			wp.move(move.Value, 0)
		case "S":
			wp.move(0, move.Value)
		case "W":
			wp.move(-move.Value, 0)
		case "R":
			wp.turn(move.Value)
		case "L":
			wp.turn(360 - move.Value)
		}
	}
	return ix.Abs(ship.X) + ix.Abs(ship.Y)
}

func main() {
	moves, _ := parseFile("day12.txt")
	fmt.Printf("Part 1: %d\n", part1(moves))
	fmt.Printf("Part 2: %d\n", part2(moves))
}
