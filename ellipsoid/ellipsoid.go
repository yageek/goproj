//This package manages the ellipsoid
//go:generate go run generator/generator.go
//go:generate gofmt -w ellipsoid_index.go
package ellipsoid

//The informations contained in an Ellipsoid
type Ellipsoid struct {
	SemiMajorAxis     float64
	SemiMinorAxis     float64
	InverseFlattening float64
}

func NewEllipsoidWithSemiAxis(a float64, b float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.SemiMinorAxis = b

	return p
}

func NewEllipsoidWithInverseFlattening(a float64, inv_f float64) *Ellipsoid {
	p := new(Ellipsoid)

	p.SemiMajorAxis = a
	p.InverseFlattening = inv_f

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
