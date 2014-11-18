package datum

import (
	"math"
)

//The informations contained in an Ellipse
type Ellipse struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	Flattening    float64
	Excentricity  float64
	Excentricity2 float64
}

func NewEllipseWithSemiAxis(a float64, b float64) *Ellipse {
	p := new(Ellipse)

	p.SemiMajorAxis = a
	p.SemiMinorAxis = b

	p.Flattening = (a - b) / a
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(b, 2)) / math.Pow(a, 2)
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}

func NewEllipseWithFlattening(a float64, f float64) *Ellipse {
	p := new(Ellipse)

	p.SemiMajorAxis = a
	p.Flattening = f

	p.SemiMinorAxis = a * (1 - f)
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(p.SemiMinorAxis, 2)) / math.Pow(a, 2)
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}
