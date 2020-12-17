package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Space struct {
	Cubes      map[Point]bool
	MinX, MaxX int
	MinY, MaxY int
	MinZ, MaxZ int
}

func (space *Space) String() string {
	var b bytes.Buffer
	for z := space.MinZ; z <= space.MaxZ; z++ {
		fmt.Fprintf(&b, "z=%d\n", z)
		for y := space.MinY; y <= space.MaxY; y++ {
			for x := space.MinX; x <= space.MaxX; x++ {
				p := Point{X: x, Y: y, Z: z}
				active, ok := space.Cubes[p]
				if ok && active {
					fmt.Fprintf(&b, "#")
				} else {
					fmt.Fprintf(&b, ".")
				}
			}
			fmt.Fprintln(&b)
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (space *Space) ToList() []Point {
	list := []Point{}
	for k := range space.Cubes {
		list = append(list, k)
	}
	return list
}

func (space *Space) ActiveNeighbours(p Point) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				active, found := space.Cubes[Point{X: p.X + x, Y: p.Y + y, Z: p.Z + z}]
				if found && active {
					count++
				}
			}
		}
	}
	return count
}

func (space *Space) Step() {
	actives := []Point{}
	inactives := []Point{}
	for z := space.MinZ - 1; z <= space.MaxZ+1; z++ {
		for x := space.MinX - 1; x <= space.MaxX+1; x++ {
			for y := space.MinY - 1; y <= space.MaxY+1; y++ {
				p := Point{X: x, Y: y, Z: z}
				//fmt.Println(p)
				active, ok := space.Cubes[p]
				count := space.ActiveNeighbours(p)
				//fmt.Println(count)
				if (!ok || !active) && count == 3 {
					actives = append(actives, p)
				} else if active && (count < 2 || count > 3) {
					inactives = append(inactives, p)
				}
			}
		}
	}
	for _, pt := range actives {
		space.Cubes[pt] = true
		if pt.X < space.MinX {
			space.MinX = pt.X
		}
		if pt.Y < space.MinY {
			space.MinY = pt.Y
		}
		if pt.Z < space.MinZ {
			space.MinZ = pt.Z
		}
		if pt.X > space.MaxX {
			space.MaxX = pt.X
		}
		if pt.Y > space.MaxY {
			space.MaxY = pt.Y
		}
		if pt.Z > space.MaxZ {
			space.MaxZ = pt.Z
		}
	}
	for _, pt := range inactives {
		delete(space.Cubes, pt)
	}
}

func parseFile(path string) (*Space, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	size := len(lines[0])
	space := &Space{
		Cubes: make(map[Point]bool),
		MinZ:  0,
		MaxZ:  0,
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if lines[y][x] == '#' {
				space.Cubes[Point{X: x, Y: y, Z: 0}] = true
				if space.MinX > x {
					space.MinX = x
				}
				if space.MinY > y {
					space.MinY = y
				}
				if space.MaxX < x {
					space.MaxX = x
				}
				if space.MaxY < y {
					space.MaxY = y
				}
			}
		}
	}
	return space, nil
}

func main() {
	space, _ := parseFile("day17.txt")
	space.Step()
	space.Step()
	space.Step()
	space.Step()
	space.Step()
	space.Step()

	count := 0
	for _, active := range space.Cubes {
		if active {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}
