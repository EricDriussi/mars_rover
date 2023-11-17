package coordinate

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
)

// TODO: add coordinateS
type Coordinate interface {
	X() int
	Y() int
}

func SumOf(coordinateOne absoluteCoordinate.AbsoluteCoordinate, coordinateTwo relativeCoordinate.RelativeCoordinate) *absoluteCoordinate.AbsoluteCoordinate {
	return absoluteCoordinate.From(coordinateOne.X()+coordinateTwo.X(), coordinateOne.Y()+coordinateTwo.Y())
}
