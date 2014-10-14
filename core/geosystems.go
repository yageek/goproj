package core

import (
	"math"
)

type GeodeticSystem interface {
	ToGeographic(point *Point) *Point
	ToGeodetic(point *Point) *Point
}

type EllipseGeodetic struct {
	ellipseParam *EllipseParameters
}

func (eg *EllipseGeodetic) ToGeographic(point *Point) *Point {

	var lat float64
	var long float64

	if x != 0 {
		long = math.Atan2()
	} else {

	}

	point.x = long
	point.y = lat

	return point
}
