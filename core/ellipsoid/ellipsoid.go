//This package manages the ellipsoid
//go:generate go run generator/generator.go
//go:generate gofmt -w ellipsoid_index.go
package ellipsoid

import (
	"math"
)

//The informations contained in an Ellipsoid
type Ellipsoid struct {
	SemiMajorAxis     float64
	SemiMinorAxis     float64
	InverseFlattening float64
	Excentricity      float64
	Excentricity2     float64
}

func NewEllipsoidWithSemiAxis(a float64, b float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.SemiMinorAxis = b
	p.InverseFlattening = a / (a - b)
	p.Excentricity2 = (math.Pow(a, 2) - math.Pow(b, 2)) / math.Pow(a, 2)
	p.Excentricity = math.Sqrt(p.Excentricity2)

	return p
}

func NewEllipsoidWithInverseFlattening(a float64, inv_f float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.InverseFlattening = inv_f

	p.SemiMinorAxis = a * (1 - 1/inv_f)
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
	return &EpsgEllipsoid{Name: name, Epsg: epsg, Ellipsoid: NewEllipsoidWithSemiAxis(a, b)}
}

func NewEpsgEllipsoidWithInverseFlattening(name string, epsg int, a float64, f float64) *EpsgEllipsoid {
	return &EpsgEllipsoid{Name: name, Epsg: epsg, Ellipsoid: NewEllipsoidWithInverseFlattening(a, f)}
}
