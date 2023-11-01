package coordinate

import "mars_rover/internal/domain/size"

type Coordinate interface {
	WrapIfOutOf(size.Size)
	IsOutsideOf(size.Size) bool
	Equals(other Coordinate) bool
	X() int
	Y() int
}
