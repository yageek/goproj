package ellipsoid

import (
	"math"
)

//The informations contained in an Ellipsoid
type Ellipsoid struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	Flattening    float64
	Excentricity  float64
	Excentricity2 float64
}

func NewEllipsoidWithSemiAxis(a float64, b float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.SemiMinorAxis = b

	p.Flattening = (a - b) / a
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(b, 2)) / math.Pow(a, 2)
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}

func NewEllipsoidWithFlattening(a float64, f float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.Flattening = f

	p.SemiMinorAxis = a * (1 - f)
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(p.SemiMinorAxis, 2)) / math.Pow(a, 2)
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}

type EpsgEllipsoid struct {
	*Ellipsoid
	Name string
	Epsg int
}

func NewEpsgEllipsoidWithSemiAxis(name string, epsg int, a float64, b float64) *EpsgEllipsoid {
	ellipsoid := &EpsgEllipsoid{Name: name, Epsg: epsg}

	ellipsoid.Ellipsoid = NewEllipsoidWithSemiAxis(a, b)
	return ellipsoid
}

func NewEpsgEllipsoidWithFlattening(name string, epsg int, a float64, f float64) *EpsgEllipsoid {
	ellipsoid := &EpsgEllipsoid{Name: name, Epsg: epsg}

	ellipsoid.Ellipsoid = NewEllipsoidWithFlattening(a, f)
	return ellipsoid
}
