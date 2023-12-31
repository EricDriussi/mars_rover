package coordinate

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
)

type Coordinate interface {
	X() int
	Y() int
}

func SumOf(coordinate AbsoluteCoordinate, relative RelativeCoordinate) *AbsoluteCoordinate {
	return absoluteCoordinate.From(coordinate.X()+relative.X(), coordinate.Y()+relative.Y())
}
