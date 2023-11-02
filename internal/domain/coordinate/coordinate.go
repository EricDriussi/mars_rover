package coordinate

import "mars_rover/internal/domain/size"

type Coordinate interface {
	// TODO: should be the rover's responsibility
	WrapIfOutOf(size.Size)
	Equals(other Coordinate) bool
	X() int
	Y() int
	// TODO: Sum(other Coordinate)
}
