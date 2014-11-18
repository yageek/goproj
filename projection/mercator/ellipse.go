package mercator

import (
	. "github.com/yageek/goproj/datum"
	. "github.com/yageek/goproj/math"
	"math"
)

type EllipseProjection struct {
}

func (e *EllipseProjection) Forward(lambda, phi float64, datum *Datum) (x, y float64) {
	a := datum.Ellipsoid.SemiMajorAxis
	x = a * (lambda - datum.CentralMeridien)
	y = a * math.Log(math.Tan(PI_4+phi/2.0)*math.Pow((1-math.E*math.Sin(phi))/(1+math.E*math.Sin(phi)), math.E/2.0))
	return
}
