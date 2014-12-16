package lambert

import (
	"github.com/go-gis/goproj/core"
	"math"
)

type ConformalConic struct {
	Phi1    float64
	Phi2    float64
	cosPhi1 float64
	cosPhi2 float64
}

func NewConformalConic(phi1, phi2 float64) *ConformalConic {

	return &ConformalConic{
		Phi1:    phi1,
		Phi2:    phi2,
		cosPhi1: math.Cos(cosPhi1),
		cosPhi2: math.Cos(cosPhi2),
	}
}

func (l *ConformalConic) Forward(lambda, phi float64, datum *core.Datum) (x, y float64) {

}
func (l *ConformalConic) Inverse(x, y float64, datum *core.Datum) (lambda, phi float64) {
}

//Forward function for the sphere
func (l *ConformalConic) sphereForward(lambda, phi float64, datum *core.Datum) (x, y float64) {

	n := math.Log(l.cosPhi1/l.cosPhi2) * math.Log(math.Tan(core.PI_4+l.Phi2/2)/math.Tan(core.PI_4+l.Phi1/2))
	F := math.Cos(l.Phi1) * math.Pow(math.Tan(core.PI_4+l.Phi1/2), n) / n
	RF := datum.Ellipsoid.SemiMajorAxis * F
	rho := RF / math.Pow(math.Tan(core.PI_4+phi/2), n)
	//	rho_0 := RF / math.Pow(math.Tan(core.PI_4+datum.EquateurOrigin/2), n)

}

//Inverse function for the sphere
func (l *ConformalConic) sphereInverse(x, y float64, datum *core.Datum) (lambda, phi float64) {

}

//Forward function for the sphere
func (l *ConformalConic) ellipsoidForward(lambda, phi float64, datum *core.Datum) (x, y float64) {

}

//Inverse function for the sphere
func (l *ConformalConic) ellipsoidInverse(x, y float64, datum *core.Datum) (lambda, phi float64) {
}
