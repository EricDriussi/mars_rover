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

func SumOf(coordinateOne AbsoluteCoordinate, coordinateTwo RelativeCoordinate) *AbsoluteCoordinate {
	return absoluteCoordinate.From(coordinateOne.X()+coordinateTwo.X(), coordinateOne.Y()+coordinateTwo.Y())
}