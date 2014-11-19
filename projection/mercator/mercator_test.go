package mercator

import (
	"encoding/json"
	. "github.com/yageek/goproj/datum"
	"io/ioutil"
	"math"
	"testing"
)

type SnyderTest struct {
	Delta float64
	R     float64
	X     float64
	Tests []SnyderValue
}

type SnyderValue struct {
	Phi float64
	Y   float64
	K   float64
}

func TestSphereForward(t *testing.T) {

	test := SnyderTest{}
	data, _ := ioutil.ReadFile("sphere_values.json")
	err := json.Unmarshal(data, &test)
	if err != nil {
		t.Error(err)
	}

	datum := &Datum{CentralMeridien: 0.0}

	sp := &SphereProjection{Radius: test.R}

	for _, testcase := range test.Tests {
		phi := testcase.Phi * math.Pi / 180.0
		_, y := sp.Forward(test.X, phi, datum)

		if math.Abs(y-testcase.Y) > test.Delta {
			t.Errorf("Failed with Phi:%.6f | Expected: %.6f | Computed: %.6f | Delta: %f\n", testcase.Phi, testcase.Y, y, test.Delta)
		}

		k := sp.ScaleFactor(0, phi)

		if math.Abs(k-testcase.K) > test.Delta {
			t.Errorf("Failed K for Phi: %.6f | Expected:%.6f | Computed:%.6f\n", testcase.Phi, testcase.K, k)
		}

	}

}
func TestEllipseForward(t *testing.T) {
	test := SnyderTest{}
	data, _ := ioutil.ReadFile("ellipse_values.json")

	err := json.Unmarshal(data, &test)

	if err != nil {
		t.Error(err)
	}

	ellipsoid := NewEllipseWithSemiAxis(1, 1)
	datum := &Datum{CentralMeridien: 0.0, Ellipsoid: ellipsoid}

	sp := &EllipseProjection{}

	for _, testcase := range test.Tests {
		phi := testcase.Phi * math.Pi / 180.0
		_, y := sp.Forward(test.X, phi, datum)

		if math.Abs(y-testcase.Y) > test.Delta {
			t.Errorf("Failed with Phi:%.6f | Expected: %.6f | Computed: %.6f | Delta: %f\n", testcase.Phi, testcase.Y, y, test.Delta)
		}

		k := sp.ScaleFactor(0, phi)

		if math.Abs(k-testcase.K) > test.Delta {
			t.Errorf("Failed K for Phi: %.6f | Expected:%.6f | Computed:%.6f\n", testcase.Phi, testcase.K, k)
		}
	}
}
