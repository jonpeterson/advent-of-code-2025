package main

import (
	"math"
)

type Point struct {
	x, y int
}

func createPoint(x, y int) Point {
	return Point{x, y}
}

func (point Point) ray(dir byte) Line {
	switch dir {
	case dirNorth:
		return createLine(point, Point{point.x, math.MinInt})
	case dirEast:
		return createLine(point, Point{math.MaxInt, point.y})
	case dirSouth:
		return createLine(point, Point{point.x, math.MaxInt})
	default:
		return createLine(point, Point{math.MinInt, point.y})
	}
}
