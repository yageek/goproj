package core

import (
	"math"
)

//The informations contained in an Ellipse
type EllipseParameters struct {
	SemiMajorAxis  float64
	SeminMinorAxis float64
	Flattening     float64
	Excentricity   float64
}

func NewEllipseParameters(a float64, b float64) *EllipseParameters {
	p := new(EllipseParameters)

	p.SemiMajorAxis = a
	p.SeminMinorAxis = b
	p.Flattening = (a - b) / a
	p.Excentricity = math.Sqrt(math.Pow(a, 2)-math.Pow(b, 2)) / a

	return p
}
