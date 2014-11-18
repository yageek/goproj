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

func TestForward(t *testing.T) {

	test := SnyderTest{}
	data, _ := ioutil.ReadFile("sphere_forward.json")
	json.Unmarshal(data, &test)

	datum := &Datum{CentralMeridien: 0.0}

	sp := &SphereProjection{Radius: test.R}

	for _, testcase := range test.Tests {
		_, y := sp.Forward(test.X, testcase.Phi*math.Pi/180.0, datum)

		if math.Abs(y-testcase.Y) > test.Delta {
			t.Errorf("Failed with Phi:%.6f | Expected: %.6f | Computed: %.6f | Delta: %f\n", testcase.Phi, testcase.Y, y, test.Delta)
		}
	}

}
