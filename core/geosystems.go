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

// This constants are defined for the GeocentricToGeodetic method
const (
	AD_C_Z1 = 1.0026000
	AD_C_Z2 = 1.00092592
	AD_C_Z3 = 0.999250297
	AD_C_Z4 = 0.997523508
)

// GeocentricToGeodetic converts the point's coordintates
// from the cartesian system (geocentric) to a geodetic system defined
// by the ellipse.
//
// This is based on the non iterative method described within paper :
//  'An Improved Algorithm for Geocentric to Geodetic Coordinate Conversion'
//  Ralph Toms, Feb 1996
//
func (eg *EllipseGeodetic) GeocentricToGeodetic(point *Point) *Point {

	var lat float64
	var long float64

	point.x

	long = math.Atan2(point.y, point.x)

	w2 := math.Pow(point.x, 2) + math.Pow(point.y, 2)

	point.x = long
	point.y = lat

	return point
}
