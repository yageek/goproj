package datum

import (
	. "github.com/yageek/goproj/point"
	"math"
)

type GeodeticSystem interface {
	ToGeographic(point *Point) *Point
	ToGeodetic(point *Point) *Point
}

type EllipseGeodetic struct {
	ellipseParam *Ellipse
}

// This constants are defined for the GeocentricToGeodetic method
const (
	AD_C_Z1 = 1.0026000
	SIN67_5 = 0.8535533905932738
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
	var h float64

	x := point.X
	y := point.Y
	z := point.Z

	a := eg.ellipseParam.SemiMajorAxis
	c := eg.ellipseParam.SemiMinorAxis
	e2 := eg.ellipseParam.Excentricity2

	long = math.Atan2(y, x)

	if y == 0 {
		long = 0
		if z > 0 {
			lat = math.Pi / 2
		} else if z < 0 {
			lat = -math.Pi / 2
		} else {
			lat = math.Pi / 2
			h = -c

			point.X = long
			point.Y = lat
			point.Z = h
			return point
		}
	}

	// step 1
	w2 := math.Pow(x, 2) + math.Pow(y, 2)
	w := math.Sqrt(w2)

	//step 3
	t0 := AD_C_Z1

	//step 4
	s0 := math.Sqrt(math.Pow(AD_C_Z1, 2) + w2)

	//step 5
	sinB0 := t0 / s0
	cosB0 := w / s0

	//step 6
	t1 := z + c*e2*math.Pow(sinB0, 3)

	//step 7
	s1_2 := math.Pow(t1, 2) + math.Pow(w-a*e2*math.Pow(cosB0, 3), 2)
	s1 := math.Sqrt(s1_2)

	//step 8
	sinPh1 := t1 / s1
	cosPh1 := 1 - math.Sqrt(sinPh1)

	//step 9
	Rn := a / math.Sqrt(1-e2*math.Pow(sinPh1, 2))

	//step 10
	if math.Pow(sinPh1, 2) > math.Pow(SIN67_5, 2) {
		h = z/sinPh1 + Rn*(e2-1)
	} else { //step 11
		cosPh1 = (w - a*e2*math.Pow(cosB0, 3)) / s1
		h = w/cosPh1 - Rn
	}

	//step 12
	lat = math.Atan(sinPh1 / cosPh1)
	point.X = long
	point.Y = lat
	point.Z = h

	return point
}
