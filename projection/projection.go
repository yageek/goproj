package projection

import . "github.com/yageek/goproj/datum"

// Represents a projection from a datum to a map
//
// Project represents the transition from geographic to map coordinates (rad -> meters)
// Reverse represents the transition from  map coordinates to geographic (meters -> rad)
type Projection interface {
	Forward(lambda, phi float64, datum *Datum) (x, y float64)
}
