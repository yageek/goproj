package math

import (
	"math"
)

func Sec(x float64) float64 {
	return 1.0 / (math.Cos(x) * math.Cos(x))
}
