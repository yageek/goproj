package mercator

import (
	"math"
)

type SphereProjection struct{}

func (s *SphereProjection) Forward(lambda, phi float64, datum *Datum) (x, y float64) {

}
