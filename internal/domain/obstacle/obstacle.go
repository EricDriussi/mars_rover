package obstacle

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/size"
)

type Obstacle interface {
	IsBeyond(Size) bool
	Occupies(AbsoluteCoordinate) bool
	Coordinates() []AbsoluteCoordinate
}
