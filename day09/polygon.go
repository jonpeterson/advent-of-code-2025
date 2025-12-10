package main

import (
	"math"
)

type Polygon struct {
	faces []PolygonFace
}

type PolygonFace struct {
	line   Line
	normal byte
}

func createPolygon(points []Point, startNormal byte) Polygon {
	faces := make([]PolygonFace, 0, len(points))

	var prevPoint, point, nextPoint Point
	normal := startNormal

	for i := 0; i < len(points); i++ {
		point = points[i]
		if i < len(points)-1 {
			nextPoint = points[i+1]
		} else {
			nextPoint = points[0]
		}

		if i > 0 {
			if point.y == nextPoint.y {
				// line is horizontal
				if (nextPoint.x-prevPoint.x)*(nextPoint.y-prevPoint.y) >= 0 {
					// both positive or negative, rotate normal counter-clockwise
					normal--
				} else {
					// otherwise rotate normal clockwise
					normal++
				}
			} else {
				// line is vertical
				if (nextPoint.x-prevPoint.x)*(nextPoint.y-prevPoint.y) >= 0 {
					// both positive or negative, rotate normal clockwise
					normal++
				} else {
					// otherwise rotate normal counter-clockwise
					normal--
				}
			}

			if normal == 4 {
				normal = 0
			} else if normal == 255 {
				normal = 3
			}
		}

		face := PolygonFace{createLine(point, nextPoint), normal}
		faces = append(faces, face)

		prevPoint = point
	}

	return Polygon{faces}
}

func (polygon Polygon) containsRect(rect Rect) bool {
	for _, point := range rect.points() {
		if !polygon.containsPoint(point) {
			return false
		}
	}

	for _, line := range rect.lines() {
		for _, face := range polygon.faces {
			crossed, _ := line.crossesLine(face.line, true)
			if crossed {
				return false
			}
		}
	}

	return true
}

func (polygon Polygon) containsPoint(point Point) bool {
	for _, face := range polygon.faces {
		if face.line.containsPoint(point) {
			return true
		}
	}

	for _, dir := range dirs {
		foundFace, closestFace := polygon.closestFace(point, dir)
		if !foundFace || closestFace.normal != dir {
			return false
		}
	}

	return true
}

func (polygon Polygon) closestFace(point Point, dir byte) (bool, PolygonFace) {
	closestFaceDist := math.MaxInt
	var closestFace PolygonFace
	ray := point.ray(dir)

	for _, face := range polygon.faces {
		crossed, crossPoint := ray.crossesLine(face.line, false)
		if !crossed {
			continue
		}
		dist := abs((point.x - crossPoint.x) + (point.y - crossPoint.y))
		if dist < closestFaceDist {
			closestFaceDist = dist
			closestFace = face
		}
	}

	return closestFaceDist < math.MaxInt, closestFace
}
