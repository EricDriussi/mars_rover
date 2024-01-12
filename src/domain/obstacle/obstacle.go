package obstacle

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle/bigRock"
	"mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/size"
)

type Obstacle interface {
	IsBeyond(Size) bool
	Occupies(AbsoluteCoordinate) bool
	Coordinates() Coordinates
}

func CreateObstacle(coordinates Coordinates) (Obstacle, error) {
	if coordinates.Amount() == 1 {
		return smallRock.In(coordinates.First())
	}
	return bigRock.In(coordinates)
}

func MaxSize() int {
	return bigRock.MaxSize
}
