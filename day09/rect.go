package main

type Rect struct {
	p1, p2 Point
	_area  int
}

func createRect(p1, p2 Point) Rect {
	return Rect{p1, p2, -1}
}

func (rect Rect) points() [4]Point {
	return [4]Point{
		rect.p1,
		createPoint(rect.p2.x, rect.p1.y),
		rect.p2,
		createPoint(rect.p1.x, rect.p2.y),
	}
}

func (rect Rect) lines() []Line {
	p := rect.points()
	return []Line{
		createLine(p[0], p[1]),
		createLine(p[1], p[2]),
		createLine(p[2], p[3]),
		createLine(p[3], p[0]),
	}
}

func (rect Rect) area() int {
	if rect._area < 0 {
		w := abs(rect.p1.x-rect.p2.x) + 1
		h := abs(rect.p1.y-rect.p2.y) + 1
		rect._area = w * h
	}
	return rect._area
}
