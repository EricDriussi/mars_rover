package obstacle

import (
	abs "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/size"
)

type Obstacle interface {
	IsBeyond(size.Size) bool
	Occupies(abs.AbsoluteCoordinate) bool
	Coordinates() []abs.AbsoluteCoordinate
}
