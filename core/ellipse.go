package core

import (
	"math"
)

//The informations contained in an Ellipse
type EllipseParameters struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	Flattening    float64
	Excentricity  float64
	Excentricity2 float64
}

func NewEllipseParametersWithSemiAxis(a float64, b float64) *EllipseParameters {
	p := new(EllipseParameters)

	p.SemiMajorAxis = a
	p.SemiMinorAxis = b

	p.Flattening = (a - b) / a
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(b, 2)) / a
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}

func NewEllipseParametersWithFlattening(a float64, f float64) *EllipseParameters {
	p := new(EllipseParameters)

	p.SemiMajorAxis = a
	p.Flattening = f

	p.SemiMinorAxis = a * (1 + f)
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(p.SemiMinorAxis, 2)) / a
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}
