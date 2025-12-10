package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Depends on input.
// TODO: figure out how to detect
var startingPolygonNormal = dirEast

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := make([]Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		points = append(points, createPoint(x, y))
	}

	rects := make([]Rect, 0)
	for i1, p1 := range points {
		for i2 := i1 + 1; i2 < len(points); i2++ {
			rects = append(rects, createRect(p1, points[i2]))
		}
	}

	slices.SortFunc(rects, func(a, b Rect) int {
		if a.area() > b.area() {
			return -1
		} else if a.area() < b.area() {
			return 1
		} else {
			return 0
		}
	})

	fmt.Printf("Part 1: %d\n", rects[0].area())

	polygon := createPolygon(points, startingPolygonNormal)

	var largestNonCrossing Rect
	for _, rect := range rects {
		if polygon.containsRect(rect) {
			largestNonCrossing = rect
			break
		}
	}

	fmt.Printf("Part 2: %d\n", largestNonCrossing.area())
}
