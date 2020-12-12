package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

type hall struct {
	Seats   []*seat
	SeatMap map[position]*seat
	Size    int
}

type seat struct {
	Pos  position
	Free bool
}

type position struct {
	X int
	Y int
}

func (hall *hall) String() string {
	b := &strings.Builder{}
	for y := 0; y < hall.Size; y++ {
		for x := 0; x < hall.Size; x++ {
			s, ok := hall.find(x, y)
			if ok {
				if s.Free {
					b.WriteString("L")
				} else {
					b.WriteString("#")
				}
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (hall *hall) hash() *big.Int {
	h := big.NewInt(0)
	for i, seat := range hall.Seats {
		b := uint(1)
		if seat.Free {
			b = uint(0)
		}
		h = h.SetBit(h, i, b)
	}
	return h
}

func (hall *hall) find(x int, y int) (*seat, bool) {
	if x < 0 || x >= hall.Size || y < 0 || y >= hall.Size {
		return nil, false
	}
	seat, ok := hall.SeatMap[position{X: x, Y: y}]
	return seat, ok
}

func (hall *hall) occupancy(seat *seat, maxRange int) int {
	occupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 1; k <= maxRange; k++ {
				s, ok := hall.find(seat.Pos.X+(i*k), seat.Pos.Y+(j*k))
				if ok {
					if !s.Free {
						occupied++
					}
					break
				}
			}
		}
	}
	return occupied
}

func (hall *hall) count() int {
	occupied := 0
	for _, seat := range hall.Seats {
		if !seat.Free {
			occupied++
		}
	}
	return occupied
}

func parseFile(path string) (*hall, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	hall := &hall{}
	hall.SeatMap = make(map[position]*seat)
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'L':
				pos := position{X: x, Y: y}
				seat := &seat{
					Pos:  pos,
					Free: true,
				}
				hall.Seats = append(hall.Seats, seat)
				hall.SeatMap[pos] = seat
			}
		}
		hall.Size = y + 1
	}
	return hall, nil
}

func round(hall *hall, maxOccupancy int, maxRange int) {
	var flip []position
	for _, seat := range hall.Seats {
		occupied := hall.occupancy(seat, maxRange)
		if (seat.Free && occupied == 0) || (!seat.Free && occupied >= maxOccupancy) {
			flip = append(flip, seat.Pos)
		}
	}
	for _, pos := range flip {
		seat, _ := hall.find(pos.X, pos.Y)
		seat.Free = !seat.Free
	}
}

func solve(hall *hall, round func(*hall)) int {
	hashes := make(map[string]int)
	for {
		round(hall)
		h := hall.hash()
		_, ok := hashes[h.String()]
		if ok {
			return hall.count()
		}
		hashes[h.String()] = 0
	}
}

func main() {
	h, _ := parseFile("day11.txt")
	fmt.Printf("Part 1: %d\n", solve(h, func(ha *hall) {
		round(ha, 4, 1)
	}))
	h, _ = parseFile("day11.txt")
	fmt.Printf("Part 2: %d\n", solve(h, func(ha *hall) {
		round(ha, 5, ha.Size-1)
	}))
}
