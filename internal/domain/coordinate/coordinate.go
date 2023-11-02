package coordinate

import "mars_rover/internal/domain/size"

type Coordinate interface {
	WrapIfOutOf(size.Size)
	Equals(other Coordinate) bool
	X() int
	Y() int
}
