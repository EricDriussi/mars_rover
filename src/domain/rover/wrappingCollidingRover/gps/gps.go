package gps

import (
	coord "mars_rover/src/domain/coordinate"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/planetMap"
)

type GPS struct {
	rover Rover
}

func Bind(rover Rover) GPS {
	return GPS{rover}
}

func (this *GPS) Ahead() AbsoluteCoordinate {
	return calculate(
		this.rover.Coordinate(),
		this.rover.Direction().RelativeCoordinateAhead(),
		this.rover.Map(),
	)
}

func (this *GPS) Behind() AbsoluteCoordinate {
	return calculate(
		this.rover.Coordinate(),
		this.rover.Direction().RelativeCoordinateBehind(),
		this.rover.Map(),
	)
}

func calculate(absolute AbsoluteCoordinate, relative RelativeCoordinate, planetMap Map) AbsoluteCoordinate {
	futureCoordinate := *coord.SumOf(absolute, relative)
	if planetMap.IsOutOfBounds(futureCoordinate) {
		futureCoordinate = wrapAround(futureCoordinate, planetMap)
	}
	return futureCoordinate
}

func wrapAround(coordinate AbsoluteCoordinate, planetMap Map) AbsoluteCoordinate {
	return *absoluteCoordinate.Build(
		// Planet size starts at 1, rover movement starts at 0
		wrap(coordinate.X(), planetMap.Width()-1),
		wrap(coordinate.Y(), planetMap.Height()-1),
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
