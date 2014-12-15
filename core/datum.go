package core

type Datum struct {
	CentralMeridien float64
	EquateurOrigin  float64
	Ellipsoid       *Ellipsoid
}
