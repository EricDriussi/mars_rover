package positionCalculator

import (
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/rover/planetMap"
)

func Forward(direction Direction, coordinate AbsoluteCoordinate, planetMap Map) AbsoluteCoordinate {
	futureCoordinate := *coord.SumOf(coordinate, direction.RelativeCoordinateAhead())
	if isOutOfBounds(futureCoordinate, planetMap) {
		futureCoordinate = wrapAround(futureCoordinate, planetMap)
	}
	return futureCoordinate
}

func Backward(direction Direction, coordinate AbsoluteCoordinate, planetMap Map) AbsoluteCoordinate {
	futureCoordinate := *coord.SumOf(coordinate, direction.RelativeCoordinateBehind())
	if isOutOfBounds(futureCoordinate, planetMap) {
		futureCoordinate = wrapAround(futureCoordinate, planetMap)
	}
	return futureCoordinate
}

func isOutOfBounds(coordinate AbsoluteCoordinate, planetMap Map) bool {
	return planetMap.IsOutOfBounds(coordinate)
}

func wrapAround(coordinate AbsoluteCoordinate, planetMap Map) AbsoluteCoordinate {
	return *absoluteCoordinate.From(
		wrap(coordinate.X(), planetMap.Width()),
		wrap(coordinate.Y(), planetMap.Height()),
	)
}

func wrap(point, limit int) int {
	if point > limit {
		return 0
	} else if point < 0 {
		return limit
	}
	return point
}
