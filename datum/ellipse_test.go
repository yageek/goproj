package datum

import (
	"math"
	"testing"
)

type testEllipseParams struct {
	Name       string
	A, B, F, E float64
	OB, OF, OE float64
}

const (
	AXIS_DELTA       = 1e-2
	FLATTENING_DELTA = 1e-7
)

var ellipsesTested []testEllipseParams

func init() {
	ellipsesTested = []testEllipseParams{
		{Name: "WGS84", A: 6378137, B: -1, F: 1 / 298.257222101, OB: 6356752.314140355847852106, OE: 0.08181919132},

		{Name: "Clarke1880-Axis", A: 6378249.145, B: 6356514.870, F: -1, OF: 1 / 293.465, OE: 0.08181919132},
		{Name: "Clarke1880-Flattening", A: 6378249.145, B: -1, F: 1 / 293.465, OB: 6356514.870, OE: 0.08181919132},
	}
}

func TestEllipseParametersCreations(t *testing.T) {
	for _, params := range ellipsesTested {

		var out *Ellipse
		if params.B < 0 {
			out = NewEllipseWithFlattening(params.A, params.F)
		} else {
			out = NewEllipseWithSemiAxis(params.A, params.B)
		}

		if params.B < 0 && math.Abs(params.OB-out.SemiMinorAxis) > AXIS_DELTA {
			t.Errorf("%s - Semi Minor Axis - Expected:%.10f | Computed:%.10f", params.Name, params.OB, out.SemiMinorAxis)
		}

		if params.F < 0 && math.Abs(params.OF-out.Flattening) > FLATTENING_DELTA {
			t.Errorf("%s - Flattening - Expected:%.10f | Computed:%.10f", params.Name, params.OF, out.Flattening)
		}
		// if math.Abs(params.OE-out.Excentricity) > ACCEPTABLE_DELTA {
		// 	t.Errorf("%s - Excentricity - Expected:%.10f | Computed:%.10f", params.Name, params.OE, out.Excentricity)
		// }

	}
}
