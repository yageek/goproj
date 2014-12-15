package datum

import "github.com/go-gis/goproj/core/ellipsoid"

type Datum struct {
	CentralMeridien float64
	EquateurOrigin  float64
	Ellipsoid       *ellipsoid.Ellipsoid
}
