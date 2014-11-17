package projection

// Represents a projection from a datum to a plane map
//
// Project represents the transition from geographic to map coordinates
// Reverse represents the transition from  map coordinates to geographic
type Projection interface {
	Project(lambda, pi) (x, y, float64)
	Reverse(x, y float64) (lambda, phi float64)
}
