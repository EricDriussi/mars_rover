package obstacle

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle/bigRock"
	"mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/size"
)

type Obstacle interface {
	IsBeyond(Size) bool
	Occupies(AbsoluteCoordinate) bool
	Coordinates() []AbsoluteCoordinate
}

func CreateObstacle(coordinates ...AbsoluteCoordinate) (Obstacle, error) {
	if len(coordinates) == 1 {
		return smallRock.In(coordinates[0]), nil
	}
	return bigRock.In(coordinates...)
}
