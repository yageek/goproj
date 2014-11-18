package mercator

type Mercator struct {
}

func (e *Mercator) Forward(lambda, phi float64, datum *Datum) (x, y float64) {

	x = datum.Ellipsoid.SemiMajorAxis * (lambda - datum.Ellipsoid.CentralMeridien)
	y = 
}
