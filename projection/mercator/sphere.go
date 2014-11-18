package mercator

import (
	. "github.com/yageek/goproj/datum"
	. "github.com/yageek/goproj/math"
	"math"
)

// This value is taken from openstreetmap : http://wiki.openstreetmap.org/wiki/Mercator#C
type SphereProjection struct {
	Radius float64 //The radius of the sphere
}

//The projection formula are taken from Snyder, p48-51
func (s *SphereProjection) Forward(lambda, phi float64, datum *Datum) (x, y float64) {
	x = s.Radius * (lambda - datum.CentralMeridien)
	y = s.Radius * math.Log(math.Tan(PI_4+phi/2.0))
	return
}

//The projection formula are taken from Snyder, p48-51
func (s *SphereProjection) Inverse(x, y float64, datum *Datum) (lambda, phi float64) {
	phi = PI_2 - 2*math.Atan(math.Exp(-y/s.Radius))
	lambda = x/s.Radius + datum.CentralMeridien
	return
}
