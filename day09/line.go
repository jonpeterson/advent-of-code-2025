package main

type Line interface {
	containsPoint(point Point) bool
	crossesLine(other Line, excludePoints bool) (bool, Point)
}

type HLine struct {
	x1, x2, y int
}

type VLine struct {
	x, y1, y2 int
}

func createLine(p1, p2 Point) Line {
	if p1.y == p2.y {
		x1 := p1.x
		x2 := p2.x
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		return HLine{x1, x2, p1.y}
	} else {
		y1 := p1.y
		y2 := p2.y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		return VLine{p1.x, y1, y2}
	}
}

func (line HLine) containsPoint(point Point) bool {
	return line.x1 <= point.x && line.x2 >= point.x && line.y == point.y
}

func (line VLine) containsPoint(point Point) bool {
	return line.x == point.x && line.y1 <= point.y && line.y2 >= point.y
}

func (line HLine) crossesLine(other Line, excludePoints bool) (bool, Point) {
	v, ok := other.(VLine)
	if !ok {
		return false, Point{}
	} else if excludePoints {
		return line.x1 < v.x && line.x2 > v.x && line.y > v.y1 && line.y < v.y2, Point{v.x, line.y}
	} else {
		return line.x1 <= v.x && line.x2 >= v.x && line.y >= v.y1 && line.y <= v.y2, Point{v.x, line.y}
	}
}

func (line VLine) crossesLine(other Line, excludePoints bool) (bool, Point) {
	h, ok := other.(HLine)
	if !ok {
		return false, Point{}
	} else if excludePoints {
		return line.x > h.x1 && line.x < h.x2 && line.y1 < h.y && line.y2 > h.y, Point{line.x, h.y}
	} else {
		return line.x >= h.x1 && line.x <= h.x2 && line.y1 <= h.y && line.y2 >= h.y, Point{line.x, h.y}
	}
}
