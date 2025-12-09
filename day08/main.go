package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z int
	circuit *Circuit
}

type BoxPair struct {
	a, b  *Box
	dist2 float64
}

type Circuit struct {
	boxes []*Box
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	// Gather boxes
	boxes := make([]*Box, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		box := boxFromCoords(coords)
		boxes = append(boxes, box)
		lines = append(lines, line)
	}

	// Gather all pairs
	boxCount := len(boxes)
	pairCount := (boxCount * (boxCount - 1)) / 2
	pairs := make([]*BoxPair, 0, pairCount)
	pairs = appendAllPairs(boxes, pairs)

	// Shortest 1000 pairs
	slices.SortFunc(pairs, func(a *BoxPair, b *BoxPair) int {
		if a.dist2 < b.dist2 {
			return -1
		} else {
			return 1
		}
	})

	fmt.Printf("Part 1: %d\n", part1(pairs))
	fmt.Printf("Part 2: %d\n", part2(pairs))
}

func part1(pairs []*BoxPair) int {
	// Create circuits
	circuits := make([]*Circuit, 0)
	for _, p := range pairs[:1000] {
		ac := p.a.circuit
		bc := p.b.circuit
		if ac == nil {
			if bc == nil {
				circuit := &Circuit{boxes: []*Box{p.a, p.b}}
				circuits = append(circuits, circuit)
				p.a.circuit = circuit
				p.b.circuit = circuit
			} else {
				p.a.circuit = bc
				bc.boxes = append(bc.boxes, p.a)
			}
		} else if bc == nil {
			p.b.circuit = ac
			ac.boxes = append(ac.boxes, p.b)
		} else if bc != ac {
			for _, bcb := range bc.boxes {
				bcb.circuit = ac
				ac.boxes = append(ac.boxes, bcb)
			}
			bc.boxes = []*Box{}
		}
	}

	// Longest 3 circuits
	slices.SortFunc(circuits, func(a *Circuit, b *Circuit) int {
		if len(a.boxes) > len(b.boxes) {
			return -1
		} else {
			return 1
		}
	})
	circuits = circuits[:3]

	total := 1
	for _, c := range circuits {
		lb := len(c.boxes)
		if lb > 0 {
			total *= lb
		}
	}

	return total
}

func part2(pairs []*BoxPair) int {
	// Create circuits
	circuits := make([]*Circuit, 0)
	for _, p := range pairs {
		ac := p.a.circuit
		bc := p.b.circuit
		if ac == nil {
			if bc == nil {
				circuit := &Circuit{boxes: []*Box{p.a, p.b}}
				circuits = append(circuits, circuit)
				p.a.circuit = circuit
				p.b.circuit = circuit
			} else {
				p.a.circuit = bc
				bc.boxes = append(bc.boxes, p.a)
			}
		} else if bc == nil {
			p.b.circuit = ac
			ac.boxes = append(ac.boxes, p.b)
		} else if bc != ac {
			for _, bcb := range bc.boxes {
				bcb.circuit = ac
				ac.boxes = append(ac.boxes, bcb)
			}
			bc.boxes = []*Box{}

			if len(ac.boxes) == 1000 {
				return p.a.x * p.b.x
			}
		}
	}

	return 0
}

func boxFromCoords(coords []string) *Box {
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return &Box{x, y, z, nil}
}

func appendAllPairs(boxes []*Box, pairs []*BoxPair) []*BoxPair {
	for i := range boxes {
		pairs = boxes[i].appendPairs(boxes, i+1, pairs)
	}
	return pairs
}

func (box *Box) appendPairs(boxes []*Box, start int, pairs []*BoxPair) []*BoxPair {
	for i := start; i < len(boxes); i++ {
		b := boxes[i]
		pair := BoxPair{box, b, box.distance2(b)}
		pairs = append(pairs, &pair)
	}
	return pairs
}

func (box *Box) distance2(b *Box) float64 {
	x := box.x - b.x
	y := box.y - b.y
	z := box.z - b.z
	return float64(x*x + y*y + z*z)
}
