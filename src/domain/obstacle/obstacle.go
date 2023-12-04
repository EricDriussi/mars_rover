package obstacle

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/size"
)

type Obstacle interface {
	IsBeyond(Size) bool
	Occupies(AbsoluteCoordinate) bool
	Coordinates() []AbsoluteCoordinate
}
