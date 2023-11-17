package obstacle

import (
	coord "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/size"
)

// TODO: add obstacleS
type Obstacle interface {
	IsBeyond(size.Size) bool
	Occupies(coord.AbsoluteCoordinate) bool
	Coordinates() []coord.AbsoluteCoordinate
}
