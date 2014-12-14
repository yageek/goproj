//The projection formula are taken from Snyder, p48-51
package mercator

import (
	. "github.com/yageek/goproj/datum"
	. "github.com/yageek/goproj/math"
	"math"
)

type SphereProjection struct {
	Radius float64 //The radius of the sphere
}

func (s *SphereProjection) Forward(lambda, phi float64, datum *Datum) (x, y float64) {
	x = s.Radius * (lambda - datum.CentralMeridien)
	y = s.Radius * math.Log(math.Tan(PI_4+phi/2.0))
	return
}

func (s *SphereProjection) Inverse(x, y float64, datum *Datum) (lambda, phi float64) {
	phi = PI_2 - 2*math.Atan(math.Exp(-y/s.Radius))
	lambda = x/s.Radius + datum.CentralMeridien
	return
}

func (s *SphereProjection) scaleFactor(lamdba, phi float64) float64 {
	return 1.0 / math.Cos(phi)
}

type EllipseProjection struct {
}

func (e *EllipseProjection) Forward(lambda, phi float64, datum *Datum) (x, y float64) {
	a := datum.Ellipsoid.SemiMajorAxis
	x = a * (lambda - datum.CentralMeridien)
	y = a * math.Log(math.Tan(PI_4+phi/2.0)*math.Pow((1-math.E*math.Sin(phi))/(1+math.E*math.Sin(phi)), math.E/2.0))
	return
}
