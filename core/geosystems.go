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

	x := point.X
	y := point.Y
	z := point.Z

	a := eg.ellipseParam.SemiMajorAxis
	c := eg.ellipseParam.SemiMinorAxis
	e2 := eg.ellipseParam.Excentricity2

	long = math.Atan2(y, x)

	// step 1
	w2 := math.Pow(x, 2) + math.Pow(y, 2)
	w := math.Sqrt(w2)

	//step 3
	t0 := AD_C_Z1

	//step 4
	s0 := math.Sqrt(math.Pow(AD_C_Z1, 2) + w2)

	//step 5
	sinB0 := t0 / s0
	cosB0 := W / s0

	//step 6
	t1 := z + c*e2*math.Pow(sinB0, 3)

	//step 7
	//s1_2 := math.Pow(t1, 2) + math.Pow(x, y)
	point.x = long
	point.y = lat

	return point
}
