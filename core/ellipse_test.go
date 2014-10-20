package core

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
	ACCEPTABLE_DELTA = 1e-10
)

var ellipsesTested []testEllipseParams

func init() {
	ellipsesTested = []testEllipseParams{
		{Name: "WGS84", A: 6378137, B: -1, F: 1 / 298.257222101, OB: 6356752.314140355847852106, OE: 0.08181919132},
	}
}

func isEllipseAsExpected(test testEllipseParams, out *EllipseParameters) bool {
	return (math.Abs(test.OB-out.SemiMinorAxis) < ACCEPTABLE_DELTA && math.Abs(test.OF-out.Flattening) < ACCEPTABLE_DELTA && math.Abs(test.OE-out.Excentricity) < ACCEPTABLE_DELTA)
}

func TestEllipseParametersCreations(t *testing.T) {
	for _, params := range ellipsesTested {

		var out *EllipseParameters
		if params.B < 0 {
			out = NewEllipseParametersWithFlattening(params.A, params.F)
		} else {
			out = NewEllipseParametersWithSemiAxis(params.A, params.B)
		}

		if params.B == -1 && math.Abs(params.OB-out.SemiMinorAxis) > ACCEPTABLE_DELTA {
			t.Errorf("%s - Semi Minor Axis - Computed:%f | Expected:%f", params.Name, params.OB, out.SemiMinorAxis)
		}

		if params.F == -1 && math.Abs(params.OF-out.Flattening) > ACCEPTABLE_DELTA {
			t.Errorf("%s - Flattening - Computed:%f | Expected:%f", params.Name, params.OF, out.Flattening)
		}
		if math.Abs(params.OE-out.Excentricity) > ACCEPTABLE_DELTA {
			t.Errorf("%s - Excentricity - Computed:%f | Expected:%f", params.Name, params.OE, out.Excentricity)
		}

	}
}
